package api_client

import (
  "endcoding/json"
  "fmt"
  "http"
  "ioutil"
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
  base_url string
  auth_header string
}

func New(base_url string, fb_auth_token string) *APIClient {
  // set a timeout on this client
  client = &http.Client{
    Timeout: time.Duration(HTTP_REQUEST_TIMEOUT)*time.Second
  }

  return &APIClient{
    base_url: base_url,
    auth_header: auth_header,
    fb_oauth_token: fb_oauth_token,
    session_auth_token: nil,
    client: Client
  }
}

func (this APIClient) buildQueryParamString(query_params) string {
  query_params_string := ''

  if query_params != nil {
    idx = 0
    for key, value := range query_params{
      if idx == 0:
        key = fmt.Sprintf('?%s', key)

      query_param_string = fmt.Sprintf('%s%s&', key, val)
      query_params_string += query_param_string
      idx += 1
    }
    query_params_string = query_params_string[:-1]  // strip off trailing '&'
  }

  return query_params_string
}

func (this APIClient) httpRequest(method string, path string) (map, error){
  if this.auth_header == nil {
    return nil, (
      "Client can not be used to make requests "
      "unless authenticated session toke is available"
      )
  }

  full_path := this.base_url + path

  switch method {
    case GET:
      resp, err := http.Get(full_path) // set headers on request
    case POST:
      resp, err := http.Post(full_path) // set headers on request
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
  parsed_response, err = json.Unmarshal(data)

  return parsed_response, err

}

func (this APIClient) Get(path string, query_params map) (map, error) {
  // Url encodes params, hits endpoint, returns map{}
  path =  path + this.buildQueryParamString(query_params)
  return this.httpRequest(GET, path+query_params_string)
}

func (this APIClient) Post(path string, query_params map) (map, error) {
  path =  path + this.buildQueryParamString(query_params)
  return this.httpRequest(POST, path)
}
