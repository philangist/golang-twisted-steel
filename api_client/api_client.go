package api_client

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

const (
	GET                  = "GET"
	POST                 = "POST"
	HTTP_REQUEST_TIMEOUT = 30
)

type HeaderParameter struct {
	key   string
	value string
}

type RequestPreprocessor struct {
	HeaderParameters []HeaderParameter
}

func (this RequestPreprocessor) Preprocess(req *http.Request) *http.Request {
	for _, headerParameter := range this.HeaderParameters {
		req.Header.Set(
			headerParameter.key, headerParameter.value,
		)
	}
	return req
}

type APIClient struct {
	BaseUrl             string
	Client              http.Client
	RequestPreprocessor RequestPreprocessor
}

func NewApiClient(baseUrl string) APIClient {
	client := http.Client{
		Timeout: time.Duration(HTTP_REQUEST_TIMEOUT) * time.Second,
	}

	requestPreprocessor := RequestPreprocessor{
		[]HeaderParameter{},
	}

	return APIClient{
		BaseUrl:             baseUrl,
		Client:              client,
		RequestPreprocessor: requestPreprocessor,
	}
}

func (this APIClient) BuildQueryParamString(queryParams map[string]string) string {
	queryParamsString := ""

	if queryParams != nil {
		idx := 0
		for key, value := range queryParams {
			if idx == 0 {
				key = fmt.Sprintf("?%s", key)
			}
			idx += 1

			queryParamString := fmt.Sprintf("%s=%s&", key, value)
			queryParamsString += queryParamString
		}
		//strip off trailing '&'
		queryParamsString = queryParamsString[:len(queryParamsString)-1]
	}

	return queryParamsString
}

func (this APIClient) Do(req *http.Request) (*http.Response, error) {
	req = this.RequestPreprocessor.Preprocess(req)
	fmt.Println("req.URL is %s", req.URL)
	return this.Client.Do(req)
}

func (this APIClient) httpRequest(method string, path string) (
	[]byte, error) {

	var resp *http.Response
	fullPath := this.BaseUrl + path

	switch method {
	case GET:
		req, err := http.NewRequest(GET, fullPath, nil)
		if err != nil {
			return nil, err
		}

		resp, err = this.Do(req)
		if err != nil {
			return nil, err
		}

	case POST:
		req, err := http.NewRequest(POST, fullPath, nil)
		if err != nil {
			return nil, err
		}

		resp, err = this.Do(req)
		if err != nil {
			return nil, err
		}
	default:
		return nil, errors.New(fmt.Sprintf(
			"Only %s and %s are supported by this client", GET, POST))
	}

	defer resp.Body.Close()
	return ioutil.ReadAll(resp.Body)
}

func (this APIClient) Get(path string, queryParams map[string]string) (
	[]byte, error) {
	path = path + this.BuildQueryParamString(queryParams)
	return this.httpRequest(GET, path)
}

func (this APIClient) Post(path string, queryParams map[string]string) (
	[]byte, error) {
	path = path + this.BuildQueryParamString(queryParams)
	return this.httpRequest(POST, path)
}
