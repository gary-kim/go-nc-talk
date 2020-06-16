package room

import (
	"errors"
	"github.com/monaco-io/request"
	"time"

	"github.com/gary-kim/go-nc-talk/message"
	"github.com/gary-kim/go-nc-talk/user"
	"github.com/gary-kim/go-nc-talk/constants"
)

type MessageType string

const (
	MessageComment MessageType = "comment"
	MessageSystem MessageType = "system"
	MessageCommand MessageType = "command"
)

type TalkRoom struct {
	User *user.TalkUser
	Token string
}

func (t *TalkRoom) SendMessage(msg string) error {
	url := t.User.NextcloudUrl + constants.BaseEndpoint + "/chat/" + t.Token
	requestParams := map[string]string{
		"message": msg,
	}

	client := request.Client{
		URL: url
		Method: "POST",
		Params: requestParams,
	}
	res, err := client.Do()
	if err != nil {
		return err
	}
	if res.StatusCode() != 201 {
		return errors.New("unexpected return code")
	}
	return nil
}

func (t *TalkRoom) RecieveMessages(c chan message) err {
	url := t.User.NextcloudUrl + constants.BaseEndpoint + "/chat/" + t.Token
	requestParam := map[string]string{
		"lookIntoFuture": "1",
		"limit": "200",
		"timeout": "60",
		"setReadMarker": "1",
		"includeLastKnown": "0",
	}
	go func() {
		for {
			client := request.Client{
				URL: url,
				Params: requestParam,
				Timeout: time.Second * 60,
			}

			res, err := client.Do()
		}
	}()

}

