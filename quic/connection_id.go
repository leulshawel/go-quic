package quic

//conections id are to identify a connection with an endpoint
//multiple connection ids can be used per connection
//connection ids have an associated sequence number
//sequence number of initial connection is is 0
//sequence numbers increament by 1 for every new connection id issued

type ConnectionId [20]byte

type ConnectionIdManager struct {
	ActiveConnectionIdLimit int
	SequenceNumber          int //start from 0 and increament by 1 for every new id/connection
	Retired                 bool
	IdPool                  []ConnectionId
}

func (cm *ConnectionIdManager) generateId() ConnectionId {
	return ConnectionId{}
}

func (cm *ConnectionIdManager) retierId() error {
	//remove the id from the pool
	return nil
}

func (cm *ConnectionIdManager) Add(id ConnectionId) error {
	cm.IdPool = append(cm.IdPool, id)
	return nil
}

func (c *Connection) AddNewId() ConnectionId {
	id := c.ConnectionIdManager.generateId()
	c.ConnectionIdManager.Add(id)
	return id
}
