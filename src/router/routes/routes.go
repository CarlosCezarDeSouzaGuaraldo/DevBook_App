package routes

import (
	"api/src/middlewares"
	"net/http"

	"github.com/gorilla/mux"
)

// Route represents all API's routes
type Route struct {
	URI            string
	Method         string
	Func           func(http.ResponseWriter, *http.Request)
	AuthIsRequired bool
}

// Configure put all routes in router
func Configure(r *mux.Router) *mux.Router {
	routes := routeUsers
	routes = append(routes, loginRoute)

	for _, route := range routes {

		if route.AuthIsRequired {
			r.HandleFunc(route.URI,
				middlewares.Logger(middlewares.Auth(route.Func)),
			).Methods(route.Method)
		} else {
			r.HandleFunc(route.URI,
				middlewares.Logger(route.Func),
			).Methods(route.Method)
		}
	}

	return r
}
