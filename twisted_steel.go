// This package should just register api handlers as sleepy resources
package main

import (
	"github.com/philangist/golang-twisted-steel/sleepy"
	"github.com/philangist/golang-twisted-steel/api_resources"
)

// api view logic
func main() {
	platformUserResource := new(api_resources.PlatformUserResource)
	var api = new(sleepy.API)
	api.AddResource(platformUserResource, "/users/")
	api.Start(3000)
}
