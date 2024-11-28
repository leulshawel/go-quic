package connections

type Initial struct {
	HFLRP         uint8
	Version       int
	DestConnIdLen uint8
	DestConnId    []byte
	SrcConnIdLen  uint8
	SrcConnId     []byte
	//...
}
