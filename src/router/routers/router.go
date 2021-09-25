package routers

import (
	"github.com/gorilla/mux"
	"net/http"
)

// Router represent all router of API
type Router struct {
	URI string
	Method string
	Function func(w http.ResponseWriter, r *http.Request)
	Authentication bool
}

// Configuration append all routers
func Configuration(r *mux.Router) *mux.Router {
	routers := usersRouters

	for _, router := range routers {
		r.HandleFunc(router.URI, router.Function).Methods(router.Method)
	}
	return r
}