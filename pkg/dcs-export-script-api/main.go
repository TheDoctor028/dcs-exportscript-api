package main

import (
	DCS "github.com/thedoctor028/dcsexportscriptapi/dcs"
	"log"
	"os"
)

var logger = log.New(os.Stdout, "Main: ", 0)

func main() {
	service := DCS.NewService()
	defer service.Destroy()

	err := service.CreateAndStartConnections()

	if err != nil {
		logger.Printf("Failed to create DCS Service! %s", err)
	}

	wait()
}

func wait() {
	for {
	}
}
