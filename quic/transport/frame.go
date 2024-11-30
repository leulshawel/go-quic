package transport

//List of frames

//TODO:
//STREAM  a frame with actual application data (has a fin bit to be set when all data is sent)

//NEW_CONNECTION_ID used to send a new connection id
//RETIRE_CONNECTION_ID  remove a connection id
//RESET_STREAM   close the stream (abrupt termination) if not already in terminal state
//STOP_SENDING   abort receipt and
//MAX_STREAM_DATA

type Frame struct {
	FrameType int
	FIN       int8
}

func (f Frame) IsOfType(frameType int) bool { return f.FrameType == frameType }
