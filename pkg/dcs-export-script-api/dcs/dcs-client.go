package DCS

import (
	"github.com/thedoctor028/dcsexportscriptapi/api"
	udpConnection "github.com/thedoctor028/dcsexportscriptapi/udp-connection"
	"log"
	"net"
	"os"
)

var dcsClientLogger = log.New(os.Stdout, "DCS Client: ", 101)

type Client struct {
	ExportIp              string // The host where the UDP server is listening Config.lua/ExportScript.Config.IkarusHost (Used by UDPServer)
	ReceiverIp            string // The host where DCS is running (Used by UDPClient)
	ExportPort            int    // Export port where the export script sends data to, see Config.lua/ExportScript.Config.IkarusPort (Used by UDPServer)
	ReceiverPort          int    // Listening port of the export script see Config.lua/ExportScript.Config.ListenerPort
	ReceiverListeningPort int    // Listening port for the UDPClient (Used by UDPClient)
	APIIp                 string // Ip for the http/ws api server
	APIPort               int    // Port for the http/ws api server
	Path                  string // Path to the DCS (game) main directory example: C:\ProgramFiles\DCS World TODO
	PathSavedGames        string // Path to the DCS (game) save game directory example: C:\Users\user\Saved Games\DCS TODO

	udpServer *udpConnection.UDPServer // Wrapped UDP Socket used to receive data from export script
	udpClient *udpConnection.UDPClient // Wrapped UDP Socket used to send data to export script
	api       *api.API                 // Wrapped WS and HTTP server to send/revise data from/to the web client
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
		api:            nil,
	}
}

func (c Client) CreateAndStartUdpConnections() error {
	var err error
	c.udpClient, err = udpConnection.NewUDPClient(c.ReceiverListeningPort)
	c.udpServer, err = udpConnection.NewUDPServer(net.UDPAddr{IP: net.ParseIP(c.ExportIp), Port: c.ExportPort})
	c.api = api.NewAPI(c.APIIp, c.APIPort)
	if err != nil {
		dcsClientLogger.Printf("Error creating DCS Client instance: %s\n", err)
		return err
	}
	return nil
}
