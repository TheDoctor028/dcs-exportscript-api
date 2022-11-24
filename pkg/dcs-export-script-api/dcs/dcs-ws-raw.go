package DCS

import (
	"github.com/gorilla/websocket"
	"github.com/thedoctor028/dcsexportscriptapi/api"
	"log"
	"net/http"
)

func (c *Service) initRawRouteWS() *api.WS {
	ws := api.NewWs(&websocket.Upgrader{CheckOrigin: func(r *http.Request) bool {
		return true
	}})

	ws.Handler = func(ws *api.WS) api.RequestHandler {
		return func(w http.ResponseWriter, r *http.Request) {
			conn, err := ws.Upgrader.Upgrade(w, r, nil)
			if err != nil {
				dcsClientLogger.Print("Upgrade connection failed! %s", err)
				return
			}
			defer conn.Close()

			ws.AddNewConnection(conn)

			listenForCommands(conn, c)
		}
	}

	return ws
}

// Waits for commands on the ws to send it to DCS ExportScript UDP server
// Example.:C12,3022,0 for more details see ExportScript docs
func listenForCommands(conn *websocket.Conn, c *Service) {
	for {
		_, message, err := conn.ReadMessage()
		if err != nil {
			log.Println("read:", err)
			break
		}
		log.Printf("recv: %s", message)

		if string(message)[0] == 'C' {
			err := c.udpClient.SendData(string(message))
			if err != nil {
				println(err)
			}
		}
	}
}
