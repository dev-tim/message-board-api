package apiserver

import (
	"database/sql"
	"encoding/json"
	"github.com/gorilla/mux"
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

func (s *APIServer) handleGetPrivateMessages() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		limit, err := extractIntParam(r, "limit", 10)
		if err != nil {
			json.NewEncoder(w).Encode(NewError(w, r, 400, "Invalid limit value"))
			return
		}

		offset, err := extractIntParam(r, "offset", 10)
		if err != nil {
			json.NewEncoder(w).Encode(NewError(w, r, 400, "Invalid offset value"))
			return
		}

		messages, err := s.store.Messages().FindLatest(limit, offset)
		if err != nil {
			json.NewEncoder(w).Encode(NewError(w, r, 500, "Unable to fetch messages"))
			return
		}

		json.NewEncoder(w).Encode(messages)
	}
}

func (s *APIServer) handleGetPrivateSingleMessage() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		messages, err := s.store.Messages().FindById(vars["messageId"])

		if err == sql.ErrNoRows {
			json.NewEncoder(w).Encode(NewError(w, r, 404, "No result found"))
		} else if err != nil {
			json.NewEncoder(w).Encode(NewError(w, r, 500, "Unable to fetch messages"))
			return
		}

		json.NewEncoder(w).Encode(messages)
	}
}

func (s *APIServer) handlePostPrivateMessage() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var post CreateMessageBodyV1ClientRequest
		err := json.NewDecoder(r.Body).Decode(post)
		if err != nil {
			json.NewEncoder(w).Encode(NewError(w, r, 400, "Unable to parse body"))
		}

		message := post.ToMessage()
		if created, err := s.store.Messages().Create(message); err != nil {
			json.NewEncoder(w).Encode(NewError(w, r, 500, "Unable to store message"))
		} else {
			json.NewEncoder(w).Encode(created)
		}
	}
}

func (s *APIServer) handleUpdatePrivateMessage() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)

		var patch PatchMessageBodyV1ClientRequest
		err := json.NewDecoder(r.Body).Decode(patch)
		if err != nil {
			json.NewEncoder(w).Encode(NewError(w, r, 400, "Unable to parse body"))
		}

		if updated, err := s.store.Messages().Update(vars["messageId"], patch.Text); err != sql.ErrNoRows {
			json.NewEncoder(w).Encode(NewError(w, r, 404, "Document not found"))
		} else if err != nil {
			json.NewEncoder(w).Encode(NewError(w, r, 500, "Unable to patch message"))
		} else {
			json.NewEncoder(w).Encode(updated)
		}
	}
}
