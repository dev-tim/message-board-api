package apiserver

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"net/http"
)

type APIServer struct {
	name   string
	config *Config
	logger *logrus.Logger
	router *mux.Router
}

func New(config *Config) *APIServer {
	return &APIServer{
		name:   "App",
		config: config,
		logger: logrus.New(),
		router: mux.NewRouter(),
	}
}

func (s *APIServer) Start() error {
	if err := s.configureLogger(s.config); err != nil {
		return err
	}

	s.configureRouter()
	s.logger.Info("Started api server")

	return http.ListenAndServe(s.config.BindAddress, s.router)
}

func (s *APIServer) configureLogger(config *Config) error {
	level, err := logrus.ParseLevel(config.LogLevel)
	if err != nil {
		return err
	}

	s.logger.SetLevel(level)
	return nil
}

func (s *APIServer) configureRouter() {
	s.router.HandleFunc("/health", s.handleHealth())
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
