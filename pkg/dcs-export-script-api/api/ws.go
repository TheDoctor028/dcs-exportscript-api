package api

import (
	"github.com/gorilla/websocket"
	udpConnection "github.com/thedoctor028/dcsexportscriptapi/udp-connection"
	"log"
	"net/http"
	"strconv"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
} // use default options

var connections = map[int]*websocket.Conn{}
var connId = 0

// LEGACY STUFF
func setUpWSConnection(udpClient *udpConnection.UDPClient) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		connection, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			log.Print("upgrade:", err)
			return
		}
		defer connection.Close()

		connection.WriteMessage(websocket.TextMessage, []byte(strconv.Itoa(connId)))
		connections[connId] = connection
		connId++
		for {
			_, message, err := connection.ReadMessage()
			if err != nil {
				log.Println("read:", err)
				break
			}
			log.Printf("recv: %s", message)

			if string(message)[0] == 'C' {
				err := udpClient.SendData(string(message))
				if err != nil {
					println(err)
				}
			}
		}
	}
}

func SendEventToAllConnections(event string) {
	for id, conn := range connections {
		log.Println("Sending data to: ", id)

		log.Println("Error: ", conn.WriteMessage(websocket.TextMessage, []byte(event)))
	}
}

// END OF LEGACY STUFF

type WS struct {
	Connections map[int]*websocket.Conn     // The currently active connection on the websocket
	Upgrader    *websocket.Upgrader         // Upgrader instance to upgrade HTTP -> WS
	Handler     func(ws *WS) RequestHandler // Http request handler that upgrades the connection to ws and implements WS functionality
	nextConnID  int                         // ID of the next incoming connection
}

func NewWs(upgrader *websocket.Upgrader) *WS {
	return &WS{
		nextConnID: 0,
		Upgrader:   upgrader,
	}
}

func (ws *WS) GetHandler() RequestHandler {
	return ws.Handler(ws)
}

func (ws *WS) AddNewConnection(conn *websocket.Conn) int {
	ws.Connections[ws.nextConnID] = conn
	ws.nextConnID++
	return ws.nextConnID
}
