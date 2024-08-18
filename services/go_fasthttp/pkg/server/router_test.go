package server

import (
	"testing"
	"time"

	"github.com/nuclio/logger"
	nucliozap "github.com/nuclio/zap"
	"github.com/stretchr/testify/suite"
)

type RouterSuite struct {
	suite.Suite
	router *Router
	logger logger.Logger
}

func (suite *RouterSuite) SetupTest() {
	var err error
	suite.logger, err = nucliozap.NewNuclioZapTest("test")
	suite.Require().NoError(err)
	suite.router = NewRouter(suite.logger, "http://py-flask:8000")
}

func (suite *RouterSuite) TestCalculateRuntime() {
	time.Sleep(1 * time.Second)
	runtime := suite.router.calculateRuntime()
	suite.GreaterOrEqual(runtime, 1*time.Second)
}

func TestRouterSuite(t *testing.T) {
	suite.Run(t, new(RouterSuite))
}
