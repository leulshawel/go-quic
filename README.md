<b><h1>go-quic</h1></b> go-quic is an implementation of the quic protocol [rfc9000](https://datatracker.ietf.org/doc/html/rfc9000) mainly for educational and small application uses<br> 


<h3><b>The plan</b></h3>
Our plan here is to build a full end to end implementation for the protocol and to build a small http3 server on top<br>

<h3><b>Listener</b></h3>
A Listener is a quic protocol running on a specific quic port

```go

udpAddr := &net.UDPAddr{IP: net.IPv4(127, 0, 0, 1), Port: 8080}


l, err := quic.CreateNewListener(nil, udpAddr, nil, nil, nil, nil, 0, 0, 0, nil);
if err_ := l.Listen(nil); err_ != nil {
	fmt.Println(err)
	return;
}

l.Accept(nil)
return

``` 

<h3><b>Server</b></h3>

A server in go-quic is the central place from where we manage quic protocol listeners running on different udp ports

```go
//udp address to listen on
udpAddr := &net.UDPAddr{IP: net.IPv4(127, 0, 0, 1), Port: 8080}
    
s := quic.CreateNewServer(context.Background()) //create a server
    
var l *quic.Listener; var err error
if l, err = s.Listen(nil, udpAddr, nil); err != nil {
    fmt.Println(err)
}

//start accepting connections on this listener
l.Accept(nil)
```

<h3><b>Contributions are welcomed</b></h3>
