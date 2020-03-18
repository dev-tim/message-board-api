package apiserver

import (
	"encoding/json"
	"github.com/dev-tim/message-board-api/internal/app/common"
	"github.com/dev-tim/message-board-api/internal/app/store"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

type APIServer struct {
	name   string
	config *Config
	router *mux.Router
	store  *store.Store
}

func New(config *Config) *APIServer {
	return &APIServer{
		name:   "App",
		config: config,
		router: mux.NewRouter(),
	}
}

func (s *APIServer) Start() error {
	if _, err := common.NewLoggerFactory(s.config.Common); err != nil {
		return err
	}

	if err := s.configureStore(); err != nil {
		return err
	}

	s.configureRouter()
	common.GetLogger().Info("Started api server")

	return http.ListenAndServe(s.config.BindAddress, s.router)
}

func (s *APIServer) configureRouter() {
	s.router.HandleFunc("/health", s.handleHealth())
	s.router.HandleFunc("/private/api/v1/messages", s.handleGetPublicMessages())
}

func (s *APIServer) configureStore() error {
	st := store.New(s.config.Store)
	if err := st.Open(); err != nil {
		return err
	}

	s.store = st

	if err := st.Migrate(); err != nil {
		return err
	}
	return nil
}

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

func extractIntParam(r *http.Request, key string, defaultVal int) (*int, error) {
	query := r.URL.Query()
	val := query.Get(key)

	if len(val) == 0 {
		return &defaultVal, nil
	}

	if atoi, err := strconv.Atoi(val); err != nil {
		return nil, err
	} else {
		return &atoi, nil
	}
}
