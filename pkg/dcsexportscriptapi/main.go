package main

import (
	"github.com/thedoctor028/dcsexportscriptapi/udpConnection"
	"github.com/thedoctor028/dcsexportscriptapi/utils"
	"log"
	"os"
)

var logger = log.New(os.Stdout, "Main: ", 0)

var dataLogger = initDataLogger()

var buffer = make([]byte, 1024*2)

func main() {
	udpConnection.ServerUDP("127.0.0.1", 8080, &buffer, cbOnBufferListening)
}

func cbOnBufferListening() {
	res := utils.ExtractUIDAndValue(string(buffer), ":")
	dataLogger.Println(res)
}

func initDataLogger() log.Logger {
	f, _ := os.OpenFile("./data.logs.txt", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0600)
	return *log.New(f, " ////// ", 99)
}
