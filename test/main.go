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
	ParentContext, cancel := context.WithCancel(context.Background())

	s := quic.CreateNewServer(ParentContext)
	defer cancel()

	l, err := s.Listen(
		udpAddr,
		func(ctx context.Context, l *quic.Listener) error {
			fmt.Printf("server listening on %d\n", l.UdpConn.LocalAddr())
			return nil
		})

	if err != nil {
		fmt.Println(err)

	}

	l.Accept()

}
