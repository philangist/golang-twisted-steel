// This package should just register api handlers as sleepy resources
package main

import (
  "philangist.github.com/twisted-steel/api_resources"
  "philangist.github.com/restful-micro-framework/sleepy"
  )


// api view logic
func main(){
  platformUserResource := new(api_resources.PlatformUserResource)
  var api = new(sleepy.API)
  api.AddResource(platformUserResource, "/users/")
  api.Start(3000)
}
