package udpConnection

import (
	"log"
	"net"
	"os"
)

// clientLogger
// Logger for client events
var clientLogger = log.New(os.Stdout, "UDP Client: ", 101)

// UDPClient
// Wrapped UDP socket for sending data to Target (target must be defined after init)
type UDPClient struct {
	Conn   *net.UDPConn // UDP socket
	Target net.UDPAddr  // Address of the destination server, to send the data
}

// NewUDPClient
// Start a UDP listener on the given port.
// The Target must be set later.
// @param {int} port - Listening port for UDP socket
func NewUDPClient(port int) (*UDPClient, error) {
	conn, err := net.ListenUDP("udp", &net.UDPAddr{
		Port: port})

	if err == nil {
		clientLogger.Printf("Listening on UDP %s...\n", conn.LocalAddr())
	}

	return &UDPClient{
		Conn: conn,
	}, err
}

// SendData
// Send the given string to the Target addr from the created connection
// {string} data - String data to send to the server
func (udp UDPClient) SendData(data string) error {
	n, err := udp.Conn.WriteTo([]byte(data), &udp.Target)
	if err != nil {
		clientLogger.Printf("Can't send command via UDP error: %s", err)
		return err
	}

	clientLogger.Printf("Sent %d bytes from %s -> %s\n", n, udp.Conn.LocalAddr(), udp.Target)
	return nil
}
