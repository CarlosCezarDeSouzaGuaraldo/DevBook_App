package routes

import (
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

	for _, route := range routes {
		r.HandleFunc(route.URI, route.Func).Methods(route.Method)
	}

	return r
}
