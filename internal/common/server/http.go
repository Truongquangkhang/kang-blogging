package server

import (
	"net/http"
	"os"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/sirupsen/logrus"
)

func RunHTTPServer(
	createHandler func(router chi.Router) http.Handler,
) {
	httpPort := os.Getenv("HTTP_PORT")
	if httpPort == "" {
		httpPort = "10080"
	}
	RunHTTPServerOnAddr(":"+httpPort, createHandler)
}

func RunHTTPServerOnAddr(
	addr string,
	createHandler func(router chi.Router) http.Handler,
) {
	apiRouter := chi.NewRouter()
	setMiddlewares(apiRouter)

	rootRouter := chi.NewRouter()
	// we are mounting all APIs under /api path
	rootRouter.Mount("/", createHandler(apiRouter))

	server := &http.Server{
		Addr:              addr,
		Handler:           rootRouter,
		ReadHeaderTimeout: 20 * time.Second,
		ReadTimeout:       1 * time.Minute,
		WriteTimeout:      2 * time.Minute,
	}
	logrus.WithField("httpEndpoint", addr).Info("Starting HTTP server")
	logrus.Fatal(server.ListenAndServe())
}

func setMiddlewares(router *chi.Mux) {
	router.Use(middleware.Heartbeat("/health"))
	router.Use(middleware.RequestID)
	router.Use(middleware.RealIP)
	router.Use(logs.NewStructuredLogger(logrus.StandardLogger()))
	router.Use(middleware.Recoverer)

	router.Use(
		middleware.SetHeader("X-Content-Type-Options", "nosniff"),
		middleware.SetHeader("X-Frame-Options", "deny"),
	)
	router.Use(middleware.NoCache)
}
