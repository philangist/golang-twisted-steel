// This package should just register api handlers as sleepy resources
package main

import (
	"philangist.github.com/restful-micro-framework/sleepy"
	"philangist.github.com/twisted-steel/api_resources"
)

// api view logic
func main() {
	platformUserResource := new(api_resources.PlatformUserResource)
	var api = new(sleepy.API)
	api.AddResource(platformUserResource, "/users/")
	api.Start(3000)
}
