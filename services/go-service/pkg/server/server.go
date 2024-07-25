package server

import (
	"fmt"
	"github.com/TomerShor/plat-ng/services/go-service/pkg/common"
	"github.com/nuclio/logger"
	nucliozap "github.com/nuclio/zap"
	"github.com/valyala/fasthttp"
	"os"
)

// Server is a simple http server that uses fasthttp, listens on port 8010 and returns "hello world" on the root path
type Server struct {
	listenPort int
	logger     logger.Logger
}

// NewServer creates a new server
func NewServer(listenPort int, logSeverity string) *Server {
	newLogger, err := nucliozap.NewNuclioZap("server",
		"console",
		nil,
		os.Stdout,
		os.Stderr,
		common.ResolveLogLevel(logSeverity))
	if err != nil {
		panic(err)

	}
	return &Server{
		listenPort: listenPort,
		logger:     newLogger,
	}
}

func (s *Server) Start() error {
	s.logger.Debug("Starting server")
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
		case "/goodbye":
			s.handleGoodbye(ctx)
		default:
			ctx.Error(fmt.Sprintf("Path not found: %s", path), fasthttp.StatusNotFound)
		}
	}
}

func (s *Server) handleHello(ctx *fasthttp.RequestCtx) {
	s.logger.Debug("Handling hello request")
	ctx.SetBodyString("Hello World")
}

func (s *Server) handleGoodbye(ctx *fasthttp.RequestCtx) {
	s.logger.Debug("Handling goodbye request")
	ctx.SetBodyString("Goodbye World")
}
