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
Request:

```bash
$ curl -X GET "localhost:3000/users/?user_id=1574083" # Snoop Doggy Dogg
```

Expected Response:
```javascript
{
  "data":{
    "id":"1574083",
    "username":"snoopdogg"
  }
}
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


### License

Copyright (c) 2011 Phil Opaola

Distributed under an [MIT-style](http://www.opensource.org/licenses/mit-license.php) license.

> Permission is hereby granted, free of charge, to any person obtaining a copy of
> this software and associated documentation files (the "Software"), to deal in
> the Software without restriction, including without limitation the rights to
> use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies
> of the Software, and to permit persons to whom the Software is furnished to do
> so, subject to the following conditions:

> The above copyright notice and this permission notice shall be included in all
> copies or substantial portions of the Software.

> THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
> IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
> FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
> AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
> LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
> OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
> SOFTWARE.
