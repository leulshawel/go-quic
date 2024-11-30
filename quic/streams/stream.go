package streams

import "go-quic/quic/connections"

type StreamId uint64

// A stream
type Stream struct {
	SendindEnd
	RecvingEnd
	Id            uint64 //id 64 integer (0x01 server/client 0x02 uni/bi directional) cannot be reused for the same connection
	PriorityLevel int    //from the application
}

// return the first two bits of the id (they represent the stream type)
func (s *Stream) getType() uint8 { return uint8(s.Id) & uint8(3) }

// check if a stream belogs to a connection
func (s *Stream) isInConnection(con *connections.Connection) bool { return false }

func (s *Stream) GenerateStreamId(con *connections.Connection) StreamId {
	var id StreamId
	con.GetNextStreamId(s.getType())

	return id
}
