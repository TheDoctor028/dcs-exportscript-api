package udpConnection

import (
	"fmt"
	"log"
	"net"
	"os"
)

var clientLogger = log.New(os.Stdout, "UDP Client: ", 101)

var SenderConn *net.UDPConn

type UDPSender struct {
	Conn   *net.UDPConn
	Target net.UDPAddr
}

func ServeUDPSender(port int) (*net.UDPConn, error) {
	return net.ListenUDP("udp", &net.UDPAddr{
		Port: port})
}

func SendDataToUDPServer(conn *net.UDPConn, addr *net.UDPAddr, data string) {
	n, err := conn.WriteTo([]byte(data), addr)
	if err != nil {
		fmt.Printf("Can't send command via UDP error: ", err)
	}

	fmt.Println("Sent", n, "bytes", conn.LocalAddr(), "->", addr)
}

// NewUDPSender
// Start a UDP listener on the given port.
// The Target must be set later.
func NewUDPSender(port int) (UDPSender, error) {
	conn, err := net.ListenUDP("udp", &net.UDPAddr{
		Port: port})

	if err == nil {
		clientLogger.Printf("Listening on UDP %s...\n", conn.LocalAddr())
	}

	return UDPSender{
		Conn: conn,
	}, err
}

// SendData
// Send the given string to the Target addr from the created connection
func (udp UDPSender) SendData(data string) error {
	n, err := udp.Conn.WriteTo([]byte(data), &udp.Target)
	if err != nil {
		clientLogger.Printf("Can't send command via UDP error: %s", err)
		return err
	}

	clientLogger.Printf("Sent %d bytes from %s -> %s\n", n, udp.Conn.LocalAddr(), udp.Target)
	return nil
}
