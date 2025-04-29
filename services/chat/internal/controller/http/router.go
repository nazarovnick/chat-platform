package http

import (
	"github.com/gorilla/mux"
	"net/http"
)

func NewRouter(gwMux http.Handler) *mux.Router {
	router := mux.NewRouter()
	RegisterHealthRoutes(router)
	router.PathPrefix("/").Handler(gwMux)

	return router
}
