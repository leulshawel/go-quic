package quic

import (
	"context"
	"fmt"
	"net"
	// "errors"
)

type ConnectionOps interface {
	Open(raddr *net.UDPAddr, laddr *net.UDPAddr) (*Connection, error)
	Listen(laddr *net.UDPAddr, cb func(conn net.UDPConn))
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
	ZeroRTTUsed    bool   //is zero rtt used for this connection
	HandshakeState string //onging, completed, failed
	context        context.Context

	sendBlocked chan struct {
		state  bool
		reason string
	} //if any go routine should block sending of streams
}

// creates a new QUIC connection  with a server as of RFC 9000
func OpenConnection(raddr *net.UDPAddr, laddr *net.UDPAddr) (*Connection, error) {

	//create a udp 'connection'
	udpConn, err := net.DialUDP(raddr.Network(), laddr, raddr)
	if err != nil {
		return nil, fmt.Errorf("couldn't send udp dial on %s from %s", raddr.String(), laddr.String())
	}

	//the quic connection "object" to be returned
	var quicConn = Connection{
		Raddr:   raddr,
		Laddr:   laddr,
		UdpConn: udpConn,
	}

	//create a tls handshake to be sent over udp

	//do tls handshake
	//send the connReq
	//wait for a response (in a different thread) incase there is sth todo

	//do the quic connection handshake

	quicConn.UdpConn = udpConn
	return &quicConn, nil
}

// liten for a connection from a client
func Listen(laddr *net.UDPAddr, cb func(net.UDPConn)) {
	panic("listen not implemented yet!")
}

func ListenEarly(laddr *net.UDPAddr, cb func(net.UDPConn)) {
	panic("listen not implemented yet!")

}

func (c *Connection) Close() error {
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
	return c.context
}
