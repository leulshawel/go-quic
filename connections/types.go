package connections

import (
	"go-quic/streams"
	"net"
)

// a QUIC connect request
type Connreq struct {
}

type Connection struct {
	UdpConn *net.UDPConn
	address net.UDPAddr       //ipv4:udp-port
	Streams []*streams.Stream //an arrat of stream pointers in this connection
}
