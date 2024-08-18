package server

import (
	"fmt"
	"time"

	"github.com/TomerShor/plat-ng/services/go_fasthttp/pkg/common"

	"github.com/nuclio/errors"
	"github.com/nuclio/logger"
	"github.com/valyala/fasthttp"
)

// Server is a simple http server that uses fasthttp, listens on port 8010 has some basic routes and can
// proxy requests to a python service
type Server struct {
	listenPort int
	logger     logger.Logger
	startTime  time.Time
}

// NewServer creates a new server
func NewServer(listenPort int, parentLogger logger.Logger) *Server {
	newLogger := parentLogger.GetChild("server")
	return &Server{
		listenPort: listenPort,
		logger:     newLogger,
	}
}

func (s *Server) Start() error {
	s.startTime = time.Now()
	s.logger.DebugWith("Starting server",
		"listenPort", s.listenPort,
		"startTime", time.Now())

	pyServiceAddress := common.GetEnvOrDefault("PY_SERVICE_URL", "http://py-flask:8000")

	router := NewRouter(s.logger, pyServiceAddress)
	if err := router.Start(); err != nil {
		return errors.Wrap(err, "Failed to start router")
	}

	if err := fasthttp.ListenAndServe(fmt.Sprintf(":%d", s.listenPort), router.Handler()); err != nil {
		panic("Error in ListenAndServe: " + err.Error())
	}

	return nil
}

func (s *Server) Stop() {
	s.logger.Debug("Stopping server")
}
