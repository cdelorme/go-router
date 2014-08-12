package routing

import "net/http"

type Controller interface {
	RegisterWithRouter(
		prefix string,
		addRoute func(
			uri string,
			callback func(
				writer http.ResponseWriter,
				request *http.Request),
			methods ...string))
}
