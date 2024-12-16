package quic

import (
	"context"
	"fmt"
	"net"
	"sync"
)

// Listener
// a QUIC transport running on a UDP port

type Listener struct {
	UdpConn              *net.UDPConn
	laddr                *net.UDPAddr
	callBack             func(context.Context, *Listener) error
	MaxConnections       int
	acceptQeue           chan net.Addr
	errorQeue            chan *error
	connections          []*Connection
	Ctx                  context.Context
	DefaultMaxStreamData int
	DefaultMaxStreams    int
	DefaultMaxConnIds    int
	onClose              func(con *Connection) error
	sendLock             sync.Mutex
	isListenning         bool
}

// accept accepts connections that have finished the handshake
func (l *Listener) Accept() {
	var connQeue = make(chan *Connection)

	go l.accept(connQeue)

	select {
	case con := <-connQeue:
		l.connections = append(l.connections, con)
	case err := <-l.errorQeue:
		l.errorQeue <- err
	}
}

// this is the actual function where we accept connections but it
func (l *Listener) accept(con chan *Connection) {
	//This is where we have to do the handshake and create a quic connection if successfull
	for {

		select {
		case <-l.acceptQeue:
			//for actual connections the fields will be filled accordingly
			conn, err := createQuicConnection(
				l.UdpConn,
				nil,
				l.laddr,
				l.DefaultMaxStreamData,
				l.DefaultMaxStreams,
				l.DefaultMaxConnIds,
				"completed",
				l.Ctx,
				l.onClose,
				&l.sendLock,
			)

			if err != nil {
				l.errorQeue <- &err
			} else {
				con <- conn
			}
		case <-l.Ctx.Done():
			goto exit
		}
	}

exit:
	return
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

func (l *Listener) listen() error {
	udpConn, err := net.ListenUDP("udp", l.laddr)
	if err != nil {
		return err
	}

	l.UdpConn = udpConn

	if l.callBack != nil {
		go l.callBack(l.Ctx, l)
	}

	l.isListenning = true

	buffer := make([]byte, 1024)

	for {
		//this is blocking so we won't get to the for loop below till we get some datagrams or error
		_, addr, err := udpConn.ReadFrom(buffer)
		if err != nil {
			l.errorQeue <- &err
		}

		select {
		case <-l.Ctx.Done():
			goto exit
		case l.acceptQeue <- addr:
		}

		fmt.Println(l.acceptQeue)

	}

exit:
	//Do all the necessary cleanups before returning
	return nil

}

// Server
// a central server that can manage multiple listeners

type Server struct {
	isListenning bool
	Listeners    []*Listener
	Ctx          context.Context
}

func CreateNewServer(ctx context.Context) *Server {
	if ctx == nil {
		ctx = context.Background()
	}
	s := &Server{
		Ctx: ctx,
	}

	return s
}

func (s *Server) AddListener(l *Listener) {
	//check if a listener can be managed by this server
	s.Listeners = append(s.Listeners, l)
}

func (s *Server) Listen(laddr *net.UDPAddr, cb func(context.Context, *Listener) error) (*Listener, error) {

	var l *Listener
	if len(s.Listeners) == 0 {
		fmt.Println("Adding a listner to the server")
		l = &Listener{
			laddr:        laddr,
			callBack:     cb,
			isListenning: false,
		}
		ctx, _ := context.WithCancel(s.Ctx)

		l.Ctx = ctx
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
	go l.listen()
	s.isListenning = true

	return l, nil
}

func (s *Server) Wait() int {
	for {
		if !s.isListenning {
			break
		}
	}
	return 0
}
