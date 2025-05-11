package http

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func RegisterHealthRoutes(router *mux.Router) {
	health := router.PathPrefix("/healthz").Subrouter()
	health.HandleFunc("/liveness", livenessHandler).Methods(http.MethodGet)
	health.HandleFunc("/readiness", readinessHandler).Methods(http.MethodGet)
}

type healthResponse struct {
	Status string `json:"status"`
}

func livenessHandler(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	resp := healthResponse{Status: "alive"}
	data, err := json.Marshal(resp)
	if err != nil {
		http.Error(w, "failed to marshal response", http.StatusInternalServerError)
		return
	}
	if _, err := w.Write(data); err != nil {
		log.Printf("failed to write response: %v", err)
	}
}

func readinessHandler(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	resp := healthResponse{Status: "ready"}
	data, err := json.Marshal(resp)
	if err != nil {
		http.Error(w, "failed to marshal response", http.StatusInternalServerError)
		return
	}
	if _, err := w.Write(data); err != nil {
		log.Printf("failed to write response: %v", err)
	}
}
