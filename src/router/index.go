package router

import (
	"api_resources/src/router/routers"
	"github.com/gorilla/mux"
)

// Generate return an ruter with configurations
func Generate() *mux.Router {
	r := mux.NewRouter()
	return routers.Configuration(r)
}