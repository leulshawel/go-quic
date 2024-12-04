package quic

type SendindEnd struct {
	StateManager    *StateManager //manage state of the sream from here
	FlowCotrolLimit int           //dont't send packate of it exeeds this
	MaxStreamData   uint          //maximum aount of data can be sent in this stream (sent/received via flow_control_limit)
}

// sending  end
func (s *Stream) Send(p Packet) int {
	sendingEnd := s.SendindEnd
	//we can also delay giving id to a stream till it sends its first stream
	var byteSent int
	//check packate is within the flow control set by peer
	//transition to SEND state

	newState := sendingEnd.StateManager.getNextState((p.Frame.FrameType))
	sendingEnd.StateManager.state = newState

	//if we are blocked from sending by the flow_control_limit then we send STREAM_SEND_BLOCKED

	return byteSent
}

func (s *Stream) End(p Packet) int {
	sendingEnd := s.SendindEnd
	var byteSent int
	//send a STREAM + FIN frame
	//trnasition to stream state data sent
	if sendingEnd.StateManager.state == STREAM_STATE_DATA_SENT {
		return byteSent
	}

	newState := sendingEnd.StateManager.getNextState(FIN)
	sendingEnd.StateManager.state = newState
	return byteSent
}

func (s *Stream) Reset() {
	sendingEnd := s.SendindEnd
	//check if stream is not in terminal state
	//send a RESET_STREAM frame
	//transition to RESET_SENT when recieve an ACK

	newState := sendingEnd.StateManager.getNextState(RESET_STREAM)
	sendingEnd.StateManager.state = newState

}

func (s SendindEnd) isTerminated() bool {
	return s.StateManager.isTerminated()
}
