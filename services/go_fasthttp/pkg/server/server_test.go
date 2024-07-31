package server

import (
	"github.com/nuclio/logger"
	nucliozap "github.com/nuclio/zap"
	"github.com/stretchr/testify/suite"
	"testing"
	"time"
)

type ServerSuite struct {
	suite.Suite
	server *Server
	logger logger.Logger
}

func (suite *ServerSuite) SetupTest() {
	var err error
	suite.logger, err = nucliozap.NewNuclioZapTest("test")
	suite.Require().NoError(err)
	suite.server = NewServer(8080, suite.logger)
	go func() {
		suite.server.Start()
	}()
}

func (suite *ServerSuite) TestCalculateRuntime() {
	time.Sleep(1 * time.Second)
	runtime := suite.server.calculateRuntime()
	suite.GreaterOrEqual(runtime, 1*time.Second)
}

func TestServerSuite(t *testing.T) {
	suite.Run(t, new(ServerSuite))
}
