package api

import (
	"log"
	"net/http"
)

func Serve(addr string) {
	initRouter()
	log.Println("Listening API on: ", addr, "...")
	log.Fatal(http.ListenAndServe(addr, nil))
}
