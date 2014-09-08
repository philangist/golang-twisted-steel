// uses api_client common logic to build platform specific logic
package platforms

import (
  "encoding/json"
  "errors"
  "fmt"
  "os"

  "philangist.github.com/twisted-steel/api_client"
  "philangist.github.com/twisted-steel/utils"
)

const (
  FACEBOOK = "FACEBOOK"
  INSTAGRAM = "INSTAGRAM"
  TWITTER = "TWITTER"
  TUMBLR = "TUMBLR"
  VINE = "VINE"
  SECRET_SUFFIX = "_TOKEN"
  )

var (
  PlatformApiBaseUrls = map[string]string{
    INSTAGRAM: "https://api.instagram.com/v1",
    /**
    We don't really care about these platforms yet
    FACEBOOK: "http://api.facebook.com/restserver.php",
    TWITTER: "http://twitter.com/statuses",
    TUMBLR: "http://www.tumblr.com/api",
    VINE: "https://api.vineapp.com",
    **/
  }
)

type PlatformUser struct{
  Username, ID string
}


type PlatformApiClient interface {
  GetMe() (PlatformUser, error)
  GetUser(string) (PlatformUser, error)
  GetFollowers(string) ([]PlatformUser, error)
  GetFollowing(string) ([]PlatformUser, error)
  FollowUser(PlatformUser) (bool, error) //accepts PlatformUser as argument
  UnfollowUser(PlatformUser) (bool, error) //accepts PlatformUser as argument
}

type InstagramApiClient struct {
  ApiClient api_client.APIClient
  AuthParams map[string]string
}


func GetPlatformApiClient (platform string) (PlatformApiClient, error) {
  apiBaseUrl := PlatformApiBaseUrls[platform]
  if apiBaseUrl == "" {
    return nil, errors.New(fmt.Sprintf(
      "Tried to instantiate PlatformApiClient with " +
      "invalid platform %s", platform))
  }
  accessToken := getTokenForPlatform(platform)
  apiClient := api_client.NewApiClient(apiBaseUrl)
  authParams := map[string]string{
    "access_token": accessToken,
  }
  return InstagramApiClient{
    ApiClient: apiClient,
    AuthParams: authParams,
  }, nil
}

func getKeyFromEnv(key string) (string, error) {
  value := os.Getenv(key)
  if value == "" {
    return "", errors.New(fmt.Sprintf(
      "Could not get value for %s from environment", key))
  }
  return value, nil
}

func getTokenForPlatform(PlatformName string) (string) {
  platformKey := fmt.Sprintf("%s%s", PlatformName, SECRET_SUFFIX)
  token, err := getKeyFromEnv(platformKey)
  utils.Check(err)
  return token
}

func (this InstagramApiClient) GetMe() (PlatformUser, error){
  // pass
  return PlatformUser{"foo", "bar"}, nil
}

func (this InstagramApiClient) GetUser(userId string) (PlatformUser, error){
  resp, err := this.ApiClient.Get(
    fmt.Sprintf("/users/%s", userId),
    this.AuthParams,
  )
  utils.Check(err)

  var dat InstagramUserResponseSchema
  err = json.Unmarshal(resp, &dat)
  utils.Check(err)

  platformUser := PlatformUser{dat.Data.Username, dat.Data.ID}
  return platformUser, nil
}

func (this InstagramApiClient) GetFollowers(userId string) ([]PlatformUser, error){
  /** pass
  resp, _ := this.ApiClient.Get(
    fmt.Sprintf("/users/%s/followed-by", userId),
    this.AuthParams,
  )**/
  dummyUser := PlatformUser{"foo", "bar"}
  return []PlatformUser{dummyUser,}, nil
}

func (this InstagramApiClient) GetFollowing(userId string) ([]PlatformUser, error){
  /** pass
  resp, _ := this.ApiClient.Get(
    fmt.Sprintf("/users/%s/follows", userId),
    this.AuthParams,
  )**/
  dummyUser := PlatformUser{"foo", "bar"}
  return []PlatformUser{dummyUser,}, nil
}

func (this InstagramApiClient) FollowUser(user PlatformUser) (bool, error){
  // pass
  return true, nil
}

func (this InstagramApiClient) UnfollowUser(user PlatformUser) (bool, error){
  // pass
  return true, nil
}
