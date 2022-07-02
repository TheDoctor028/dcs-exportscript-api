package udpConnection

import (
	"fmt"
	"log"
	"net"
)

var SenderConn *net.UDPConn

func ServeUDPSender(port int) {
	var err error
	SenderConn, err = net.ListenUDP("udp", &net.UDPAddr{Port: port})

	if err != nil {
		log.Fatal("Listen:", err)
	}
}

func SendDataToUDPServer(conn *net.UDPConn, addr *net.UDPAddr, data string) {
	n, err := conn.WriteTo([]byte(data), addr)
	if err != nil {
		fmt.Printf("Can't send command via UDP error: ", err)
	}

	fmt.Println("Sent", n, "bytes", conn.LocalAddr(), "->", addr)
}
