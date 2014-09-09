// JSON types for various platforms. Largely inspired by
// https://github.com/carbocation/go.instagram/blob/master/types.go

package platforms

type InstagramUserSchema struct {
	Bio            string
	FullName       string `json:full_name`
	ID             string
	ProfilePicture string `json:full_name`
	Username       string
	Website        string
}

type InstagramUserResponseSchema struct {
	Meta struct {
		Code int
	}
	Data InstagramUserSchema
}
