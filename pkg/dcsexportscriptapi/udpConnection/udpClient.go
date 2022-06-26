package udpConnection

import (
	"fmt"
	"net"
)

func SendDataToUDPServer(addr string, port int, data string) {
	conn, err := net.Dial("udp", "127.0.0.1:26027")

	if err != nil {
		fmt.Printf("Some error %v\n", err)
		return
	}
	_, err = fmt.Fprintf(conn, data)

	if err != nil {
		fmt.Printf("Can't send command via UDP error: ", err)
	}
}
