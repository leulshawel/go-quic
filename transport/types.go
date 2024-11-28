package transport

import (
	"go-quic/streams"
)

type Packet_ interface {
	Go(s *streams.Stream)
	Resend(s *streams.Stream)
	CanBeSent(s *streams.Stream)
}

type Packet struct {
	PacketId int
}
