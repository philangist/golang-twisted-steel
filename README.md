[![Build Status](https://travis-ci.org/philangist/golang-twisted-steel.svg)](https://travis-ci.org/philangist/golang-twisted-steel)


Twisted Steel
==
RESTful backend for an application to manage your followers on the social medias.


### Running the server

__Setup__
```bash
$ touch credentials.env
$ emacs credentials.env
$ source credentials.env
$ go run twisted_steel.go &
```
N.B: `credentials.env` should be a config file with OAuth access tokens for various social media platforms:

__Make a test request__

```bash
$ curl -X GET "localhost:3000/users/?user_id=1574083" # Snoop Doggy Dogg
```

__Stop the server__
```bash
$ fg
$ [Ctrl-c]
```

```
export INSTAGRAM_TOKEN="INSTAGRAM_AUTH_PAYLOAD"
export FACEBOOK_TOKEN="FACEBOOK_AUTH_PAYLOAD"
export TWITTER_TOKEN="TWITTER_AUTH_PAYLOAD"
export TUMBLR_TOKEN="TUMBLR_AUTH_PAYLOAD"
export VINE_TOKEN="VINE_AUTH_PAYLOAD"
```

__Tests__
```bash
$ go test -v ./tests
```
