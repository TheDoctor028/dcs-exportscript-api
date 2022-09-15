package api

import (
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"os"
)

var wsLogger = log.New(os.Stdout, "WS Service: ", 101)

type WS struct {
	Connections map[int]*websocket.Conn     // The currently active connection on the websocket
	Upgrader    *websocket.Upgrader         // Upgrader instance to upgrade HTTP -> WS
	Handler     func(ws *WS) RequestHandler // Http request handler that upgrades the connection to ws and implements WS functionality
	nextConnID  int                         // ID of the next incoming connection
}

func NewWs(upgrader *websocket.Upgrader) *WS {
	return &WS{
		nextConnID:  0,
		Upgrader:    upgrader,
		Connections: map[int]*websocket.Conn{},
		Handler: func(ws *WS) RequestHandler {
			return func(w http.ResponseWriter, r *http.Request) {

			}
		},
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

func (ws *WS) SendToAllConnections(data string) {
	for _, conn := range ws.Connections {
		err := conn.WriteMessage(websocket.TextMessage, []byte(data))

		if err != nil {
			wsLogger.Printf("Can't send data to client: %s\n", err)
		}
	}
}
