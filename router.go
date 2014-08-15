package routing

import (
	"encoding/json"
	"errors"
	"net/http"
)

type Router struct {
	Routes         []Route
	IgnoreList     []string
	HandleNotFound func(writer http.ResponseWriter, request *http.Request)
}

func (router *Router) ServeHTTP(writer http.ResponseWriter, request *http.Request) {

	if router.Ignored() {
		return
	}

	for _, at := range router.Routes {
		if at.URI == request.URL.String() || at.URI+"/" == request.URL.String() {
			for _, method := range at.RequestMethods {
				if method == request.Method {
					at.Callback(writer, request)
					return
				}
			}
		}
	}

	router.NotFound(writer, request)
}

func (router *Router) Ignore(uri string) {
	router.IgnoreList = append(router.IgnoreList, uri)
}

func (router *Router) Ignored(uri string) bool {
	for _, ignore := range router.IgnoreList {
		if ignore == uri {
			return true
		}
	}
	return false
}

func (router *Router) NotFound(writer http.ResponseWriter, request *http.Request) {
	if router.HandleNotFound != nil {
		router.HandleNotFound(writer, request)
		return
	}
	message := struct{ Error string }{Error: "No resources available at " + request.URL.String() + " using request method " + request.Method}
	jsonMessage, _ := json.Marshal(message)
	writer.Header().Add("Content-Type", "application/json")
	writer.Write(jsonMessage)
}

func (router *Router) RegisterRoute(routes ...Route) error {
	if len(routes) == 0 {
		return errors.New("No routes provided")
	}
	router.Routes = append(router.Routes, routes...)
	return nil
}

func (router *Router) CreateRoute(
	uri string,
	callback func(writer http.ResponseWriter, request *http.Request),
	methods ...string) Route {
	if len(methods) == 0 {
		methods = []string{"GET"}
	}
	return Route{URI: uri, Callback: callback, RequestMethods: methods}
}

func (router *Router) CreateAndRegisterRoute(
	uri string,
	callback func(writer http.ResponseWriter, request *http.Request),
	methods ...string) {
	route := router.CreateRoute(uri, callback, methods...)
	router.RegisterRoute(route)
}

func (router *Router) RegisterController(controller Controller) {
	controller.RegisterWithRouter(router.CreateAndRegisterRoute)
}
