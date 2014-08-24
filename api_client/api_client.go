package api_client

import (
  "endcoding/json"
  "fmt"
  "http"
  "ioutil"
  "http"
  "time"
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

func New(baseUrl string, fb_auth_token string) *APIClient {
  // set a timeout on this client
  client = &http.Client{
    Timeout: time.Duration(HTTP_REQUEST_TIMEOUT)*time.Second
  }

  return &APIClient{
    baseUrl: baseUrl,
    fbOauthToken: fbOauthToken,
    sessionAuthToken: 'SOME-SECRET-HERE',
    client: Client
  }
}

func (this APIClient) buildQueryParamString(queryParams) string {
  queryParamsString := ''

  if queryParams != nil {
    idx = 0
    for key, value := range queryParams{
      if idx == 0:
        key = fmt.Sprintf('?%s', key)

      queryParamString = fmt.Sprintf('%s%s&', key, val)
      queryParamsString += queryParamString
      idx += 1
    }
    queryParamsString = queryParamsString[:-1]  // strip off trailing '&'
  }

  return queryParamsString
}

func (this APIClient) httpRequest(method string, path string) (map, error){
  if this.sessionAuthToken == nil {
    return nil, (
      "Client can not be used to make requests "
      "unless authenticated session toke is available"
      )
  }

  fullPath := this.baseUrl + path

  switch method {
    case GET:
      req, _ = http.NewRequest(GET, fullPath)
      req.Header.Set("X-Auth-Token", this.sessionAuthToken)
      resp, err := this.client.Do(req)
    case POST:
      req, _ = http.NewRequest(POST, fullPath)
      req.Header.Set("X-Auth-Token", this.sessionAuthToken)
      resp, err := this.client.Do(req)
    default:
      return nil, fmt.Sprintf(
        "Only %s and %s are supported by this client", GET, POST)
  }

  if err != nil {
    return nil, err
  }

  defer resp.Body.Close()
  data, err = ioutil.ReadAll(resp.Body)

  if err != nil{
    return nil, err
  }
  parsedResponse, err = json.Unmarshal(data)

  return parsedResponse, err

}

func (this APIClient) Get(path string, queryParams map) (map, error) {
  // Url encodes params, hits endpoint, returns map{}
  path =  path + this.buildQueryParamString(queryParams)
  return this.httpRequest(GET, path+queryParamsString)
}

func (this APIClient) Post(path string, queryParams map) (map, error) {
  path =  path + this.buildQueryParamString(queryParams)
  return this.httpRequest(POST, path)
}
