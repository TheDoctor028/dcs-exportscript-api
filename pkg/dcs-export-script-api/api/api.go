package api

import (
	"log"
	"net/http"
	"os"
)

var apiLogger = log.New(os.Stdout, "API Server: ", 101)

// API
// Wrapper for the http/ws server
type API struct {
	Ip     string // Host ip of the Http server
	Port   int    // Host port of the Http server
	Router Router // Router instance to handel requests
}

// NewAPI
// Constructor for the API struct
func NewAPI(ip string, port int) *API {
	api := API{
		ip,
		port,
		NewRouter(),
	}

	return &api
}

// Serve
// Fires up the http server on the given ip/port (in the constructor)
func (a API) Serve() error {
	mergedAddr := mergeIpAndPort(a)

	err := http.ListenAndServe(mergedAddr, nil)
	if err == nil {
		apiLogger.Printf("API server listening on %s:%d...", mergedAddr)
		a.Router.InitRoutes()
	}

	return err
}

// mergeIpAndPort
// Returns the merged ip:port as a string
func mergeIpAndPort(a API) string {
	return a.Ip + ":" + string(rune(a.Port))
}
