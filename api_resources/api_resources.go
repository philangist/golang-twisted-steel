package api_resources

import (
	"fmt"
	"net/url"

	"github.com/philangist/golang-twisted-steel/sleepy"
	"github.com/philangist/golang-twisted-steel/platforms"
	"github.com/philangist/golang-twisted-steel/utils"
)

type PlatformUserResource struct {
	sleepy.PostNotSupported
	sleepy.PutNotSupported
	sleepy.DeleteNotSupported
}

func (PlatformUserResource) Get(values url.Values) (int, interface{}) {
	userId := values.Get("user_id")
	if userId == "" {
		return 400, map[string]string{
			"error": "user_id must be specified",
		}
	}

	instagramApiClient, err := platforms.GetPlatformApiClient(
		platforms.INSTAGRAM)
	utils.Check(err)
	platformUser, err := instagramApiClient.GetUser(userId)

	utils.Check(err)
	userFollowers, err := instagramApiClient.GetFollowers(userId)
	utils.Check(err)
	fmt.Println("userFollowers is %s", userFollowers)

	return 200, map[string]map[string]string{
		"data": {
			"id":       platformUser.ID,
			"username": platformUser.Username,
		},
	}
}
