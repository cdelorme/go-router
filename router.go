package routing

import (
	"encoding/json"
	"errors"
	"net/http"
)

type Router struct {
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

	// add 404 handler to replace this logic
	// router.Log.Debug("No callback registered for: %s", request.URL.String())
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
	methods ...string) (Route, error) {
	if len(methods) == 0 {
		return Route{}, errors.New("No request method supplied")
	}
	return Route{URI: uri, Callback: callback, RequestMethods: methods}, nil
}

func (router *Router) CreateAndRegisterRoute(
	uri string,
	callback func(writer http.ResponseWriter, request *http.Request),
	methods ...string) error {
	route, err := router.CreateRoute(uri, callback, methods...)
	if err != nil {
		return err
	}
	router.RegisterRoute(route)
	return nil
}

func (router *Router) RegisterController(controller Controller) error {
	if controller == nil {
		return errors.New("No controller supplied to registering with router")
	}
	controller.RegisterWithRouter(router.CreateAndRegisterRoute)
	return nil
}
