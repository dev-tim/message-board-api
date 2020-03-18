package apiserver

import (
	"github.com/dev-tim/message-board-api/internal/app/common"
	"github.com/dev-tim/message-board-api/internal/app/store"
	"github.com/gorilla/mux"
	"github.com/urfave/negroni"
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
	s.router.Use(contextMiddleware)
	s.router.Use(loggingMiddleware)

	publicPrefix := "/public"
	privatePrefix := "/private"

	publicRouter := mux.NewRouter().PathPrefix(publicPrefix).Subrouter().StrictSlash(true)
	privateRouter := mux.NewRouter().PathPrefix(privatePrefix).Subrouter().StrictSlash(true)

	s.router.HandleFunc("/health", s.handleHealth())
	privateRouter.HandleFunc("/v1/messages", s.handleGetPublicMessages())
	publicRouter.HandleFunc("/v1/messages", s.handleGetPublicMessages())

	n := negroni.New()
	recovery := negroni.NewRecovery()
	recovery.PanicHandlerFunc = func(panic *negroni.PanicInformation) {
		common.GetLogger().Error("Caught panic", panic.RequestDescription(), panic.StackAsString())
	}

	s.router.PathPrefix(privatePrefix).Handler(negroni.New(
		negroni.HandlerFunc(func(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
			if BasicAuth(w, r, "Provide user name and password") {
				next(w, r)
			}
		}),
		negroni.Wrap(privateRouter),
	))

	s.router.PathPrefix(publicPrefix).Handler(negroni.New(
		negroni.HandlerFunc(func(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
			// TODO do rate limiting here
			next(w, r)
		}),
		negroni.Wrap(publicRouter),
	))

	n.Use(recovery)
	n.UseHandler(s.router)
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
