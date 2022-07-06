package api

import (
	udpConnection "github.com/thedoctor028/dcsexportscriptapi/udp-connection"
	"net/http"
)

// RequestHandler alias for http handler function
type RequestHandler = func(w http.ResponseWriter, r *http.Request)

// Router
// Wrapper for store the routes
type Router struct {
	Routes []Route
}

// Route
// A http route with a handler
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
