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
	udpAddr := &net.UDPAddr{IP: net.IPv4(127, 0, 0, 1), Port: 8080}

	s := quic.CreateNewServer(context.Background()) //create a server

	var l *quic.Listener
	var err error
	if l, err = s.Listen(nil, udpAddr, nil); err != nil {
		fmt.Println(err)
	}

	//start accepting connections on this listener
	l.Accept(nil)

}
