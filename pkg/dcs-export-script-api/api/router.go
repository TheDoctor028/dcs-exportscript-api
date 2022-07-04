package api

import (
	udpConnection "github.com/thedoctor028/dcsexportscriptapi/udp-connection"
	"net/http"
)

type RequestHandler = func(w http.ResponseWriter, r *http.Request)

type Router struct {
	Routes []Route
}

type Route struct {
	Path    string
	Handler RequestHandler
}

func NewRouter() Router {
	return Router{}
}

func (r Router) InitRoutes() {
	for _, route := range r.Routes {
		route.InitRoute()
	}
}

func (r Route) InitRoute() {
	http.HandleFunc(r.Path, r.Handler)
}

func initRouter(udpClient *udpConnection.UDPClient) {
	http.HandleFunc("/test", setUpWSConnection(udpClient))
}
