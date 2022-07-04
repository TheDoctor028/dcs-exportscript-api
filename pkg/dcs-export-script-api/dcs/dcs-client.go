package DCS

import (
	udpConnection "github.com/thedoctor028/dcsexportscriptapi/udp-connection"
)

type Client struct {
	ExportIp       string // The host where the UDP server is listening Config.lua/ExportScript.Config.IkarusHost (This is what our UDP server uses)
	ReceiverIp     string // The host where DCS is running
	ExportPort     int    // Export port of the export script see Config.lua/ExportScript.Config.IkarusPort (This is what our UDP server uses)
	ReceiverPort   int    // Listening port of the export script see Config.lua/ExportScript.Config.ListenerPort
	Path           string // Path to the DCS (game) main directory example: C:\ProgramFiles\DCS World TODO
	PathSavedGames string // Path to the DCS (game) save game directory example: C:\Users\user\Saved Games\DCS TODO

	udpServer *udpConnection.UDPServer // Wrapped UDP Socket used to receive data from export script
	udpClient *udpConnection.UDPClient // Wrapped UDP Socket used to send data to export script
}

func NewClient() *Client {
	return &Client{
		ExportIp:       "127.0.0.1",
		ReceiverIp:     "127.0.0.1",
		ExportPort:     1625,
		ReceiverPort:   1626,
		Path:           "C:\\Program Files\\DCS World",
		PathSavedGames: "C:\\Users\\user\\DCS",
		udpServer:      nil,
		udpClient:      nil,
	}
}

func (c Client) CreateAndStartUdpConnections() error {
	var err error
	//c.udpClient, err = udpConnection.ServeUDPSender(c.ExportPort)
	//c.udpServer, err = udpConnection.ServeUDPServer(c.ExportPort)

	if err != nil {
		return err
	}
	return nil
}
