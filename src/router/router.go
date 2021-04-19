package router

import (
	"api/src/router/routes"

	"github.com/gorilla/mux"
)

//Generate will return one router with configurated routes
func Generate() *mux.Router {
	r := mux.NewRouter()
	return routes.Config(r)
}
