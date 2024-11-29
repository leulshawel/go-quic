package streams

import (
	"go-quic/quic/connections"
	"go-quic/quic/transport"
)

type RecvingEnd struct {
	Id               StreamId
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
		//new stream creation and state handling
		if s.lastUnreadPacket.Frame.IsOfType(transport.STREAM) ||
			s.lastUnreadPacket.Frame.IsOfType(transport.STREAM_DATA_BLOCKED) {
			s.RecvingEnd.StateManager.toRecv()

		}
	}
	return byteRecvd

}
