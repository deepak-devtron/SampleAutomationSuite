package SSOLoginRouter

import (
	"SampleAutomationSuite/testdata/Utils"
	"encoding/json"
	"fmt"
	"github.com/go-resty/resty/v2"
	"github.com/stretchr/testify/suite"
	"net/http"
)

type GetListResponseDto struct {
	Code   int    `json:"code"`
	Status string `json:"status"`
	Result []struct {
		Id     int    `json:"id"`
		Name   string `json:"name"`
		Url    string `json:"url"`
		Active bool   `json:"active"`
		Label  string `json:"label,omitempty"`
	} `json:"result"`
}

type GetSSODetailsResponse struct {
	Code   int    `json:"code"`
	Status string `json:"status"`
	Result struct {
		Id     int    `json:"id"`
		Name   string `json:"name"`
		Url    string `json:"url"`
		Config struct {
			Id     string `json:"id"`
			Label  string `json:"label"`
			Type   string `json:"type"`
			Name   string `json:"name"`
			Config struct {
				Issuer        string   `json:"issuer"`
				ClientID      string   `json:"clientID"`
				ClientSecret  string   `json:"clientSecret"`
				RedirectURI   string   `json:"redirectURI"`
				HostedDomains []string `json:"hostedDomains"`
			} `json:"config"`
		} `json:"config"`
		Active bool `json:"active"`
	} `json:"result"`
}

type InterfaceSSOLoginRouter interface {
	UnmarshalGivenResponseBody(response []byte, apiName string) StructSSOLoginRouter
}

type StructSSOLoginRouter struct {
	getListResponseDto    GetListResponseDto
	getSSODetailsResponse GetSSODetailsResponse
}

func (structSSOLoginRouter StructSSOLoginRouter) UnmarshalGivenResponseBody(response []byte, apiName string) StructSSOLoginRouter {
	switch apiName {
	case "GetList":
		json.Unmarshal(response, &structSSOLoginRouter.getListResponseDto)
	case "GetSSODetails":
		json.Unmarshal(response, &structSSOLoginRouter.getSSODetailsResponse)
	case "GetSSOConfigByName":
		json.Unmarshal(response, &structSSOLoginRouter.getSSODetailsResponse)
	case "GetTeamById":
		json.Unmarshal(response, &structSSOLoginRouter.getListResponseDto)
	case "UpdateSSODetails":
		json.Unmarshal(response, &structSSOLoginRouter.getSSODetailsResponse)
	}
	return structSSOLoginRouter
}

func (ssoLoginStruct StructSSOLoginRouter) HitGetListApi(suite *SSOLoginTestSuite) GetListResponseDto {

	baseStructRef := Utils.BaseStruct{}

	resp, err := baseStructRef.MakeApiCall(suite.SSOLoginApiClient, "/orchestrator/sso/list", http.MethodGet, "", true, nil)
	baseStructRef.HandleError(err, "TestGetList")

	ssoRouter := ssoLoginStruct.UnmarshalGivenResponseBody(resp.Body(), "GetList")
	return ssoRouter.getListResponseDto
}

func (ssoLoginStruct StructSSOLoginRouter) HitGetSSODetailsApi(suite *SSOLoginTestSuite, ssoDetailsId string) GetSSODetailsResponse {
	baseStructRef := Utils.BaseStruct{}
	resp, err := baseStructRef.MakeApiCall(suite.SSOLoginApiClient, "/orchestrator/sso/"+ssoDetailsId, http.MethodGet, "", true, nil)
	baseStructRef.HandleError(err, "TestGetSsoLoginWithCorrectId")

	ssoRouter := ssoLoginStruct.UnmarshalGivenResponseBody(resp.Body(), "GetSSODetails")
	return ssoRouter.getSSODetailsResponse
}

/*
func HitGetLoginConfigByNameApi(queryParams map[string]string) GetSSODetailsResponse {
	resp, err := Base.MakeApiCall("/orchestrator/sso", http.MethodGet, "", true, queryParams)
	Base.HandleError(err, "TestGetSsoLoginConfigWithCorrectName")

	structSSOLoginRouter := StructSSOLoginRouter{}
	ssoRouter := structSSOLoginRouter.UnmarshalGivenResponseBody(resp.Body(), "GetSSOConfigByName")
	return ssoRouter.getSSODetailsResponse
}

func HitUpdateSSODetailsApi(byteValue []byte) GetSSODetailsResponse {
	resp, err := Base.MakeApiCall("/orchestrator/sso/update", http.MethodPut, string(byteValue), true, nil)
	Base.HandleError(err, "TestUpdateSsoLoginWithCorrectArgs")

	structSSOLoginRouter := StructSSOLoginRouter{}
	ssoRouter := structSSOLoginRouter.UnmarshalGivenResponseBody(resp.Body(), "UpdateSSODetails")
	return ssoRouter.getSSODetailsResponse
}
*/

type SSOLoginTestSuite struct {
	suite.Suite
	SSOLoginApiClient *resty.Client
}

func (suite *SSOLoginTestSuite) SetupTest() {
	fmt.Println("I am calling setup test method")
	suite.SSOLoginApiClient = resty.New()
}
