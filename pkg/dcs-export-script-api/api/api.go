package api

import (
	"log"
	"net/http"
	"os"
)

var apiLogger = log.New(os.Stdout, "API Server: ", 101)

type API struct {
	Ip     string // Host ip of the Http server
	Port   int    // Host port of the Http server
	Router Router // Router instance to handel requests
}

func NewAPI(ip string, port int) *API {
	api := API{
		ip,
		port,
		NewRouter(),
	}

	return &api
}

func (a API) Serve() error {
	mergedAddr := mergeIpAndPort(a)

	err := http.ListenAndServe(mergedAddr, nil)
	if err == nil {
		apiLogger.Printf("API server listening on %s:%d...", mergedAddr)
		a.Router.InitRoutes()
	}

	return err
}

func mergeIpAndPort(a API) string {
	return a.Ip + ":" + string(rune(a.Port))
}
