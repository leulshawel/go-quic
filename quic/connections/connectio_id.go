package connections

//conections id are to identify a connection with an endpoint
//multiple connection ids can be used per connection
//connectio ids have an associated sequence number
//sequence number of initial connection is is 0
//sequence numbers increament by 1 for every new connection id issued

type ConnectionId struct {
	Id             [20]byte
	SequenceNumber int
}

func Generate_connection_id() ConnectionId {
	var id ConnectionId
	//generate a connetion id following the rules

	return id
}
