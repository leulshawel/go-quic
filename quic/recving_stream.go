package quic

type RecvingEnd struct {
	StateManager     *StateManager //manage state of the sream from here
	FlowCotrolLimit  int           //dont't send packate of it exids this
	connection       *Connection   //the connection this stream belongs to
	lastUnreadPacket *Packet
	MaxStreamData    uint //maximum aount of data can be sent in this stream (sent/received via flow_control_limit)
}

// Receive data
func (s *Stream) Receive() int {
	var byteRecvd int
	//check if the stream already exists or is newly created from the sending side
	if s.isInConnection(s.connection) {
		//data from an exsting stream
	} else {
		//this are the only frame types that change states on in a recving/sending stream
		stream_ := s.lastUnreadPacket.Frame.IsOfType(STREAM)
		stream_data_blocked_ := s.lastUnreadPacket.Frame.IsOfType(STREAM_DATA_BLOCKED)
		reset_stream_ := s.lastUnreadPacket.Frame.IsOfType(RESET_STREAM)

		if stream_ || stream_data_blocked_ || reset_stream_ {
			//handle the state change

		}
	}
	return byteRecvd

}

// when an application aborts we send a STOP_SENDING frame
// if the receiving part in the RECV or SIZE_KNOWN state
func (s *Stream) Abort(erro_code int) {
	if s.RecvingEnd.StateManager.isTerminated() {
		panic("Aborting a terminated stream")
	}
	//send a packet with a STOP_SENDING frame
	s.Send(Packet{})
}

func (r RecvingEnd) isTerminated() bool {
	return r.StateManager.isTerminated()
}
