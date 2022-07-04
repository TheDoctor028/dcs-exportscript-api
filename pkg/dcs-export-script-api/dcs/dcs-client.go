package DCS

import (
	udpConnection "github.com/thedoctor028/dcsexportscriptapi/udp-connection"
)

type Client struct {
	SenderDestinationIp string // The host where the UDP server is listening Config.lua/ExportScript.Config.IkarusHost (This is what our UDP server uses)
	ReceiverIp          string // The host where DCS is running
	SenderPort          int    // Export port of the export script see Config.lua/ExportScript.Config.IkarusPort (This is what our UDP server uses)
	ReceiverPort        int    // Listening port of the export script see Config.lua/ExportScript.Config.ListenerPort
	Path                string // Path to the DCS (game) main directory example: C:\ProgramFiles\DCS World
	PathSavedGames      string // Path to the DCS (game) save game directory example: C:\Users\user\Saved Games\DCS

	udpServer *udpConnection.UDPServer // Wrapped UDP Socket used to receive data from export script
	udpClient *udpConnection.UDPSender // Wrapped UDP Socket used to send data to export script
}

func NewClient() *Client {
	return &Client{
		SenderDestinationIp: "127.0.0.1",
		ReceiverIp:          "127.0.0.1",
		SenderPort:          1625,
		ReceiverPort:        1626,
		Path:                "C:\\Program Files\\DCS World",
		PathSavedGames:      "C:\\Users\\user\\DCS",
		udpServer:           nil,
		udpClient:           nil,
	}
}

func (c Client) CreateAndStartUdpConnections() error {
	var err error
	//c.udpClient, err = udpConnection.ServeUDPSender(c.SenderPort)
	//c.udpServer, err = udpConnection.ServeUDPServer(c.SenderPort)

	if err != nil {
		return err
	}
	return nil
}
