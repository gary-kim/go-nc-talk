package talk

import (
	"github.com/gary-kim/go-nc-talk/room"
	"github.com/gary-kim/go-nc-talk/user"
)

// NewUser returns a TalkUser instance
func NewUser(url string, username string, password string) *user.TalkUser {
	return &user.TalkUser{
		NextcloudURL: url,
		User:         username,
		Pass:         password,
	}
}

// NewRoom returns a new TalkRoom instance
func NewRoom(tuser *user.TalkUser, token string) *room.TalkRoom {
	tr := &room.TalkRoom{
		User:  tuser,
		Token: token,
	}
	return tr
}
