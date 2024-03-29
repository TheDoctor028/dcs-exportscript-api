package api

import (
	"errors"
	"log"
	"net/http"
	"os"
	"strconv"
)

var apiLogger = log.New(os.Stdout, "API Server: ", 101)

// WSName
// Type alias for WebSockets names (used for 'enums')
type WSName string

// API
// Wrapper for the http/ws server
type API struct {
	Ip         string         // Host ip of the Http server
	Port       int            // Host port of the Http server
	Router     *Router        // Router instance to handel requests
	Websockets map[WSName]*WS // Websockets of the api
}

// NewAPI
// Constructor for the API struct
func NewAPI(ip string, port int) *API {
	api := API{
		ip,
		port,
		NewRouter(),
		map[WSName]*WS{},
	}

	return &api
}

// Serve
// Fires up the http server on the given ip/port (in the constructor)
func (a *API) Serve() {
	mergedAddr := mergeIpAndPort(a)
	a.Router.InitRoutes()

	apiLogger.Printf("API server listening on %s...", mergedAddr)
	apiLogger.Fatal("API Failed: %s", http.ListenAndServe(mergedAddr, nil))
}

// AddWS
// Adds a new websocket connection
func (a *API) AddWS(name WSName, ws *WS) error {
	if a.Websockets[name] != nil {
		return errors.New("websocket already exists")
	}

	a.Websockets[name] = ws

	return nil
}

// mergeIpAndPort
// Returns the merged ip:port as a string
func mergeIpAndPort(a *API) string {
	return a.Ip + ":" + strconv.Itoa(a.Port)
}
