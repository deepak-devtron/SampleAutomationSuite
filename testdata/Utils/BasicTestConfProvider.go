package Utils

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/caarlos0/env"
	"github.com/go-resty/resty/v2"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"os"
)

type LogInResult struct {
	Token string `json:"token"`
}

type LogInResponse struct {
	Code   int         `json:"code"`
	Status string      `json:"status"`
	Result LogInResult `json:"result"`
}

type ApiErrorDto struct {
	Code   int    `json:"code"`
	Status string `json:"status"`
	Errors []struct {
		Code            string `json:"code"`
		InternalMessage string `json:"internalMessage"`
		UserMessage     string `json:"userMessage"`
	} `json:"errors"`
}

type EnvironmentConfig struct {
	BaseServerUrl string `env:"BASE_SERVER_URL" envDefault:"https://deepak.devtron.info:32443/"`
	LogInUserName string `env:"LOGIN_USERNAME" envDefault:"admin" `
	LogInUserPwd  string `env:"LOGIN_PASSWORD" envDefault:"argocd-server-5495f64f7-pf75s" `
}

type BaseStruct struct {
	logInResult       LogInResult
	logInResponse     LogInResponse
	apiErrorDto       ApiErrorDto
	environmentConfig EnvironmentConfig
}

// MakeApiCall make the api call to the requested url based on http method requested
func (baseStructRef BaseStruct) MakeApiCall(client *resty.Client, apiUrl string, method string, body string, setCookie bool, queryParams map[string]string) (*resty.Response, error) {
	var resp *resty.Response
	var err error
	switch method {
	case "GET":
		if queryParams != nil {
			return baseStructRef.setCookieInApi(client, setCookie).R().SetQueryParams(queryParams).Get(apiUrl)
		}

		return baseStructRef.setCookieInApi(client, setCookie).R().Get(apiUrl)

	case "POST":
		log.Println("Hitting the Post API..." + apiUrl)
		return baseStructRef.setCookieInApi(client, setCookie).R().SetBody(body).Post(apiUrl)

	case "PUT":
		return baseStructRef.setCookieInApi(client, setCookie).R().SetBody(body).Put(apiUrl)
	case "DELETE":
		return baseStructRef.setCookieInApi(client, setCookie).R().SetBody(body).Delete(apiUrl)

	}
	return resp, err
}

/*func (baseStructRef BaseStruct) getRestyClient(client *resty.Client) *resty.Client {
	envConf, _ := baseStructRef.GetEnvironmentConfig()
	//client := suite.teamTestSuite.TeamApiClient
	client.SetBaseURL(envConf.BaseServerUrl)
	if setCookie {
		client.SetCookie(&http.Cookie{Name: "argocd.token",
			Value:  baseStructRef.getAuthToken(suite),
			Path:   "/",
			Domain: envConf.BaseServerUrl})
	}
	return client
}*/

func (baseStructRef BaseStruct) setCookieInApi(client *resty.Client, needCookie bool) *resty.Client {
	envConf, _ := baseStructRef.GetEnvironmentConfig()
	client.SetBaseURL(envConf.BaseServerUrl)
	if needCookie {
		client.SetCookie(&http.Cookie{Name: "argocd.token",
			Value:  baseStructRef.getAuthToken(client),
			Path:   "/",
			Domain: envConf.BaseServerUrl})
	}
	return client
}

// HandleError Log the error and return boolean value indicating whether error occurred or not
func (baseStructRef BaseStruct) HandleError(err error, testName string) {
	if nil != err {
		log.Println("Error occurred while invoking api for test:"+testName, "err", err)
	}
}

func GetByteArrayOfGivenJsonFile(filePath string) ([]byte, error) {
	testDataJsonFile, err := os.Open(filePath)
	if nil != err {
		log.Println("Unable to open the file. Error occurred !!", "err", err)
	}
	log.Println("Opened the given json file successfully !!!")
	defer testDataJsonFile.Close()

	byteValue, err := ioutil.ReadAll(testDataJsonFile)
	return byteValue, err
}

//support function to return auth token after log in
//TODO : if token is valid, don't call api again, error handling in invoking functions
func (baseStructRef BaseStruct) getAuthToken(client *resty.Client) string {
	envConf, _ := baseStructRef.GetEnvironmentConfig()
	jsonString := fmt.Sprintf(`{"username": "%s", "password": "%s"}`, envConf.LogInUserName, envConf.LogInUserPwd)
	log.Println("Trying to get the Auth Token via hitting the session API....")
	//resp, err := BaseStruct.MakeApiCall(client, "/orchestrator/api/v1/session", http.MethodPost, jsonString, false, nil)
	resp, err := client.SetBaseURL(envConf.BaseServerUrl).R().SetBody(jsonString).Post("/orchestrator/api/v1/session")

	baseStructRef.HandleError(err, "getAuthToken")
	var logInResponse LogInResponse
	json.Unmarshal(resp.Body(), &logInResponse)
	log.Println("Getting Auth token from the Response Successfully....")
	return logInResponse.Result.Token
}

func (baseStructRef BaseStruct) GetEnvironmentConfig() (*EnvironmentConfig, error) {
	cfg := &EnvironmentConfig{}
	err := env.Parse(cfg)
	if err != nil {
		return nil, errors.New("could not get config from environment")
	}
	return cfg, err
}

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func (baseStructRef BaseStruct) GetRandomStringOfGivenLength(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

func (baseStructRef BaseStruct) GetRandomNumberOf9Digit() int {
	return 100000000 + rand.Intn(999999999-100000000)
}
