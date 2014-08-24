package main

// import "philangist.github.com/twisted-steel/sleepy"
import "philangist.github.com/twisted-steel/api_client"

func main(){
  ApiClient := api_client.New("https://api.gotinder.com", "foobar")
  queryParams := map[string]string{}
  me, _ := ApiClient.Get("/me", queryParams)
  print("me is ", me)]
}
