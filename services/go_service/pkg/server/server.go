package server

import (
	"fmt"
	"github.com/nuclio/logger"
	"github.com/valyala/fasthttp"
	"time"
)

// Server is a simple http server that uses fasthttp, listens on port 8010 and returns "hello world" on the root path
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
		"startTime", s.startTime.String())
	if err := fasthttp.ListenAndServe(fmt.Sprintf(":%d", s.listenPort), s.requestHandler()); err != nil {
		panic("Error in ListenAndServe: " + err.Error())
	}

	return nil
}

func (s *Server) Stop() {
	s.logger.Debug("Stopping server")
}

func (s *Server) requestHandler() func(ctx *fasthttp.RequestCtx) {
	return func(ctx *fasthttp.RequestCtx) {
		path := string(ctx.Path())
		s.logger.DebugWith("Handling request", "path", path)

		switch path {
		case "/", "/hello":
			s.handleHello(ctx)
		case "/runtime":
			s.handleRuntime(ctx)
		default:
			ctx.Error(fmt.Sprintf("Path not found: %s", path), fasthttp.StatusNotFound)
		}
	}
}

func (s *Server) handleHello(ctx *fasthttp.RequestCtx) {
	s.logger.Debug("Handling hello request")
	ctx.SetBodyString("Hello World")
}

func (s *Server) handleRuntime(ctx *fasthttp.RequestCtx) {
	s.logger.Debug("Handling runtime request")
	runtime := s.calculateRuntime()
	ctx.SetBodyString("App runtime: " + runtime.String())
}

func (s *Server) calculateRuntime() time.Duration {
	return time.Since(s.startTime)
}
