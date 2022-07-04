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
