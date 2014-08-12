package routing

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/cdelorme/go-log"
)

type Router struct {
	Log    *log.Logger
	Routes []Route
}

var ignoreList = []string{"/favicon.ico"}

func (router *Router) ServeHTTP(writer http.ResponseWriter, request *http.Request) {

	for _, ignore := range ignoreList {
		if ignore == request.URL.String() {
			return
		}
	}

	for _, at := range router.Routes {
		if at.URI == request.URL.String() {
			for _, method := range at.RequestMethods {
				if method == request.Method {
					at.Callback(writer, request)
					return
				}
			}
		}
	}

	router.Log.Debug("No callback registered for: %s", request.URL.String())
	message := struct{ Error string }{Error: "No resources available at " + request.URL.String() + " using request method " + request.Method}
	jsonMessage, _ := json.Marshal(message)
	writer.Header().Add("Content-Type", "application/json")
	writer.Write(jsonMessage)
}

func (router *Router) RegisterRoute(routes ...Route) {
	if len(routes) == 0 {
		router.Log.Error("No routes provided")
		return
	}
	router.Routes = append(router.Routes, routes...)
}

func (router *Router) CreateRoute(
	uri string,
	callback func(writer http.ResponseWriter, request *http.Request),
	methods ...string) (Route, error) {
	if len(methods) == 0 {
		return Route{}, errors.New("No request method supplied")
	}
	return Route{URI: uri, Callback: callback, RequestMethods: methods}, nil
}

func (router *Router) CreateAndRegisterRoute(
	uri string,
	callback func(writer http.ResponseWriter, request *http.Request),
	methods ...string) {
	route, err := router.CreateRoute(uri, callback, methods...)
	if err != nil {
		router.Log.Error("%s", err)
		return
	}
	router.RegisterRoute(route)
}

func (router *Router) RegisterController(controller Controller, prefix string) {
	if controller == nil {
		router.Log.Error("No controller supplied to registering with router")
		return
	}
	controller.RegisterWithRouter(prefix, router.CreateAndRegisterRoute)
}

func (router *Router) Start(server *http.Server, serverAddress string) {
	if server == nil {
		if serverAddress == "" {
			router.Log.Error("No address and port provided to serve content on")
			return
		}
		server = &http.Server{}
	}
	server.Handler = router
	if serverAddress != "" {
		server.Addr = serverAddress
	}
	server.MaxHeaderBytes = 1 << 20
	err := server.ListenAndServe()
	if err != nil {
		router.Log.Error("Encountered an error setting up the server: %s", err.Error())
	}
}
