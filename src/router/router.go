package router

import "github.com/gorilla/mux"

//Generate will return one router with configurated routes
func Generate() *mux.Router {
	return mux.NewRouter()
}
