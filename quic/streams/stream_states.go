package streams

// a state manager for managing the state of streams
type StateManager struct {
	state  int
	stream *Stream
}

// Sending
// newly created ready to accept from application
// if bidirectional wait until the stream is created on the receiving to transition to this state
func (s *StateManager) toReady() error {
	if s.state == STREAM_STATE_READY {
		panic(("trying to transion from READY to READY"))
	}

	//validate all conditions for transition are met
	if s.stream.getType() == STREAM_TYPE_CLIENT_BI || s.stream.getType() == STREAM_TYPE_SERVER_BI {
		//wait (may be in a little goroutin)
	}
	s.state = STREAM_STATE_READY
	return nil
}

// sending the first SREAM or STREAM_DATA_BLOCKED puts us in this state
// we can trnasmit and retransmit as neccessary bu respecting flow_control_limit
// we can also receive and process MAX_STREAM_DATA
func (s *StateManager) toSend() error {
	if s.state == STREAM_STATE_SEND {
		panic("trying to transition from SEND to SEND state")
	}
	//validate the transition
	//do the transition to send state
	s.state = STREAM_STATE_SEND
	return nil
}

// transition after sendinf FIN bit set stream (can only restransmit from this state)
// retransmit as needed, no need to check flow_control_limit,
// no need to send STEAM_DATA_BLOCKED
// can safely ignore any MAX_STREAM_DATA frames
func (s *StateManager) toDataSent() error {
	if s.state == STREAM_STATE_DATA_SENT {
		panic("trying to transition from DATA_SENT to DATA_SENT state")
	}
	//validate the transition
	//do the transition to send state
	s.state = STREAM_STATE_DATA_SENT
	return nil
}

// Terminal state
func (s *StateManager) toResetRecvdSending() error {
	if s.state == STREAM_STATE_RESET_RECVD_SENDING {
		panic("trying to transition from RESET_RECVD to RESET_RECVD state")
	}
	//validate the transition
	//do the transition to send state
	s.state = STREAM_STATE_RESET_RECVD_SENDING
	return nil
}

// transition to this state only happens once an ACK is recieved for all stream data
func (s *StateManager) toDataRecvd() error {
	if s.state == STREAM_STATE_DATA_RECVD_SENDING {
		panic("trying to transition from DATA_RECVD to DATA_RECVD state")
	}
	//validate the transition
	//do the transition to send state
	s.state = STREAM_STATE_DATA_RECVD_SENDING
	return nil
}

func (s *StateManager) toResetSent() error {
	if s.state == STREAM_STATE_RESET_SENT {
		panic("trying to transition from RESET_SENT to RESET_SENT state")
	}
	//validate the transition
	//do the transition to send state
	s.state = STREAM_STATE_RESET_SENT
	return nil
}

// Recieving part
// when a a recieving part is created
// can recieve data till MAX_STREAM_DATA
func (s *StateManager) toRecv() error {
	if s.state == STREAM_STATE_RECV {
		panic("trying to transition from RECV to RECV state")
	}
	return nil
}

// application reads All data
func (s *StateManager) toDataRead() error {
	if s.state == STREAM_STATE_DATA_READ {
		panic("trying to transition from READ to READ state")
	}
	return nil
}

func (s *StateManager) toSizeKnown() error {
	if s.state == STREAM_STATE_RECV {
		panic("trying to transition from STATE_KNOWN to STATE_KNOWN state")
	}
	return nil
}

// Reset recieved
func (s *StateManager) toResetRecvdRecving() error {
	if s.state == STREAM_STATE_RESET_RECVD_RECVING {
		panic("trying to transition from RESET_RECVD to RESET_RECVD state")
	}
	//validate the transition
	//do the transition to send state
	s.state = STREAM_STATE_RESET_RECVD_RECVING
	return nil
}

// Reciet Read
func (s *StateManager) toResetRead() error {
	if s.state == STREAM_STATE_RESET_READ {
		panic("trying to transition from RESET_RECVD to RESET_RECVD state")
	}
	//validate the transition
	//do the transition to send state
	s.state = STREAM_STATE_RESET_READ
	return nil
}
