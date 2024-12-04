package quic

import (
	"net"
)

// type TrConfig struct {
// 	zeroRTT bool
// 	timeOut uint //in seconds
// }

type Transport struct {
	Connections []*Connection
	Conn        net.UDPConn
}

type Listener struct {
}

func (l *Listener) Accept() (*Connection, error) {
	panic("Accept not implemented yet!")
}

func (t *Transport) Dial(raddr *net.UDPAddr, laddr *net.UDPAddr) (*Connection, error) {
	//open a QUIC connection
	conn, err := OpenConnection(raddr, laddr)
	if err != nil {
		return nil, err
	}

	//add the connectio to the transport
	t.Connections = append(t.Connections, conn)
	return conn, nil
}

func (t *Transport) Listen(laddr *net.UDPAddr) *Listener {
	return nil
}

func LIsten(laddr *net.UDPAddr) *Listener {
	tr := Transport{}

	return tr.Listen(laddr)

}
