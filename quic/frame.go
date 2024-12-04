package quic

//List of frames

//TODO:
//PING   ping (mainly to keep the connection fro m silently closing)
//CRYPTO used to transmit handshake
//STREAM  a frame with actual application data (has a fin bit to be set when all data is sent)

//NEW_CONNECTION_ID used to send a new connection id
//RETIRE_CONNECTION_ID  remove a connection id
//RESET_STREAM   close the stream (abrupt termination) if not already in terminal state
//STOP_SENDING   abort receipt and
//MAX_STREAM_DATA	the maximum number of bytes that a stream can transport
//MAX_DATA   		the maximum number of bytes of data that can be exchanges in a connection
//MAX_STREAM        maximum number of streams per connection (>2^60) -> TRANSPORT_PARAMET_ERROR/FRAME_ENCODING_ERROR
//STREAMS_BLOCKED	same as SEND_DATA_BLOCKED but for streams
//NEW_CONNECTION_ID		adding a new connection id
//CONNECTION_CLOSE       close a connection
//RETIRE_CONNECTION_ID	retiering a connection id

type Frame struct {
	FrameType int
}

func (f Frame) IsOfType(frameType int) bool { return f.FrameType == frameType }
