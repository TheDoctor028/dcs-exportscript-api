package DCS

import (
	"github.com/thedoctor028/dcsexportscriptapi/api"
	udpConnection "github.com/thedoctor028/dcsexportscriptapi/udp-connection"
	"log"
	"net"
	"net/http"
	"os"
	"strconv"
	"time"
)

var dcsClientLogger = log.New(os.Stdout, "DCS Service: ", 101)
var dataLogger, loggerFile = initDataLogger()

const (
	WebsocketRaw api.WSName = "raw"
)

type Service struct {
	ExportIp              string // The host where the UDP server is listening Config.lua/ExportScript.Config.IkarusHost (Used by UDPServer)
	ReceiverIp            string // The host where DCS is running (Used by UDPClient)
	ExportPort            int    // Export port where the export script sends data to, see Config.lua/ExportScript.Config.IkarusPort (Used by UDPServer)
	ReceiverPort          int    // Listening port of the export script see Config.lua/ExportScript.Config.ListenerPort
	ReceiverListeningPort int    // Port for the UDPClient to send data from (Used by UDPClient)
	APIIp                 string // Ip for the http/ws api server
	APIPort               int    // Port for the http/ws api server
	Path                  string // Path to the DCS (game) main directory example: C:\ProgramFiles\DCS World TODO NOT USED YET
	PathSavedGames        string // Path to the DCS (game) save game directory example: C:\Users\user\Saved Games\DCS TODO NOT USED YET

	udpServer *udpConnection.UDPServer // Wrapped UDP Socket used to receive data from export script
	udpClient *udpConnection.UDPClient // Wrapped UDP Socket used to send data to export script
	api       *api.API                 // Wrapped WS and HTTP server to send/revise data from/to the web client
}

// NewService
// Returns a new DCS.Service struct
func NewService() *Service {
	return &Service{
		ExportIp:              "127.0.0.1",
		ReceiverIp:            "127.0.0.1",
		ExportPort:            1625,
		ReceiverPort:          1626,
		ReceiverListeningPort: 1627,
		APIIp:                 "127.0.0.1",
		APIPort:               8000,
		Path:                  "C:\\Program Files\\DCS World",
		PathSavedGames:        "C:\\Users\\user\\DCS",
		udpServer:             nil,
		udpClient:             nil,
		api:                   nil,
	}
}

func (c *Service) CreateAndStartConnections() error {
	var err error
	// UDP
	c.udpClient, err = udpConnection.NewUDPClient(c.ReceiverListeningPort)
	c.udpServer, err = udpConnection.NewUDPServer(net.UDPAddr{IP: net.ParseIP(c.ExportIp), Port: c.ExportPort})
	c.initUDPServer()
	c.udpServer.Serve()

	// API
	c.api = api.NewAPI(c.APIIp, c.APIPort)
	c.setUpApiRoutes()
	go c.api.Serve()

	if err != nil {
		dcsClientLogger.Printf("Error creating DCS Service instance: %s\n", err)
		return err
	}
	return nil
}

func (c *Service) Destroy() {
	defer loggerFile.Close()
}

func (c *Service) setUpApiRoutes() {
	c.initWebSockets()

	c.api.Router.AddRoute(api.Route{Path: "/hello", Handler: func(w http.ResponseWriter, r *http.Request) {
		//  Route for network exploration if the user is not familiar with the ip of the dcs-service host
		w.WriteHeader(200)
		defer w.Write([]byte("Hello!"))
	}}) // HELLO

	// Route for getting all the data raw from the UDP socket
	c.api.Router.AddRoute(api.Route{Path: "/raw", Handler: c.api.Websockets[WebsocketRaw].GetHandler()}) // RAW
}

func (c *Service) initWebSockets() {
	c.api.AddWS(WebsocketRaw, c.initRawRouteWS())
}

func (c *Service) initUDPServer() {
	c.udpServer.CB = func(buffer *[]byte, remoteAddr *net.UDPAddr) {
		res := ExtractUIDAndValue(string(*buffer), ":")
		dataLogger.Println(res.ToString())

		// TODO Implement new logic here
		dataScreenData := res.GetDataByUid(50)
		if dataScreenData != nil {
			c.api.Websockets[WebsocketRaw].SendToAllConnections(*dataScreenData)
		}
	}
}

func initDataLogger() (*log.Logger, *os.File) {
	f, _ := os.OpenFile("./logs/data.logs-"+strconv.FormatInt(time.Now().Unix(), 10)+".txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0600)
	return log.New(f, " //////\n ", 101), f
}
