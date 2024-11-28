package main

import (
	"fmt"
	"go-quic/connections"
	"net"
)

func main() {
	var address net.UDPAddr = net.UDPAddr{
		IP:   net.IPv4(127, 0, 0, 1),
		Port: 1234,
	}

	quicConn, err := connections.NewQuicConnection(address)
	if err != nil {
		fmt.Println(err)
	}

	defer quicConn.UdpConn.Close()

	quicConn.UdpConn.Write([]byte("hello"))

	fmt.Println(quicConn)
}
