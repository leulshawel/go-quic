package quic

const (
	TRANSPORT_PARAMETER_ERROR = iota
	FRAME_ENCODING_ERROR
	STREAM_LIMIT_ERROR
	CONNECTION_ID_LIMIT_ERROR
	CONNECTION_REFUSED //refuse an incomming connection

)
