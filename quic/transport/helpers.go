package transport

const (
	STREAM = iota
	STREAM_DATA_BLOCKED
	RESET_STREAM
	FIN
	MAX_STREAM_DATA
	ALL_ACK
	APP_READ_ALLDATA
	APP_READ_RESET

	FRAME_CONST_LIMIT //used for safety (boundary)
)
