package quic

import (
// "go-quic/quic/api"
// "go-quic/quic/connections"
)

type Packet struct {
	DestConnectionId uint
	Frame            *Frame
}
