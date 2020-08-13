package talk

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"gomod.garykim.dev/nc-talk/room"
)

func TestNewTalkRoom(t *testing.T) {
	username := "username"
	password := "password"
	url := "http://localhost:8080"
	for i, test := range []struct {
		token string
		ok    bool
	}{
		{"d6zoa2zs", true},
		{"h6xo3ba9", true},
		{"", false},
	} {
		tr, err := room.NewTalkRoom(NewUser(url, username, password), test.token)
		if test.ok {
			assert.NoError(t, err, "Test %d: NewTalkRoom error is not nil", i)
			assert.NotNil(t, tr, "Test %d: NewTalkRoom return is nil", i)
		} else {
			assert.Error(t, err, "Test %d: NewTalkRoom error is nil", i)
			assert.Nil(t, tr, "Test %d: NewTalkRoom return is not nil", i)
		}
	}
}
