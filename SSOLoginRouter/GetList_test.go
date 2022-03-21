package SSOLoginRouter

import (
	"github.com/stretchr/testify/assert"
	"log"
)

func (suite *SSOLoginTestSuite) TestGetList() {
	structSSOLoginRouter := StructSSOLoginRouter{}
	getListResponse := structSSOLoginRouter.HitGetListApi(suite)
	log.Println("Asserting the API Response...")
	assert.Equal(suite.T(), 2, getListResponse.Result[0].Id)
	assert.True(suite.T(), getListResponse.Result[0].Active)
	assert.NotNil(suite.T(), getListResponse.Result[0].Url)
}
