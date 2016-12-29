package main

import (
	"fmt"
	"net"
	"os"
	"protocol"
	"strconv"
	"time"
)

func sender1(conn net.Conn) {
	words := "hello world!"
	conn.Write([]byte(words))
	fmt.Println("send over")

}

func send2(conn net.Conn) {
	for i := 0; i < 100; i++ {
		session := GetSession()
		words := "{\"ID\":" + strconv.Itoa(i) + "\",\"Session\":" + session + "2015073109532345\",\"Meta\":\"golang\",\"Content\":\"message\"}"
		//conn.Write(protocol.Enpacket([]byte(words)))
		conn.Write(protocol.Enpack([]byte(words)))
	}
	fmt.Println("send over")
	defer conn.Close()
}

func senderLong(conn net.Conn) {
	for i := 0; i < 5; i++ {
		words := strconv.Itoa(i) + "testlong"
		conn.Write(protocol.Enpack([]byte(words)))
		time.Sleep(2 * time.Second)
	}
	fmt.Println("send over")
}

func GetSession() string {
	gs1 := time.Now().Unix()
	gs2 := strconv.FormatInt(gs1, 10)
	return gs2
}

func main() {
	server := "127.0.0.1:1024"
	tcpAddr, err := net.ResolveTCPAddr("tcp4", server)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}

	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}

	fmt.Println("connect success")
	//sender1(conn)
	//send2(conn)
	senderLong(conn)
}
