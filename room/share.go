package room

import (
	"encoding/json"
	"net/http"

	"github.com/monaco-io/request"

	"gomod.garykim.dev/nc-talk/constants"
	"gomod.garykim.dev/nc-talk/ocs"
)

// ShareFile shares the file at the given path with the talk room
func (t *TalkRoom) ShareFile(path string) (string, error) {
	req := t.User.RequestClient(request.Client{
		URL: constants.FilesSharingEndpoint + "shares",
		Params: map[string]string{
			"shareType": ocs.ShareTypeRoom,
			"path":      path,
			"shareWith": t.Token,
		},
	})
	resp, err := req.Do()
	if err != nil {
		return "", err
	}
	if resp.StatusCode() != http.StatusOK {
		return "", ErrUnexpectedReturnCode
	}
	data := &ocs.ShareReturn{}
	err = json.Unmarshal(resp.Data, data)
	if err != nil {
		return "", err
	}
	return data.OCS.Data.URL, nil
}
