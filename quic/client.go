package quic

import (
	"net"
)

type Client struct {
}

func (c *Client) Dial(raddr *net.UDPAddr, laddr *net.UDPAddr) (*Connection, error) {
	//open a QUIC connection
	panic("Client.dial not implemented yet")
}
