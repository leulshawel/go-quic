package streams

const (
	READY = iota
	SEND
	DATA_SENT
	RESET_SENT
	DATA_RECVD
	RESET_RECVD
)

type Stream struct {
	Id    uint64 //id 64 integer (0x01 server/client 0x02 uni/bi directional)
	State uint   //current state of the stream
}

func (s *Stream) Send(d []byte) {
	//check if data can be sent from the current state
	if s.State == READY {
		//send the STREAM
	} else {
		//return an ERROR in case data can't
	}

}
