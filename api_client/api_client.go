package api_client

import (
  "encoding/json"
  "fmt"
  "io/ioutil"
  "net/http"
  "time"
  "errors"
)

// Tinder api client only supports get and post requests
const (
  GET = "GET"
  POST = "POST"
  HTTP_REQUEST_TIMEOUT = 30
)

// interface to http library, returns native go types from parsed
// http responses, sets auth parameters before every request, and handles
// boilerplate so that TinderApiClient can just work
type APIClient struct {
  baseUrl string
  fbOauthToken string
  sessionAuthToken string
  client  http.Client
}

func NewApiClient(baseUrl string, fbOauthToken string) *APIClient {
  // set a timeout on this client
  client := http.Client{
    Timeout: time.Duration(HTTP_REQUEST_TIMEOUT) * time.Second,
  }

  return &APIClient{
    baseUrl: baseUrl,
    fbOauthToken: fbOauthToken,
    sessionAuthToken: "SOME-SECRET-HERE",
    client: client,
  }
}

func (this APIClient) buildQueryParamString(queryParams map[string]string) (
                                            string) {
  queryParamsString := ""

  if queryParams != nil {
    idx := 0
    for key, value := range queryParams{
      if idx == 0{
        key = fmt.Sprintf("?%s", key)
      }
      idx += 1

      queryParamString := fmt.Sprintf("%s=%s&", key, value)
      queryParamsString += queryParamString
    }
    //strip off trailing '&'
    queryParamsString = queryParamsString[:len(queryParamsString) - 1]
  }

  return queryParamsString
}

func (this APIClient) Do(req *http.Request) (*http.Response, error){
  req.Header.Set("X-Auth-Token", this.sessionAuthToken)
  return this.client.Do(req)
}

func (this APIClient) httpRequest(method string, path string) (
                                  map[string]string, error) {
  if this.sessionAuthToken == "" {
    return nil, errors.New(
      "Client can not be used to make requests unless authenticated session" +
      " token is available")
  }

  var resp *http.Response
  var err error

  fullPath := this.baseUrl + path


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
  data, err := ioutil.ReadAll(resp.Body)

  if err != nil{
    return nil, err
  }


  var parsedResponse map[string]string
  err = json.Unmarshal(data, &parsedResponse)

  if err != nil {
    return nil, err
  }

  return parsedResponse, err

}

func (this APIClient) Get(path string, queryParams map[string]string) (
                          map[string]string, error) {
  // Url encodes params, hits endpoint
  path =  path + this.buildQueryParamString(queryParams)
  return this.httpRequest(GET, path)
}

func (this APIClient) Post(path string, queryParams map[string]string) (
                           map[string]string, error) {
  path =  path + this.buildQueryParamString(queryParams)
  return this.httpRequest(POST, path)
}
