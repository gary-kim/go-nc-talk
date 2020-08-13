package integration

import (
	"flag"
	"testing"

	"github.com/stretchr/testify/assert"

	"gomod.garykim.dev/nc-talk/room"
	"gomod.garykim.dev/nc-talk/user"
)

var instance = flag.String("instance", "", "Nextcloud instance to test against")
var username = flag.String("username", "", "Nextcloud instance bot user")
var password = flag.String("password", "", "Nextcloud instance bot user password")
var token = flag.String("token", "", "Nextcloud instance talk room to test against")

func TestConnection(t *testing.T) {
	checkFlagsSet(t)

	tuser, err := user.NewUser(*instance, *username, *password, &user.TalkUserConfig{})

	assert.NoError(t, err)

	troom, err := room.NewTalkRoom(tuser, *token)

	assert.NoError(t, err)
	assert.NotNil(t, troom)

	err = troom.TestConnection()
	assert.NoError(t, err)
}

func checkFlagsSet(t *testing.T) {
	// Make sure all required flags are set
	if *instance == "" ||
		*username == "" ||
		*password == "" ||
		*token == "" {
		t.Skip("Skipping as required flags are not set")
	}
}
