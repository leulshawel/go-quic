package streams

import (
	"go-quic/quic/transport"
)

type SendindEnd struct {
	Id              StreamId
	StateManager    *StateManager //manage state of the sream from here
	FlowCotrolLimit int           //dont't send packate of it exids this
	MaxStreamData   uint          //maximum aount of data can be sent in this stream (sent/received via flow_control_limit)
}

// sending  end
func (s SendindEnd) Send(p transport.Packet) int {
	//we can also delay giving id to a stream till it sends its first stream
	var byteSent int
	//check packate is within the flow control set by peer
	//transition to SEND state
	if s.StateManager.state == STREAM_STATE_SEND {
		return byteSent
	}

	if p.Frame.FrameType == transport.STREAM_DATA_BLOCKED ||
		p.Frame.FrameType == transport.STREAM {

		if err := s.StateManager.toSend(); err != nil {
			//handle the error
		}

		if p.Frame.FIN != 0 {
			if err := s.StateManager.toDataSent(); err != nil {
				//handle
			}
		}
	}

	//if we are blocked from sending by the flow_control_limit then we send STREAM_SEND_BLOCKED

	return byteSent
}

func (s SendindEnd) End(p transport.Packet) int {
	var byteSent int
	//send a STREAM + FIN frame
	//trnasition to stream state data sent
	if s.StateManager.state == STREAM_STATE_DATA_SENT {
		return byteSent
	}

	//wait for all ack
	if err := s.StateManager.toDataSent(); err != nil {
		//handle the error
	}
	return byteSent
}

func (s SendindEnd) Reset() {
	//check if stream is not in terminal state
	//send a RESET_STREAM frame
	//transition to RESET_SENT when recieve an ACK
	if s.StateManager.state == STREAM_STATE_SEND {
		return
	}
	if err := s.StateManager.toResetSent(); err != nil {
		//handle the error
	}

}
