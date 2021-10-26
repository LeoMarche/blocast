package exchange

import (
	"net"
)

var DEFAULT_IPV6 = "[::1]"
var DEFAULT_IPV4 = "127.0.0.1"

var DEFAULT_SERVER_PORT = "23654"

var CONN_LIST []net.Conn
var RES_LIST [][]byte
var SEND_LIST []byte

//Add struct sendQ which is an array of bytes with a mutex

//Add struct resQ which is an array of array of bytes with a mutex

//Add connPurge who, given a list of indexs in error, removes the connection and the corresponding res string

//Add parser who, given a byte array, tries to find a valid communication and gives it back to blocast-core for processing and returns the bytes to remove at the start of the string
func Parse(b []byte) ([]byte, int) {
	//	_, _ := regexp.Compile("(\\D*)(\\d*)(.*)")
	return nil, 0

}

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

//Function Broadcast broadcasts a message to a list of conn, reporting any errors to the caller
func Broadcast(b []byte) ([]int, []error) {
	var reti []int
	var rete []error
	for i, c := range CONN_LIST {
		_, err := c.Write(b)
		if err != nil {
			reti = append(reti, i)
			rete = append(rete, err)
		}
	}
	return reti, rete
}

func Receive() ([][]byte, []error) {

	var resb [][]byte
	var rese []error

	for _, c := range CONN_LIST {
		b, err := ReceiveConn(c)
		resb = append(resb, b)
		rese = append(rese, err)
	}

	return resb, rese
}

func ReceiveConn(c net.Conn) ([]byte, error) {
	var b []byte
	_, err := c.Read(b)
	return b, err
}
