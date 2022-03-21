package TeamRouter

import (
	"github.com/stretchr/testify/suite"
	"testing"
)

func TestTeamRouterSuite1(t *testing.T) {
	suite.Run(t, new(TeamTestSuite))
}
