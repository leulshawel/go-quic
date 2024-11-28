package transport

import (
	"go-quic/streams"
)

func (p *Packet) Go(s streams.Stream) (int, error) {
	//check if the stream can send this packet

	//send the packet
	return 0, nil
}

func (p *Packet) Resend(s streams.Stream) (int, error) {
	//anylise why resend needed

	//check if can be resent

	//resend
	return p.Go(s)

}

func (p *Packet) CanBeSent(s streams.Stream) bool {
	//return true if can be sent of the stream s
	return false
}
