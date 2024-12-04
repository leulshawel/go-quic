package quic

import (
	"errors"
)

type StreamId uint64

// A stream
type Stream struct {
	SendindEnd
	RecvingEnd
	Id                  uint64 //id 64 integer (0x01 server/client 0x02 uni/bi directional) cannot be reused for the same connection
	PriorityLevel       int    //from the application
	MaxStreamData       int    //max stream data that can be sent over the stream
	DataCreditRemaining int    //remaining amount of data
	connectionIdx       int    //the index of this stream in the Streams slice in Connection
}

// return the first two bits of the id (they represent the stream type)
func (s *Stream) getType() uint8 { return uint8(s.Id) & uint8(3) }

// check if a stream belogs to a connection
func (s *Stream) isInConnection(con *Connection) bool { return false }

func (s *Stream) isClientInitiated() bool { return (s.Id % 2) == 0 }

func CreateStream(conn *Connection, stream_type uint8) (*Stream, error) {
	//check if a connection has the required MAX_STREAM_DATA available

	s := Stream{}

	if stream_type == STREAM_TYPE_CLIENT_UNI || stream_type == STREAM_TYPE_SERVER_UNI {

	} else {
		return nil, errors.New("can't create bidirectionsal streams")
	}
	//initiate the sending end
	// recievingEnd := &s.RecvingEnd

	//initiate the sendning end
	// sendingEnd = &s.SendingEnd

	return &s, nil
}

func (s *Stream) AddToConnection(c *Connection) {
	//validate the stream
	c.Streams = append(c.Streams, s)
}

//When terminating a stream the two sides always have to agree on how much flow control credit was consumed
//final size of a stream can't chanfge STREAM/RESET_STREAM frame indicating a change -> FINAL_SIZE_ERROR.
//data at or beyond the final size -> FINAL_SIZE_ERROR

//CONCURENCY
//possible values of stream id (max_streams * 4 + first_stream_of_type)
//Initial limits are sent in the transport parameters
