package main

import (
	"context"
	"fmt"
	"go-quic/quic"
	"net"
)

func main() {

	udpAddr := &net.UDPAddr{
		IP:   net.IPv4(0, 0, 0, 0),
		Port: 8080,
	}

	s := quic.CreateNewServer(context.Background())

	s.Listen(
		udpAddr,
		func(ctx context.Context, udpConn *net.UDPConn) error {
			fmt.Printf("server listening on %d\n", udpAddr.Port)
			return nil
		})
}
