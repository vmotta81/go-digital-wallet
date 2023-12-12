package route

import (
	"net/http"

	"github.com/gorilla/mux"
)

type routeModel struct {
	uri          string
	method       string
	function     func(w http.ResponseWriter, r *http.Request)
	authRequired bool
}

func CreateRouter() *mux.Router {
	router := mux.NewRouter()
	configureRoutes(router)
	return router
}

func configureRoutes(router *mux.Router) {
	for _, route := range accoutRoutes {
		router.HandleFunc(route.uri, route.function).Methods(route.method)
	}
}
