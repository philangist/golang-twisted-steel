package tests

import (
	"testing"

	"philangist.github.com/twisted-steel/api_client"
	"philangist.github.com/twisted-steel/utils"
)

func TestApiClient(t *testing.T) {
	baseUrl := "https://api.instagram.com"
	// invalidPath := "/doesNotExist"

	ApiClient := api_client.NewApiClient(baseUrl)

	utils.Equal(t, ApiClient.BaseUrl, baseUrl)

	// resp, err := ApiClient.Get(invalidPath, nil)
	//utils.Equal(t, resp, map[string]string{"status":"not found"})
	// utils.Equal(t, err, nil)

	queryParameters := map[string]string{
		"foo":  "bar",
		"baaz": "quux",
	}

	queryString := "?foo=bar&baaz=quux"

	utils.Equal(
		t,
		ApiClient.BuildQueryParamString(queryParameters),
		queryString,
	)
}
