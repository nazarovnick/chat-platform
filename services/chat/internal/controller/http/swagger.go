package http

import (
	"github.com/gorilla/mux"
	"net/http"
)

func RegisterSwaggerRoutes(router *mux.Router) {
	router.HandleFunc("/swagger", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/swagger/", http.StatusMovedPermanently)
	})

	swagger := router.PathPrefix("/swagger/").Subrouter()

	// OpenAPI .json-file
	swagger.HandleFunc("/{filename:.+\\.json}", swaggerJSONHandler).Methods(http.MethodGet)

	// Swagger UI
	swagger.PathPrefix("/").Handler(
		http.StripPrefix("/swagger/", http.FileServer(http.Dir("./web/static/swagger-ui-5.21.0/dist/"))),
	)
}

func swaggerJSONHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	filename := vars["filename"]
	http.ServeFile(w, r, "./api/openapi/v1/chat/"+filename)
}
