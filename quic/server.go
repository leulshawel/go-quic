package quic

import (
	"context"
	"errors"
	"net"
	"sync"
)

// Listener
// a QUIC transport running on a UDP port

var DefaultListenerOps = struct {
	DefaultMaxStreamData int
	DefaultMaxStreams    int
	DefaultMaxConns      int
}{
	DefaultMaxStreamData: 5,
	DefaultMaxStreams:    5,
	DefaultMaxConns:      5}

type acceptChann struct {
	connChan chan *Connection
	errChan  chan error
}

type Listener struct {
	s              *Server
	UdpConn        *net.UDPConn
	laddr          *net.UDPAddr
	callBack       func(context.Context, *Listener) error
	MaxConnections int
	acceptChann    acceptChann
	acceptQeue     chan net.Addr
	errorQeue      chan error
	connections    []*Connection
	ctx            context.Context
	Cancel         context.CancelFunc
	MaxStreamData  int
	MaxStreams     int
	MaxConns       int
	onClose        func(con *Connection) error
	sendLock       sync.Mutex
	isListenning   bool
}

func CreateNewListener(
	s *Server,
	udpConn *net.UDPConn,
	laddr *net.UDPAddr,
	cb func(context.Context, *Listener) error,
	connections []*Connection,
	ctx context.Context,
	maxStreamData int,
	maxStreams int,
	maxConns int,
	onClose func(con *Connection) error,
) (*Listener, error) {
	if udpConn != nil && laddr != nil {
		if laddr == udpConn.LocalAddr() {
			return nil, errors.New("udpConn.LocalAddr != laddr")
		}
	}

	if maxConns == 0 {
		maxConns = DefaultListenerOps.DefaultMaxConns
	}
	if maxStreamData == 0 {
		maxStreamData = DefaultListenerOps.DefaultMaxStreamData
	}
	if maxStreams == 0 {
		maxStreams = DefaultMaxStreams
	}

	ctx_, cancel := context.WithCancel(ctx)
	l := &Listener{
		s:             s,
		UdpConn:       udpConn,
		laddr:         laddr,
		Cancel:        cancel,
		callBack:      cb,
		connections:   make([]*Connection, 1),
		ctx:           ctx_,
		MaxStreamData: maxStreamData,
		MaxStreams:    maxStreams,
		MaxConns:      maxConns,
		onClose:       onClose,
	}

	return l, nil
}

// accept accepts connections that have finished the handshake
// takes a callback func where the user app can access the connection and return error if it doesn't want the connection
func (l *Listener) Accept(
	cb func(con *Connection) error,
) {

	go l.accept()
	for {
		select {
		case con := <-l.acceptChann.connChan:
			if cb != nil {
				if err := cb(con); err == nil {
					l.connections = append(l.connections, con)
				}
			}
		case err := <-l.acceptChann.errChan:
			l.errorQeue <- err
		}
	}
}

// this is the actual function where we accept connections but it
func (l *Listener) accept() {
	//This is where we have to do the handshake and create a quic connection if successfull
	for {
		select {
		case <-l.acceptQeue:
			//for actual connections the fields will be filled accordingly
			conn, err := createQuicConnection(
				l.UdpConn,
				nil,
				l.laddr,
				DefaultListenerOps.DefaultMaxStreamData,
				DefaultListenerOps.DefaultMaxStreams,
				DefaultListenerOps.DefaultMaxConns,
				"completed",
				l.ctx,
				l.onClose,
				&l.sendLock,
			)

			if err != nil {
				l.acceptChann.errChan <- err
			} else {
				l.acceptChann.connChan <- conn
			}
		case <-l.ctx.Done():
			l.errorQeue <- l.ctx.Err()
			return
		}
	}
}

func (l *Listener) CloseAll() []error {
	var errors []error = nil

	//loop through the connections and close them all and append errors to the slice
	for _, con := range l.connections {
		if err := con.Close(); err != nil {
			errors = append(errors, err)
		}
	}

	return errors
}

func (l *Listener) errorHandler() {
	for {
		err := <-l.errorQeue
		if err == context.Canceled {
			return
		}
	}

}

func (l *Listener) Listen(laddr *net.UDPAddr) error {
	errorChan := make(chan *error)
	go l.listen(errorChan)

	err := <-errorChan
	return *err

}

func (l *Listener) listen(errChan chan *error) error {
	udpConn, err := net.ListenUDP("udp", l.laddr)
	if err != nil {
		errChan <- &err
	} else {
		errChan <- nil
	}

	l.UdpConn = udpConn

	if l.callBack != nil {
		go l.callBack(l.ctx, l)
	}

	l.isListenning = true

	buffer := make([]byte, 1024)
	l.errorHandler()

	for {
		//this is blocking so we won't get to the for loop below till we get some datagrams or error
		_, addr, err := udpConn.ReadFrom(buffer)
		if err != nil {
			l.errorQeue <- err
		}

		select {
		case <-l.ctx.Done():
			l.errorQeue <- l.ctx.Err()
			return nil
		case l.acceptQeue <- addr:
		}

	}

}

// Server
// a central server that can manage multiple listeners
type Server struct {
	isListenning bool
	Listeners    []*Listener
	ctx          context.Context    //the parent context fo the entire server (all listeners stop if cancelled)
	Cancel       context.CancelFunc //cancel function for Server.ctx
}

func CreateNewServer(ctx context.Context) *Server {

	ctx_, cancel := context.WithCancel(ctx)
	s := &Server{
		ctx:    ctx_,
		Cancel: cancel,
	}

	return s
}

func (s *Server) AddListener(l *Listener) error {
	//check if a listener can be managed by this server
	//check if sever not already added
	if l.s == s {
		err := errors.New("listener already in the server")
		return err
	}
	s.Listeners = append(s.Listeners, l)
	l.s = s

	return nil
}

// func (s *Server) findListenerByPort(port int) *Listener {
// 	for _, l := range s.Listeners {
// 		if l.laddr.Port == port {
// 			return l
// 		}
// 	}
// 	return nil
// }

func (s *Server) Listen(lis *Listener, laddr *net.UDPAddr, cb func(context.Context, *Listener) error) (*Listener, error) {
	laddr_ := laddr != nil
	lis_ := lis != nil

	if lis_ && laddr_ {
		return nil, errors.New("needs a listener or a local addr to listen")
	}

	var l *Listener
	if lis_ && lis.laddr != nil {
		if lis.s != s {
			s.AddListener(lis)
		}
		l = lis
	} else {
		if len(s.Listeners) == 0 {
			if !laddr_ {
				return nil, errors.New("server has no listeners. laddr is required")
			}

			l = &Listener{
				laddr:        laddr,
				callBack:     cb,
				isListenning: false,
			}
			ctx, cancel := context.WithCancel(s.ctx)

			l.ctx = ctx
			l.Cancel = cancel
			//add it to this server
			s.Listeners = append(s.Listeners, l)
		} else if len(s.Listeners) == 1 {
			l = s.Listeners[0]
		} else {
			for _, ll := range s.Listeners {
				if !ll.isListenning {
					l = ll
					break
				}
			}
		}
	}
	errorChan := make(chan *error)
	go l.listen(errorChan)

	err := <-errorChan
	if err != nil {
		return nil, *err
	}

	s.isListenning = true
	return l, nil
}

// stop all listeners on the server gracefully.
// Handle or return errors if any
func (s *Server) Down() error {
	return nil
}

func (s *Server) ForceDown() {
	s.Cancel()
}
