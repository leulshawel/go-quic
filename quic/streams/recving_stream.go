package streams

import (
	"go-quic/quic/connections"
	"go-quic/quic/transport"
)

type RecvingEnd struct {
	StateManager     *StateManager           //manage state of the sream from here
	FlowCotrolLimit  int                     //dont't send packate of it exids this
	connection       *connections.Connection //the connection this stream belongs to
	lastUnreadPacket *transport.Packet
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
		stream_ := s.lastUnreadPacket.Frame.IsOfType(transport.STREAM)
		stream_data_blocked_ := s.lastUnreadPacket.Frame.IsOfType(transport.STREAM_DATA_BLOCKED)
		reset_stream_ := s.lastUnreadPacket.Frame.IsOfType(transport.RESET_STREAM)

		if stream_ || stream_data_blocked_ || reset_stream_ {
			//handle the state chane

		}
	}
	return byteRecvd

}
