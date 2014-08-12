package routing

import "net/http"

type Route struct {
	URI            string
	Callback       func(writer http.ResponseWriter, request *http.Request)
	RequestMethods []string
}
