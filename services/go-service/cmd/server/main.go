package main

import (
	"flag"
	"github.com/TomerShor/plat-ng/services/go-service/pkg/common"
	"github.com/TomerShor/plat-ng/services/go-service/pkg/server"
	"github.com/nuclio/errors"
	"os"
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
	server := server.NewServer(listenPort, logLevel)
	if err := server.Start(); err != nil {
		return errors.Wrap(err, "Failed to start server")
	}
	return nil
}
