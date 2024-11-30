package connections

import (
	"fmt"
	"net"
)

type Connection struct {
	ConnectionId            []ConnectionId
	ActiveConnectionIdLimit int
	UdpConn                 *net.UDPConn
	Raddr                   net.UDPAddr //ipv4:udp-port
	Laddr                   net.UDPAddr
	Streams                 []uint64
}

// creates a new QUIC connection as of RFC 9000
func NewQuicConnection(raddr net.UDPAddr, laddr net.UDPAddr) (*Connection, error) {

	//the quic connection "object" to be returned
	var quicConn = Connection{
		Raddr:   raddr,
		Laddr:   laddr,
		Streams: nil,
	}

	//create a udp 'connection'
	udpConn, err := net.DialUDP(raddr.Network(), &laddr, &raddr)
	if err != nil {
		return nil, fmt.Errorf("couldn't send udp dial on %s from %s", raddr.String(), laddr.String())
	}

	//create a tls handshake to be sent over udp

	//do tls handshake
	//send the connReq
	//wait for a response (in a different thread) incase there is sth todo

	//do the quic connection handshake

	quicConn.UdpConn = udpConn
	return &quicConn, nil
}

func (c *Connection) GetNextStreamId(streamType uint8) uint64 {
	var streamId uint64

	return streamId
}
