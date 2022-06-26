package udpConnection

import (
	"fmt"
	"log"
	"net"
	"os"
)

var logger = log.New(os.Stdout, "UDP Server: ", 100)

func sendResponse(conn *net.UDPConn, addr *net.UDPAddr) {
	_, err := conn.WriteToUDP([]byte("From server: Hello I got your message "), addr)
	if err != nil {
		logger.Printf("Couldn't send response %v", err)
	}
}

func ServerUDP(address string, port int, buffer *[]byte, cb func()) {
	addr := net.UDPAddr{
		Port: port,
		IP:   net.ParseIP(address),
	}
	ser, err := net.ListenUDP("udp", &addr)
	if err != nil {
		fmt.Printf("Some error %v\n", err)
		return
	}

	logger.Printf("Listening on UDP %s:%d...", addr.IP, addr.Port)

	for {
		_, remoteaddr, err := ser.ReadFromUDP(*buffer)
		//logger.Printf("Read a message from %v \n", remoteaddr)
		if err != nil {
			logger.Printf("Some error  %v", err)
			continue
		}
		cb()
		go sendResponse(ser, remoteaddr)
	}
}
