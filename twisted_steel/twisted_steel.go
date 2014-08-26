package main

// import "philangist.github.com/twisted-steel/sleepy"
import (
  "philangist.github.com/twisted-steel/api_client"
  "encoding/json"
  )

type InstagramApiClient struct {
  ApiClient api_client.APIClient
  AuthParams map[string]string
}

func (this InstagramApiClient) GetMe() ([]byte, error){
  return this.ApiClient.Get("/me", this.AuthParams)
}


func main(){
  ApiClient := api_client.NewApiClient("https://api.instagram.com/v1")
  AuthParams := map[string]string{
    "access_token": "SOME_KEY_HERE",
  }
  InstagramClient := InstagramApiClient{
    ApiClient: ApiClient,
    AuthParams: AuthParams,
  }
  data, err := InstagramClient.ApiClient.Get("/users/1574083", AuthParams)

  if err != nil {
    panic(err)
  }

  var dat map[string]interface{}
  err = json.Unmarshal(data[:], &dat)
  if err !=nil {
    panic(err)
  }

  meta := dat["meta"].(map[string]interface{})
  print("dat is ", string(data))
  print("meta is ", meta["code"].(float64))
}
