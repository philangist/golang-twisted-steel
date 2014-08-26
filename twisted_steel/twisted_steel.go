package main

// import "philangist.github.com/twisted-steel/sleepy"
import "philangist.github.com/twisted-steel/api_client"

type InstagramApiClient struct {
  ApiClient api_client.APIClient
  AuthParams map[string]string
}

func (this InstagramApiClient) GetMe() (map[string]string, error){
  return this.ApiClient.Get("/me", this.AuthParams)
}


func main(){
  ApiClient := api_client.NewApiClient("https://api.instagram.com")
  InstagramClient := InstagramApiClient{
    ApiClient: ApiClient,
    AuthParams: map[string]string{"client_key": "SOME_KEY_HERE"},
  }
  me, _ := InstagramClient.GetMe()
  print("me is ", me)
}
