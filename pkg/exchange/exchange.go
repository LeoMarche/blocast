package exchange

import "net"

var DEFAULT_IPV6 = "[::1]"
var DEFAULT_IPV4 = "127.0.0.1"

var DEFAULT_SERVER_PORT = "23654"

var CONN_LIST []net.Conn

func InitializeList() error {
	var conn net.Conn
	var err error
	conn, err = net.Dial("tcp", DEFAULT_IPV6+":"+DEFAULT_SERVER_PORT)
	if err != nil {
		conn, err = net.Dial("tcp", DEFAULT_IPV4+":"+DEFAULT_SERVER_PORT)
		if err != nil {
			return err
		}
	}
	CONN_LIST = append(CONN_LIST, conn)
	return nil
}

func Broadcast(b []byte) error {
	//TODO implement the function to broadcast a message

	return nil
}

func Receive() ([]byte, error) {
	//TODO implement the function that receives a message

	return nil, nil
}
