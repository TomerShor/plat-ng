package main

import (
	"flag"
	"os"

	"github.com/TomerShor/plat-ng/services/go_fasthttp/pkg/common"
	"github.com/TomerShor/plat-ng/services/go_fasthttp/pkg/server"

	"github.com/nuclio/errors"
	nucliozap "github.com/nuclio/zap"
)

func main() {
	listenPort := flag.Int("listen-port", common.GetEnvOrDefault("LISTEN_PORT", 8010), "Port to listen on")
	logLevel := flag.String("log-level", common.GetEnvOrDefault("LOG_LEVEL", "debug"), "Log level")

	flag.Parse()

	if err := run(*listenPort, *logLevel); err != nil {
		errors.PrintErrorStack(os.Stderr, err, 5)

		os.Exit(1)
	}
	os.Exit(0)
}

func run(listenPort int, logLevel string) error {
	rootLogger, err := nucliozap.NewNuclioZap("app",
		"console",
		nil,
		os.Stdout,
		os.Stderr,
		common.ResolveLogLevel(logLevel))
	if err != nil {
		return errors.Wrap(err, "Failed to create logger")
	}
	server := server.NewServer(listenPort, rootLogger)
	if err := server.Start(); err != nil {
		return errors.Wrap(err, "Failed to start server")
	}
	return nil
}
