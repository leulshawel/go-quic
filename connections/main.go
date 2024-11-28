package connections

import (
	"fmt"
	"net"
)

// creates a new QUIC connection as of RFC 9000
func NewQuicConnection(address net.UDPAddr) (*Connection, error) {

	//the quic connection "object" to be returned
	var quicConn = Connection{
		address: address,
		Streams: nil,
	}

	//create a udp 'connection'
	udpConn, err := net.DialUDP(address.Network(), nil, &address)
	if err != nil {
		return nil, fmt.Errorf("Couldn't send udp dial on %s", address.String())
	}

	//create a tls handshake to be sent over udp

	//do tls handshake
	//send the connReq
	//wait for a response (in a different thread) incase there is sth todo

	//do the quic connection handshake

	quicConn.UdpConn = udpConn
	return &quicConn, nil
}
