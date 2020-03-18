package apiserver

import (
	"github.com/dev-tim/message-board-api/internal/app/common"
	"github.com/dev-tim/message-board-api/internal/app/store"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"github.com/urfave/negroni"
	"net/http"
)

type APIServer struct {
	router *mux.Router
	logger *logrus.Logger
	store  store.IStore
	config *Config
}

func New(config *Config, store store.IStore, logger *logrus.Logger) *APIServer {
	return &APIServer{
		logger: logger,
		router: mux.NewRouter(),
		store:  store,
		config: config,
	}
}

func (s *APIServer) Start() error {
	s.configureRouter()
	s.logger.Info("Started api server")

	return http.ListenAndServe(s.config.BindAddress, s.router)
}

func (s *APIServer) configureRouter() {
	s.router.Use(ContextMiddleware)
	s.router.Use(LoggingMiddleware)

	publicPrefix := "/public"
	privatePrefix := "/private"

	publicRouter := mux.NewRouter().PathPrefix(publicPrefix).Subrouter().StrictSlash(true)
	privateRouter := mux.NewRouter().PathPrefix(privatePrefix).Subrouter().StrictSlash(true)

	s.router.HandleFunc("/health", s.HandleHealth())
	privateRouter.HandleFunc("/v1/messages", s.HandleGetPrivateMessages()).Methods(http.MethodGet)
	privateRouter.HandleFunc("/v1/messages/{messageId}", s.HandleGetPrivateSingleMessage()).Methods(http.MethodGet)
	privateRouter.HandleFunc("/v1/messages", s.HandlePostPrivateMessage()).Methods(http.MethodPost)
	privateRouter.HandleFunc("/v1/messages/{messageId}", s.HandleUpdatePrivateMessage()).Methods(http.MethodPatch)

	privateRouter.HandleFunc("/v1/messages", s.HandlePostPrivateMessage()).Methods(http.MethodPost)

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
