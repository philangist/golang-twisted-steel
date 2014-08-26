package api_client

import (
  "testing"
  "reflect"
  "runtime"
  "path/filepath"

)

//stoled this from bitly's nsq equal implementation
//https://github.com/bitly/nsq/blob/master/nsqd/test/cluster_test.go#L16
func Equal(t *testing.T, actual, expectation interface{}){
  if !reflect.DeepEqual(actual, expectation){
    _, file, line, _ := runtime.Caller(1)
    t.Logf("\033[31m%s:%d:\n\n\texpected: %#v\n\n\tactual: %#v\033[39m\n\n",
           filepath.Base(file), line, expectation, actual)
    t.FailNow()
  }
}


func TestApiClient(t *testing.T){
  baseUrl := "https://api.gotinder.com"
  // fbOauthToken := "foobar"
  invalidPath := "/doesNotExist"

  ApiClient := NewApiClient(baseUrl)

  Equal(t, ApiClient.baseUrl, baseUrl)
  // Equal(t, ApiClient.fbOauthToken, fbOauthToken)

  resp, err := ApiClient.Get(invalidPath, nil)
  Equal(t, resp, map[string]string{"status":"not found"})
  Equal(t, err, nil)

  queryParameters := map[string]string{
    "foo": "bar",
    "baaz": "quux",
  }

  queryString := "?foo=bar&baaz=quux"

  Equal(
    t,
    ApiClient.buildQueryParamString(queryParameters),
    queryString,
  )
}
