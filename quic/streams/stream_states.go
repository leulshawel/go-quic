package streams

import (
	"go-quic/quic/transport"
)

// a state manager for managing the state of streams
// map currentState -> posiible states -> next state
var frameToNextstate = map[int]map[int]int{
	//sending states
	STREAM_STATE_READY: {
		transport.STREAM:              STREAM_STATE_SEND,
		transport.STREAM_DATA_BLOCKED: STREAM_STATE_SEND,
		transport.RESET_STREAM:        STREAM_STATE_RESET_SENT,
	},
	STREAM_STATE_SEND: {
		transport.FIN:          STREAM_STATE_DATA_SENT,
		transport.RESET_STREAM: STREAM_STATE_RESET_SENT,
	},
	STREAM_STATE_DATA_SENT: {
		transport.ALL_ACK: STREAM_STATE_DATA_RECVD_SENDING,
	},

	//recieving states
	STREAM_STATE_RECV: {
		transport.FIN:          STREAM_STATE_SIZE_KNOWN,
		transport.RESET_STREAM: STREAM_STATE_RESET_RECVD_RECVING,
	},
	STREAM_STATE_SIZE_KNOWN: {
		transport.FIN:          STREAM_STATE_DATA_RECVD_RECVING,
		transport.RESET_STREAM: STREAM_STATE_RESET_SENT,
	},
	STREAM_STATE_DATA_RECVD_RECVING: {
		transport.APP_READ_ALLDATA: STREAM_STATE_DATA_READ,
	},
	STREAM_STATE_RESET_RECVD_RECVING: {
		transport.APP_READ_RESET: STREAM_STATE_RESET_READ,
	},
}

type StateManager struct {
	state int
}

func (s *StateManager) getNextState(frameType int) int {
	var newState int
	//validate frmatype value before using it to access the transition map
	if frameType < 0 || frameType > transport.FRAME_CONST_LIMIT {
		return -1
	}

	return newState
}

// // Sending
// // newly created ready to accept from application
// // if bidirectional wait until the stream is created on the receiving to transition to this state
// func (s *StateManager) toReady() error {
// 	if s.state == STREAM_STATE_READY {
// 		panic(("trying to transion from READY to READY"))
// 	}

// 	//validate all conditions for transition are met
// 	if s.stream.getType() == STREAM_TYPE_CLIENT_BI || s.stream.getType() == STREAM_TYPE_SERVER_BI {
// 		//wait (may be in a little goroutin)
// 	}
// 	s.state = STREAM_STATE_READY
// 	return nil
// }

// // sending the first SREAM or STREAM_DATA_BLOCKED puts us in this state
// // we can trnasmit and retransmit as neccessary bu respecting flow_control_limit
// // we can also receive and process MAX_STREAM_DATA
// func (s *StateManager) toSend() error {
// 	if s.state == STREAM_STATE_SEND {
// 		panic("trying to transition from SEND to SEND state")
// 	}
// 	//validate the transition
// 	//do the transition to send state
// 	s.state = STREAM_STATE_SEND
// 	return nil
// }

// // transition after sendinf FIN bit set stream (can only restransmit from this state)
// // retransmit as needed, no need to check flow_control_limit,
// // no need to send STEAM_DATA_BLOCKED
// // can safely ignore any MAX_STREAM_DATA frames
// func (s *StateManager) toDataSent() error {
// 	if s.state == STREAM_STATE_DATA_SENT {
// 		panic("trying to transition from DATA_SENT to DATA_SENT state")
// 	}
// 	//validate the transition
// 	//do the transition to send state
// 	s.state = STREAM_STATE_DATA_SENT
// 	return nil
// }

// // Terminal state
// func (s *StateManager) toResetRecvdSending() error {
// 	if s.state == STREAM_STATE_RESET_RECVD_SENDING {
// 		panic("trying to transition from RESET_RECVD to RESET_RECVD state")
// 	}
// 	//validate the transition
// 	//do the transition to send state
// 	s.state = STREAM_STATE_RESET_RECVD_SENDING
// 	return nil
// }

// // transition to this state only happens once an ACK is recieved for all stream data
// func (s *StateManager) toDataRecvd() error {
// 	if s.state == STREAM_STATE_DATA_RECVD_SENDING {
// 		panic("trying to transition from DATA_RECVD to DATA_RECVD state")
// 	}
// 	//validate the transition
// 	//do the transition to send state
// 	s.state = STREAM_STATE_DATA_RECVD_SENDING
// 	return nil
// }

// func (s *StateManager) toResetSent() error {
// 	if s.state == STREAM_STATE_RESET_SENT {
// 		panic("trying to transition from RESET_SENT to RESET_SENT state")
// 	}
// 	//validate the transition
// 	//do the transition to send state
// 	s.state = STREAM_STATE_RESET_SENT
// 	return nil
// }

// // Recieving part
// // when a a recieving part is created
// // can recieve data till MAX_STREAM_DATA
// func (s *StateManager) toRecv() error {
// 	if s.state == STREAM_STATE_RECV {
// 		panic("trying to transition from RECV to RECV state")
// 	}
// 	return nil
// }

// // application reads All data
// func (s *StateManager) toDataRead() error {
// 	if s.state == STREAM_STATE_DATA_READ {
// 		panic("trying to transition from READ to READ state")
// 	}

// 	s.state = STREAM_STATE_DATA_READ
// 	return nil
// }

// func (s *StateManager) toSizeKnown() error {
// 	if s.state == STREAM_STATE_SIZE_KNOWN {
// 		panic("trying to transition from STATE_KNOWN to STATE_KNOWN state")
// 	}

// 	s.state = STREAM_STATE_SIZE_KNOWN
// 	return nil
// }

// // Reset recieved
// func (s *StateManager) toResetRecvdRecving() error {
// 	if s.state == STREAM_STATE_RESET_RECVD_RECVING {
// 		panic("trying to transition from RESET_RECVD to RESET_RECVD state")
// 	}
// 	//validate the transition
// 	//do the transition to send state
// 	s.state = STREAM_STATE_RESET_RECVD_RECVING
// 	return nil
// }

// // Reciet Read
// func (s *StateManager) toResetRead() error {
// 	if s.state == STREAM_STATE_RESET_READ {
// 		panic("trying to transition from RESET_RECVD to RESET_RECVD state")
// 	}
// 	//validate the transition
// 	//do the transition to send state
// 	s.state = STREAM_STATE_RESET_READ
// 	return nil
// }
