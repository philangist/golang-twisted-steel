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
  SECRET_SUFFIX = "_TOKEN"
  FACEBOOK = "FACEBOOK"
  INSTAGRAM = "INSTAGRAM"
  TWITTER = "TWITTER"
  TUMBLR = "TUMBLR"
  VINE = "VINE"
  )

var (
  InstagramClient = createInstagramClient()
)

func check(err error){
  if err != nil{
    panic(err)
  }
}

func createInstagramClient() InstagramApiClient{
  access_token, err := getTokenForPlatform(INSTAGRAM)
  ApiClient := api_client.NewApiClient("https://api.instagram.com/v1")
  check(err)
  AuthParams := map[string]string{
    "access_token": access_token,
  }
  return InstagramApiClient{
    ApiClient: ApiClient,
    AuthParams: AuthParams,
  }
}

func getTokenForPlatform(PlatformName string) (string, error) {
  var key, value string

  key = fmt.Sprintf("%s%s", PlatformName, SECRET_SUFFIX)
  value = os.Getenv(key)

  if value == "" {
    return "", errors.New(fmt.Sprintf(
      "Could not get value for %s from environment", key))
  }
  return value, nil
}

type PlatformUser struct{
  UserName, UserId string
}

type PlatformApiClient interface {
  getMe() PlatformUser
  getFollowers() []PlatformUser
  getFollowing() []PlatformUser
  followUser() (bool, error) //accepts PlatformUser as argument
  unfollowUser() (bool, error) //accepts PlatformUser as argument
}

type InstagramApiClient struct {
  ApiClient api_client.APIClient
  AuthParams map[string]string
}

func (this InstagramApiClient) GetUser(userId string) ([]byte, error){
  return this.ApiClient.Get(
    fmt.Sprintf("/users/%s", userId),
    this.AuthParams,
  )
}


func getInstagramUser(userId string) []byte{
  userBuffer, err := InstagramClient.GetUser(userId)
  check(err)
  return userBuffer
}


type InstagramUserResource struct {
  sleepy.PostNotSupported
  sleepy.PutNotSupported
  sleepy.DeleteNotSupported
}

// api view logic

func (InstagramUserResource) Get(values url.Values) (int, interface{}){
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
  instagramResource := new(InstagramUserResource)
  var api = new(sleepy.API)
  api.AddResource(instagramResource, "/users/")
  api.Start(3000)
}
