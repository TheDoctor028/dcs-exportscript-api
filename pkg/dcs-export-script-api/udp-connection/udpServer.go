package udpConnection

import (
	"log"
	"net"
	"os"
)

// serverLogger
// Logger for logging server events
var serverLogger = log.New(os.Stdout, "UDP Server: ", 101)

// UDPServer
// Wrapped UDP socket to receive UDP traffic
type UDPServer struct {
	Conn *net.UDPConn                                  // UDP socket
	Addr net.UDPAddr                                   // Host address
	CB   func(buffer *[]byte, remoteAddr *net.UDPAddr) // Callback function that handles the incoming data

	buffer []byte // Buffer to receive data to
}

// sendResponse
// A simple ack answer to the sender
func sendResponse(conn *net.UDPConn, addr *net.UDPAddr) {
	_, err := conn.WriteToUDP([]byte("ack"), addr)
	if err != nil {
		serverLogger.Printf("Couldn't send response %v", err)
	}
}

// Serve
// Start to process the incoming data
func (s *UDPServer) Serve() {
	for {
		_, remoteaddr, err := s.Conn.ReadFromUDP(s.buffer)

		if err != nil {
			serverLogger.Printf("Some error  %v", err)
			continue
		}
		go s.CB(&s.buffer, remoteaddr)
		go sendResponse(s.Conn, remoteaddr)
	}
}

// NewUDPServer creates a new UDPServer instance
// with 1024 size buffer with an empty logging callBack
func NewUDPServer(addr net.UDPAddr) (UDPServer, error) {
	server := UDPServer{
		Addr: addr,
		CB: func(buffer *[]byte, remoteAddr *net.UDPAddr) {
			serverLogger.Printf("Data %d bytes received from %s", len(*buffer), remoteAddr)
		},
		buffer: make([]byte, 1024),
	}

	conn, err := net.ListenUDP("udp", &addr)

	if err == nil {
		serverLogger.Printf("Listening on UDP %s:%d...", addr.IP, addr.Port)
		server.Conn = conn
	}

	return server, err
}
