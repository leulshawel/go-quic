package main

//this is a test/example implementation of a server application that uses
//the go-quic api

import (
	"context"
	"fmt"
	"go-quic/quic"
	"net"
)

func main() {
	//udp address to listen on
	udpAddr := &net.UDPAddr{
		IP:   net.IPv4(0, 0, 0, 0),
		Port: 8080,
	}

	//we can cancel this context and every
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

	l.Accept(nil)
	fmt.Print("Helllo")

}
