package SSOLoginRouter

import (
	"github.com/stretchr/testify/assert"
	"log"
)

func (suite *SSOLoginTestSuite) TestGetSsoLoginWithCorrectId() {
	structSSOLoginRouter := StructSSOLoginRouter{}
	log.Println("Hitting the Get SSO Details API")
	actualSSODetailsResponse := structSSOLoginRouter.HitGetSSODetailsApi(suite, "1")

	log.Println("Asserting the API Response...")
	assert.Equal(suite.T(), "https://deepak.devtron.info:32443/orchestrator", actualSSODetailsResponse.Result.Url)
	assert.Equal(suite.T(), "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", actualSSODetailsResponse.Result.Config.Config.ClientSecret)
	assert.Equal(suite.T(), "https://deepak.devtron.info:32443/orchestrator/api/dex/callback", actualSSODetailsResponse.Result.Config.Config.RedirectURI)
}

func (suite *SSOLoginTestSuite) TestGetSsoLoginWithInCorrectId() {
	structSSOLoginRouter := StructSSOLoginRouter{}
	log.Println("Hitting the Get SSO Details API")
	actualSSODetailsResponse := structSSOLoginRouter.HitGetSSODetailsApi(suite, "500")

	log.Println("Asserting the API Response...")
	assert.Empty(suite.T(), actualSSODetailsResponse.Result.Url)
	assert.Empty(suite.T(), actualSSODetailsResponse.Result.Config.Config.ClientSecret)
	assert.Empty(suite.T(), actualSSODetailsResponse.Result.Config.Config.RedirectURI)
}
