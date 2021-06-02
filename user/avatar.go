package user

import (
	"net/url"
	"strconv"
)

// GetAvatarURL returns the URL to the avatar with the given userID and size in pixels
func (t *TalkUser) GetAvatarURL(userID string, size int) (avatarURL string) {
	return t.NextcloudURL + "/index.php/avatar/" + url.PathEscape(userID) + "/" + strconv.Itoa(size)
}
