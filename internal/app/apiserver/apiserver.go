package apiserver

import (
	"encoding/json"
	"github.com/dev-tim/message-board-api/internal/app/common"
	"github.com/dev-tim/message-board-api/internal/app/store"
	"github.com/gorilla/mux"
	"net/http"
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
