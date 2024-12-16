package quic

import (
	"context"
	"errors"
	"net"
	"sync"
	// "errors"
)

var DefaultMaxStreamData = 1000
var DefaultMaxStreams = 1000

type ConnectionOps interface {
	Open(raddr *net.UDPAddr, laddr *net.UDPAddr) (*Connection, error)
	Listen(laddr *net.UDPAddr, cb func(conn net.UDPConn))
}

type SendBlokedChannel chan struct {
	state  bool
	reason string
}

type Connection struct {
	ConnectionIdManager
	UdpConn        *net.UDPConn
	Raddr          *net.UDPAddr //ipv4:udp-port
	Laddr          *net.UDPAddr
	Streams        []*Stream //first stream id indexed 0 (used in CanHandleMoreStreams)
	MaxStreamData  int
	MaxStreams     int
	MaxIds         int    //the maximum number of connection id can be given to this connection
	HandshakeState string //onging, completed, failed
	ctx            context.Context

	onClose func(con *Connection) error

	sendLock *sync.Mutex
	//if any go routine should block sending of streams
}

func createQuicConnection(
	UdpConn *net.UDPConn,
	Raddr *net.UDPAddr,
	Laddr *net.UDPAddr,
	MaxStreamData int,
	MaxStreams int,
	MaxIds int,
	HandshakeState string,
	ctx context.Context,
	onClose func(con *Connection) error,
	sendLock *sync.Mutex,
) (*Connection, error) {
	var err error = nil

	if UdpConn == nil {
		err = errors.New("udp connection can't be nil")
	}
	if Raddr == nil {
		//get it from udpConn
	}

	if MaxStreamData == 0 {
		MaxStreamData = DefaultMaxStreamData
	}

	if MaxStreams == 0 {
		MaxStreams = DefaultMaxStreams
	}

	if ctx == nil {
		ctx = context.Background()
	}

	if err != nil {
		return nil, err
	}

	return &Connection{
		UdpConn:        UdpConn,
		Raddr:          Raddr,
		Laddr:          Laddr,
		MaxStreamData:  MaxStreamData,
		MaxStreams:     MaxStreams,
		MaxIds:         MaxIds,
		HandshakeState: "completed",
		ctx:            ctx,
		onClose:        onClose,
		sendLock:       sendLock,
	}, nil
}

func (c *Connection) Close() error {
	if c.onClose != nil {
		c.onClose(c)
	}

	// do the acual closing
	panic("conn.Close not implemeted yet")
}

func (c *Connection) OpenStream(stream_type uint8) (*Stream, error) {
	//create a stream
	//bind it to the connection
	panic("OpenStream not implemented yet")
	// return nil, errors.New("open stream not implemented yet")
}

func (c *Connection) HasStream(stream *Stream) bool {
	if stream.connectionIdx > len(c.Streams) {
		return false
	}
	return c.Streams[stream.connectionIdx] == stream
}

func (c *Connection) OpenStreamBi(stream_type uint8) (*Stream, error) {
	panic("OpenStreamBi not implemented yet")
	// return nil, errors.New("OpenStreamBi not implemented yet")
}

func (c *Connection) CanHandleMoreStreams(stream_type uint8) bool {
	first_stream_id := c.Streams[0].Id
	nextId := c.GetNextStreamId(stream_type)
	return nextId <= uint64(c.MaxStreams*4+int(first_stream_id))
}

func (c *Connection) GetNextStreamId(streamType uint8) uint64 {
	var streamId uint64

	return streamId
}

func GetConnectionById(id uint) *Connection {
	return &Connection{}
}

func (c *Connection) Context() context.Context {
	return c.ctx
}
