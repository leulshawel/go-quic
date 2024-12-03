package main

import (
	"fmt"
	"go-quic/quic/connections"
	"go-quic/quic/streams"
	"net"
)

func main() {

	conn, err := connections.NewConnection(&net.UDPAddr{}, nil)
	stream, err := streams.CreateStream(conn, streams.STREAM_TYPE_CLIENT_UNI)
	if err != nil {
		fmt.Print(err.Error())
	}

	fmt.Println(conn)
	fmt.Println(stream)

}
