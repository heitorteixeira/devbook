package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

//Route represents all API routes
type Route struct {
	URI         string
	Method      string
	Function    func(http.ResponseWriter, *http.Request)
	RequestAuth bool
}

//Config set all routes in the router
func Config(r *mux.Router) *mux.Router {
	routes := userRoutes
	routes = append(routes, routeLogin)

	for _, route := range routes {
		r.HandleFunc(route.URI, route.Function).Methods(route.Method)
	}

	return r
}
