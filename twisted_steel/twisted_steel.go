package main

import (
  "fmt"
  "net/url"

  "philangist.github.com/twisted-steel/api_client"
  "philangist.github.com/restful-micro-framework/sleepy"
  )

var (
  ApiClient = api_client.NewApiClient("https://api.instagram.com/v1")
  AuthParams = map[string]string{
    "access_token": "SOME_KEY_HERE", //get this from environment
  }
  InstagramClient = InstagramApiClient{
    ApiClient: ApiClient,
    AuthParams: AuthParams,
  }
)


type InstagramApiClient struct {
  ApiClient api_client.APIClient
  AuthParams map[string]string
}

type InstagramUserResource struct {
  sleepy.PostNotSupported
  sleepy.PutNotSupported
  sleepy.DeleteNotSupported
}

func check(err error){
  if err != nil{
    panic(err)
  }
}

func (InstagramUserResource) Get(values url.Values) (int, interface{}){
  var err error

  userId := values.Get("user_id")
  if userId == "" {
    return 400, map[string]string{
      "error": "user_id must be specified",
      }
  }

  userData := getInstagramUser(userId)
  return 200, map[string]string{"data": string(userData[:])}
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


func main(){
  instagramResource := new(InstagramUserResource)

  var api = new(sleepy.API)
  api.AddResource(instagramResource, "/users/")
  api.Start(3000)
}
