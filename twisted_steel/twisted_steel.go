package main

import (
  "encoding/json"
  "errors"
  "fmt"
  "net/url"
  "os"

  "philangist.github.com/twisted-steel/api_client"
  "philangist.github.com/restful-micro-framework/sleepy"
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
    We don't care about these platforms yet
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

//-----------------------------------------------------------------
// https://github.com/carbocation/go.instagram/blob/master/types.go
type InstagramUserSchema struct{
  Bio string
  FullName string  `json:full_name`
  ID string
  ProfilePicture string `json:full_name`
  Username string
  Website string
}

type InstagramUserResponseSchema struct{
  Meta struct {
    Code int
  }
  Data InstagramUserSchema
}
//
//-----------------------------------------------------------------

type PlatformUserResource struct {
  sleepy.PostNotSupported
  sleepy.PutNotSupported
  sleepy.DeleteNotSupported
}

func check(err error){
  if err != nil{
    panic(err)
  }
}

func GetPlatformApiClient (platform string) (PlatformApiClient, error) {
  apiBaseUrl := PlatformApiBaseUrls[platform]
  if apiBaseUrl == "" {
    return nil, errors.New(fmt.Sprintf(
      "Tried to instantiate PlatformApiClient with" +
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
  check(err)
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
  check(err)

  var dat InstagramUserResponseSchema
  err = json.Unmarshal(resp, &dat)
  check(err)

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


// api view logic

func (PlatformUserResource) Get(values url.Values) (int, interface{}){
  userId := values.Get("user_id")
  if userId == "" {
    return 400, map[string]string{
      "error": "user_id must be specified",
      }
  }

  instagramApiClient, err := GetPlatformApiClient(INSTAGRAM)
  check(err)
  platformUser, err := instagramApiClient.GetUser(userId)

  check(err)
  userFollowers, err := instagramApiClient.GetFollowers(userId)
  check(err)
  fmt.Println("userFollowers is %s", userFollowers)

  return 200, map[string]map[string]string {
    "data": {
      "id": platformUser.ID,
      "username": platformUser.Username,
      },
    }
}

func main(){
  platformUserResource := new(PlatformUserResource)
  var api = new(sleepy.API)
  api.AddResource(platformUserResource, "/users/")
  api.Start(3000)
}
