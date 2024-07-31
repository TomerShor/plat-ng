package server

import (
	"fmt"
	"github.com/buaazp/fasthttprouter"
	"github.com/nuclio/errors"
	"github.com/nuclio/logger"
	"github.com/valyala/fasthttp"
	"time"
)

type Router struct {
	router           *fasthttprouter.Router
	logger           logger.Logger
	pyServiceAddress string
	startTime        time.Time
}

func NewRouter(parentLogger logger.Logger, pyServiceAddress string) *Router {
	return &Router{
		router:           fasthttprouter.New(),
		logger:           parentLogger.GetChild("router"),
		pyServiceAddress: pyServiceAddress,
		startTime:        time.Now(),
	}
}

func (r *Router) Start() error {
	r.startTime = time.Now()
	r.logger.DebugWith("Starting router", "startTime", r.startTime)
	if err := r.registerRoutes(); err != nil {
		return errors.Wrap(err, "Failed to register routes")
	}
	return nil
}

func (r *Router) Stop() {
	r.logger.Debug("Stopping router")
}

func (r *Router) Handler() fasthttp.RequestHandler {
	return r.router.Handler
}

func (r *Router) registerRoutes() error {
	r.router.GET("/", r.index)
	r.router.GET("/hello", r.hello)
	r.router.GET("/py-proxy", r.pyProxy)
	r.router.GET("/runtime", r.runtime)

	return nil
}

func (r *Router) index(ctx *fasthttp.RequestCtx) {
	r.logger.Debug("Handling index request")
	ctx.SetBodyString("Hello World")
}

func (r *Router) hello(ctx *fasthttp.RequestCtx) {
	r.logger.Debug("Handling Hello request")
	ctx.SetBodyString("Hello from Go-Fasthttp service")
}

func (r *Router) pyProxy(ctx *fasthttp.RequestCtx) {
	subPath := string(ctx.QueryArgs().Peek("path"))
	r.logger.DebugWith("Handling Python proxy request", "subPath", subPath)
	url := fmt.Sprintf("%s/%s", r.pyServiceAddress, subPath)
	r.logger.DebugWith("Proxying request to Python service", "url", url)

	// Redirecting like this doesn't work...
	ctx.Redirect(url, fasthttp.StatusMovedPermanently)
}

func (r *Router) runtime(ctx *fasthttp.RequestCtx) {
	r.logger.Debug("Handling runtime request")
	runtime := r.calculateRuntime()
	ctx.SetBodyString("App runtime: " + runtime.String())
}

func (r *Router) calculateRuntime() time.Duration {
	return time.Since(r.startTime)
}
