package routing

import "net/http"

type Controller interface {
	RegisterWithRouter(
		addRoute func(
			uri string,
			callback func(
				writer http.ResponseWriter,
				request *http.Request),
			methods ...string))
}
