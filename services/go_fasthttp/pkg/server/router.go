package server

import (
	"fmt"
	"time"

	"github.com/fate-lovely/phi"
	"github.com/google/uuid"
	"github.com/nuclio/errors"
	"github.com/nuclio/logger"
	"github.com/valyala/fasthttp"
)

type Router struct {
	router           *phi.Mux
	logger           logger.Logger
	pyServiceAddress string
	startTime        time.Time
}

func NewRouter(parentLogger logger.Logger, pyServiceAddress string) *Router {
	return &Router{
		router:           phi.NewRouter(),
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

func (r *Router) InstallMiddleware(ctx *fasthttp.RequestCtx) {
	r.router.Use(r.requestID)
}

func (r *Router) Handler() fasthttp.RequestHandler {
	return r.router.ServeFastHTTP
}

func (r *Router) registerRoutes() error {
	r.router.Get("/", r.index)
	r.router.Get("/hello", r.hello)
	r.router.Get("/py-proxy", r.pyProxy)
	r.router.Get("/runtime", r.runtime)

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

	responseBuffer := ctx.Response.Body()
	statusCode, responseBody, err := fasthttp.Get(responseBuffer, url)
	if err != nil {
		r.logger.ErrorWith("Failed to proxy request to Python service", "url", url, "err", err.Error())
		ctx.SetStatusCode(fasthttp.StatusInternalServerError)
		return
	}

	ctx.SetStatusCode(statusCode)
	ctx.SetBody(responseBody)
}

func (r *Router) runtime(ctx *fasthttp.RequestCtx) {
	r.logger.Debug("Handling runtime request")
	runtime := r.calculateRuntime()
	ctx.SetBodyString("App runtime: " + runtime.String())
}

func (r *Router) calculateRuntime() time.Duration {
	return time.Since(r.startTime)
}

func (r *Router) requestID(next phi.HandlerFunc) phi.HandlerFunc {
	id := uuid.New().String()
	return func(ctx *fasthttp.RequestCtx) {
		next(ctx)
		ctx.WriteString(id) // nolint: errcheck
	}
}
