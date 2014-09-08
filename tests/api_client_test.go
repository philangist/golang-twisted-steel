package tests

import (
  "testing"
  "reflect"
  "runtime"
  "path/filepath"

  "philangist.github.com/twisted-steel/api_client"
  "philangist.github.com/twisted-steel/utils"
)



func TestApiClient(t *testing.T){
  baseUrl := "https://api.instagram.com"
  invalidPath := "/doesNotExist"

  ApiClient := api_client.NewApiClient(baseUrl)

  utils.Equal(t, ApiClient.baseUrl, baseUrl)

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
