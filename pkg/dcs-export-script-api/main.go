package main

import (
	"github.com/thedoctor028/dcsexportscriptapi/api"
	"github.com/thedoctor028/dcsexportscriptapi/udp-connection"
	"github.com/thedoctor028/dcsexportscriptapi/utils"
	"log"
	"net"
	"os"
	"strconv"
	"time"
)

var logger = log.New(os.Stdout, "Main: ", 0)

var dataLogger = initDataLogger()

func main() {
	server, _ := udpConnection.NewUDPServer(net.UDPAddr{IP: net.ParseIP("127.0.0.1"), Port: 1625})
	udpClient, _ := udpConnection.NewUDPClient(1627)
	udpClient.Target = net.UDPAddr{IP: net.ParseIP("127.0.0.1"), Port: 1626}

	server.CB = cbOnDataReceived
	go server.Serve()
	go api.Serve("127.0.0.1:8000", udpClient)
	wait()
}

func cbOnDataReceived(buffer *[]byte, remoteAddr *net.UDPAddr) {
	res := utils.ExtractUIDAndValue(string(*buffer), ":")
	dataLogger.Println(res.ToString())
	dataScreenData := res.GetDataByUid(50)
	if dataScreenData != nil {
		api.SendEventToAllConnections(*dataScreenData)
	}
}

func initDataLogger() *log.Logger {
	f, _ := os.OpenFile("./logs/data.logs-"+strconv.FormatInt(time.Now().Unix(), 10)+".txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0600)
	defer f.Close()
	return log.New(f, " ////// ", 101)
}

func wait() {
	for {
	}
}
