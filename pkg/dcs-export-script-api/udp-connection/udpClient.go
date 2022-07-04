package udpConnection

import (
	"fmt"
	"net"
)

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
	return UDPSender{
		Conn: conn,
	}, err
}

// SendData
// Send the given string to the Target addr from the created connection
func (udp UDPSender) SendData(data string) error {
	n, err := udp.Conn.WriteTo([]byte(data), &udp.Target)
	if err != nil {
		fmt.Printf("Can't send command via UDP error: %s", err)
		return err
	}

	fmt.Printf("Sent %d bytes from %s -> %s\n", n, udp.Conn.LocalAddr(), udp.Target)
	return nil
}
