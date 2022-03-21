package TeamRouter

import (
	"SampleAutomationSuite/testdata/Utils"
	"encoding/json"
	"github.com/go-resty/resty/v2"
	"github.com/stretchr/testify/suite"
	"net/http"
)

type SaveTeamResponseDto struct {
	Code   int    `json:"code"`
	Status string `json:"status"`
	Result struct {
		Id     int    `json:"id"`
		Name   string `json:"name"`
		Active bool   `json:"active"`
	} `json:"result"`
}

type SaveTeamRequestDto struct {
	Name   string `json:"name"`
	Active bool   `json:"active"`
}

type UpdateTeamRequestDto struct {
	Id     int    `json:"id"`
	Name   string `json:"name"`
	Active bool   `json:"active"`
}

type DeleteTeamResponseDto struct {
	Code   int    `json:"code"`
	Status string `json:"status"`
	Result string `json:"result"`
}

type FetchAllTeamResponseDto struct {
	Code   int    `json:"code"`
	Status string `json:"status"`
	Result []struct {
		Id     int    `json:"id"`
		Name   string `json:"name"`
		Active bool   `json:"active"`
	} `json:"result"`
}

type TeamTestSuite struct {
	suite.Suite
	TeamApiClient *resty.Client
}

func (suite *TeamTestSuite) SetupTest() {
	suite.TeamApiClient = resty.New()
}

func (teamRouterStruct TeamRouterStruct) HitSaveTeamApi(suite *TeamTestSuite, payload []byte) SaveTeamResponseDto {
	var saveTeamRequestDto SaveTeamRequestDto

	baseStructRef := Utils.BaseStruct{}
	teamName := baseStructRef.GetRandomStringOfGivenLength(10)
	saveTeamRequestDto.Name = teamName
	saveTeamRequestDto.Active = true
	byteValueOfStruct, _ := json.Marshal(saveTeamRequestDto)
	var payloadOfApi string

	if payload != nil {
		payloadOfApi = string(payload)
	} else {
		payloadOfApi = string(byteValueOfStruct)
	}

	resp, err := baseStructRef.MakeApiCall(suite.TeamApiClient, "/orchestrator/team", http.MethodPost, payloadOfApi, true, nil)
	baseStructRef.HandleError(err, "SaveTeamApi")

	teamRouter := teamRouterStruct.UnmarshalGivenResponseBody(resp.Body(), "SaveTeam")
	return teamRouter.saveTeamResponseDto
}

func (teamRouterStruct TeamRouterStruct) HitFetchAllTeamApi(suite *TeamTestSuite) FetchAllTeamResponseDto {
	baseStructRef := Utils.BaseStruct{}
	resp, err := baseStructRef.MakeApiCall(suite.TeamApiClient, "/orchestrator/team", http.MethodGet, "", true, nil)
	baseStructRef.HandleError(err, "FetchAllTeamApi")

	teamRouter := teamRouterStruct.UnmarshalGivenResponseBody(resp.Body(), "FetchAllTeam")
	return teamRouter.fetchAllTeamResponseDto
}

func (teamRouterStruct TeamRouterStruct) HitDeleteTeamApi(suite *TeamTestSuite, byteValueOfStruct []byte) DeleteTeamResponseDto {
	baseStructRef := Utils.BaseStruct{}
	resp, err := baseStructRef.MakeApiCall(suite.TeamApiClient, "/orchestrator/team", http.MethodDelete, string(byteValueOfStruct), true, nil)
	baseStructRef.HandleError(err, "DeleteTeamApi")

	teamRouter := teamRouterStruct.UnmarshalGivenResponseBody(resp.Body(), "DeleteTeam")
	return teamRouter.deleteTeamResponseDto
}

/*
func HitGetTeamByIdApi(id string) *resty.Response {
	resp, err := Base.MakeApiCall("/orchestrator/team/"+id, http.MethodGet, "", true, nil)
	Base.HandleError(err, "GetTeamByIdApiS")
	return resp
}

func HitUpdateTeamApi(byteValueOfStruct []byte) *resty.Response {
	resp, err := Base.MakeApiCall("/orchestrator/team/", http.MethodGet, string(byteValueOfStruct), true, nil)
	Base.HandleError(err, "UpdateTeamApi")
	return resp
}

func HitFetchForAutocompleteApi(byteValueOfStruct []byte) *resty.Response {
	resp, err := Base.MakeApiCall("/orchestrator/team/autocomplete", http.MethodGet, string(byteValueOfStruct), true, nil)
	Base.HandleError(err, "FetchForAutocompleteApi")
	return resp
}*/

func (teamRouterStruct TeamRouterStruct) GetPayLoadForDeleteAPI(id int, name string, isActive bool) []byte {
	var updateTeamDto UpdateTeamRequestDto
	updateTeamDto.Id = id
	updateTeamDto.Name = name
	updateTeamDto.Active = isActive
	byteValueOfStruct, _ := json.Marshal(updateTeamDto)
	return byteValueOfStruct
}

/*
func ValidateGetTeamByIdApi(t *testing.T, saveTeamResponseBody []byte, getTeamByIdResponseBody []byte) {
	var saveTeamResponseDto SaveTeamResponseDto
	json.Unmarshal(saveTeamResponseBody, &saveTeamResponseDto)
	var getTeamByIdResponseDto SaveTeamResponseDto
	json.Unmarshal(getTeamByIdResponseBody, &getTeamByIdResponseDto)

	assert.Equal(t, saveTeamResponseDto.Result.Id, getTeamByIdResponseDto.Result.Id)
	assert.Equal(t, saveTeamResponseDto.Result.Name, getTeamByIdResponseDto.Result.Name)
	assert.Equal(t, saveTeamResponseDto.Result.Active, getTeamByIdResponseDto.Result.Active)
}*/

/*func ValidateDeleteTeamApi(t *testing.T, responseBody []byte) {
	var deleteTeamResponse DeleteTeamResponseDto
	json.Unmarshal(responseBody, &deleteTeamResponse)
	assert.Equal(t, "Project deleted successfully.", deleteTeamResponse.Result)
	assert.Equal(t, "OK", deleteTeamResponse.Status)
}
*/
/*func GetSaveTeamRequestDto() SaveTeamRequestDto {
	var saveTeamRequestDto SaveTeamRequestDto
	teamName := testUtils.GetRandomStringOfGivenLength(10)
	saveTeamRequestDto.Name = teamName
	saveTeamRequestDto.Active = true
	return saveTeamRequestDto
}*/

/*type TeamRouterInterface interface {
	UnmarshalGivenResponseBody(response []byte, apiName string) TeamRouterStruct
}
*/
type TeamRouterStruct struct {
	saveTeamResponseDto     SaveTeamResponseDto
	deleteTeamResponseDto   DeleteTeamResponseDto
	fetchAllTeamResponseDto FetchAllTeamResponseDto
	saveTeamRequestDto      SaveTeamRequestDto
}

func (teamRouterStruct TeamRouterStruct) UnmarshalGivenResponseBody(response []byte, apiName string) TeamRouterStruct {
	switch apiName {
	case "FetchAllTeam":
		json.Unmarshal(response, &teamRouterStruct.fetchAllTeamResponseDto)
	case "SaveTeam":
		json.Unmarshal(response, &teamRouterStruct.saveTeamResponseDto)
	case "DeleteTeam":
		json.Unmarshal(response, &teamRouterStruct.saveTeamResponseDto)
	case "DELETE":
		json.Unmarshal(response, &teamRouterStruct.saveTeamResponseDto)
	}
	return teamRouterStruct
}
