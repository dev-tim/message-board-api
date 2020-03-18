package apiserver

import (
	"encoding/json"
	"net/http"
)

func (s *APIServer) handleHealth() http.HandlerFunc {

	type HealthResponse struct {
		Status string `json:"status"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		response := HealthResponse{
			Status: "OK",
		}

		json.NewEncoder(w).Encode(response)
	}
}

func (s *APIServer) handleGetPublicMessages() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		limit, err := extractIntParam(r, "limit", 10)
		if err != nil {
			http.Error(w, "Invalid limit value", 400)
			return
		}

		offset, err := extractIntParam(r, "offset", 10)
		if err != nil {
			http.Error(w, "Invalid offset value", 400)
			return
		}

		messages, err := s.store.Messages().FindLatest(limit, offset)
		if err != nil {
			http.Error(w, "Unable to fetch messages", 500)
			return
		}

		json.NewEncoder(w).Encode(messages)
	}
}

func (s *APIServer) handlePostPublicMessage() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "Not implemented", 412)
	}
}
