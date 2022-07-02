package api

import "net/http"

func initRouter() {
	http.HandleFunc("/test", setUpWSConnection)
}
