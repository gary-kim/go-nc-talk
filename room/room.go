package room

import (
	"context"
	"encoding/json"
	"errors"
	"io/ioutil"
	"time"

	"github.com/monaco-io/request"

	"gomod.garykim.dev/nc-talk/constants"
	"gomod.garykim.dev/nc-talk/ocs"
	"gomod.garykim.dev/nc-talk/user"
)

// TalkRoom represents a room in Nextcloud Talk
type TalkRoom struct {
	User  *user.TalkUser
	Token string
}

// SendMessage sends a message in the Talk room
func (t *TalkRoom) SendMessage(msg string) (*ocs.TalkRoomMessage, error) {
	url := t.User.NextcloudURL + constants.BaseEndpoint + "/chat/" + t.Token
	requestParams := map[string]string{
		"message": msg,
	}

	client := t.User.RequestClient(request.Client{
		URL:    url,
		Method: "POST",
		Params: requestParams,
	})
	res, err := client.Do()
	if err != nil {
		return nil, err
	}
	if res.StatusCode() != 201 {
		return nil, errors.New("unexpected return code")
	}
	msgInfo := &ocs.OCSTalkRoomMessage{}
	err = json.Unmarshal(res.Data, msgInfo)
	return &msgInfo.TalkRoomMessage[0], err
}

// ReceiveMessages starts watching for new messages
func (t *TalkRoom) ReceiveMessages(ctx context.Context) (chan ocs.TalkRoomMessage, error) {
	c := make(chan ocs.TalkRoomMessage)
	url := t.User.NextcloudURL + constants.BaseEndpoint + "/chat/" + t.Token
	requestParam := map[string]string{
		"lookIntoFuture":   "1",
		"limit":            "200",
		"timeout":          "60",
		"setReadMarker":    "1",
		"includeLastKnown": "0",
	}
	err := t.TestConnection()
	if err != nil {
		return nil, err
	}
	lastKnown := ""
	go func() {
		for {
			if ctx.Err() != nil {
				return
			}
			if lastKnown != "" {
				requestParam["lastKnownMessageId"] = lastKnown
			}
			client := t.User.RequestClient(request.Client{
				URL:     url,
				Params:  requestParam,
				Timeout: time.Second * 60,
			})

			res, err := client.Resp()
			if err != nil {
				continue
			}
			if res.StatusCode == 200 {
				lastKnown = res.Header.Get("X-Chat-Last-Given")
				message := ocs.OCSTalkRoomMessage{}
				data, err := ioutil.ReadAll(res.Body)
				if err != nil {
					continue
				}
				err = json.Unmarshal(data, &message)
				if err != nil {
					continue
				}
				for _, msg := range message.TalkRoomMessage {
					c <- msg
				}
			}
		}
	}()
	return c, nil
}

func (t *TalkRoom) TestConnection() error {
	url := t.User.NextcloudURL + constants.BaseEndpoint + "/chat/" + t.Token
	requestParam := map[string]string{
		"lookIntoFuture":   "0",
		"limit":            "1",
		"timeout":          "30",
		"setReadMarker":    "0",
		"includeLastKnown": "0",
	}
	client := t.User.RequestClient(request.Client{
		URL:     url,
		Params:  requestParam,
		Timeout: time.Second * 30,
	})

	res, err := client.Do()
	if err != nil {
		return err
	}
	switch res.StatusCode() {
	case 200:
		return nil
	case 304:
		return nil
	case 404:
		return errors.New("room could not be found")
	case 412:
		return errors.New("room is in lobby mode but user is not a moderator")
	}
	return errors.New("unknown return code")
}
