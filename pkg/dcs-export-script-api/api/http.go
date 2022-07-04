package api

import (
	udpConnection "github.com/thedoctor028/dcsexportscriptapi/udp-connection"
	"log"
	"net/http"
)

func Serve(addr string, udpClient *udpConnection.UDPSender) {
	initRouter(udpClient)
	log.Println("Listening API on: ", addr, "...")
	log.Fatal(http.ListenAndServe(addr, nil))
}
