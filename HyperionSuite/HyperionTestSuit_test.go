package HyperionSuite

import (
	"SampleAutomationSuite/nishant"
	"github.com/stretchr/testify/suite"
	"testing"
)

func TestSSOLoginRouterSuite1(t *testing.T) {
	suite.Run(t, new(nishant.ExampleTestSuite))
}
