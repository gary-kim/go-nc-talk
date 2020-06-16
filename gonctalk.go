package gonctalk

import (
	"github.com/gary-kim/go-nc-talk/user"
)

func New(url string, user string, password string) *user.TalkUser {
	return &user.TalkUser{
		NextcloudUrl: url,
		User: user,
		Pass: password,
	}
}




