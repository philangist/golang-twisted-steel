package main

import (
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

type PlatformUser struct{
  UserName, UserId string
}

type PlatformApiClient interface {
  GetMe() (PlatformUser, error)
  GetUser(string) ([]byte, error)
  GetFollowers() ([]PlatformUser, error)
  GetFollowing() ([]PlatformUser, error)
  FollowUser(PlatformUser) (bool, error) //accepts PlatformUser as argument
  UnfollowUser(PlatformUser) (bool, error) //accepts PlatformUser as argument
}

type InstagramApiClient struct {
  ApiClient api_client.APIClient
  AuthParams map[string]string
}

func (this InstagramApiClient) GetMe() (PlatformUser, error){
  // pass
  return PlatformUser{"foo", "bar"}, nil
}


func (this InstagramApiClient) GetUser(userId string) ([]byte, error){
  return this.ApiClient.Get(
    fmt.Sprintf("/users/%s", userId),
    this.AuthParams,
  )
}


func (this InstagramApiClient) GetFollowers() ([]PlatformUser, error){
  // pass
  dummyUser := PlatformUser{"foo", "bar"}
  return []PlatformUser{dummyUser,}, nil
}


func (this InstagramApiClient) GetFollowing() ([]PlatformUser, error){
  // pass
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


func getInstagramUser(userId string) []byte{
  instagramApiClient, err := GetPlatformApiClient(INSTAGRAM)
  check(err)
  userBuffer, err := instagramApiClient.GetUser(userId)
  check(err)
  return userBuffer
}


type PlatformUserResource struct {
  sleepy.PostNotSupported
  sleepy.PutNotSupported
  sleepy.DeleteNotSupported
}

// api view logic

func (PlatformUserResource) Get(values url.Values) (int, interface{}){
  userId := values.Get("user_id")
  if userId == "" {
    return 400, map[string]string{
      "error": "user_id must be specified",
      }
  }

  userData := getInstagramUser(userId)
  return 200, map[string]string{"data": string(userData[:])}
}


func main(){
  platformUserResource := new(PlatformUserResource)
  var api = new(sleepy.API)
  api.AddResource(platformUserResource, "/users/")
  api.Start(3000)
}
