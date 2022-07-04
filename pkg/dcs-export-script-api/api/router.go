package api

import (
	udpConnection "github.com/thedoctor028/dcsexportscriptapi/udp-connection"
	"net/http"
)

func initRouter(udpClient *udpConnection.UDPSender) {
	http.HandleFunc("/test", setUpWSConnection(udpClient))
}
