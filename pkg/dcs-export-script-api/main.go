package main

import (
	DCS "github.com/thedoctor028/dcsexportscriptapi/dcs"
	"github.com/thedoctor028/dcsexportscriptapi/utils"
	"log"
	"net"
	"os"
	"strconv"
	"time"
)

var logger = log.New(os.Stdout, "Main: ", 0)

var dataLogger, loggerFile = initDataLogger()

func main() {
	service := DCS.NewService()

	err := service.CreateAndStartConnections()

	if err != nil {
		logger.Printf("Failed to create DCS Service! %s", err)
	}

	defer loggerFile.Close()
	wait()
}

func cbOnDataReceived(buffer *[]byte, remoteAddr *net.UDPAddr) { // TODO move to Dcs Service
	res := utils.ExtractUIDAndValue(string(*buffer), ":")
	dataLogger.Println(res.ToString())
	dataScreenData := res.GetDataByUid(50)
	if dataScreenData != nil {
		//api.SendEventToAllConnections(*dataScreenData)
	}
}

func initDataLogger() (*log.Logger, *os.File) {
	f, _ := os.OpenFile("./logs/data.logs-"+strconv.FormatInt(time.Now().Unix(), 10)+".txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0600)
	return log.New(f, " //////\n ", 101), f
} // TODO move to DCS service

func wait() {
	for {
	}
}
