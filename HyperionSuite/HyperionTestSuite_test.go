package HyperionSuite

import (
	"SampleAutomationSuite/SSOLoginRouter"
	"github.com/stretchr/testify/suite"
	"testing"
)

func TestSSOLoginRouterSuite(t *testing.T) {
	suite.Run(t, new(SSOLoginRouter.SSOLoginTestSuite))
}
