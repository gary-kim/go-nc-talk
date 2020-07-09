package user

import (
	"encoding/xml"
	"reflect"

	"github.com/monaco-io/request"
)

const (
	ocsCapabilitiesEndpoint = "/ocs/v2.php/cloud/capabilities"
)

// TalkUser represents a user of Nextcloud Talk
type TalkUser struct {
	User         string
	Pass         string
	NextcloudURL string
	capabilities *Capabilities
}

// Capabilities describes the capabilities that the Nextcloud Talk instance is capable of. Visit https://nextcloud-talk.readthedocs.io/en/latest/capabilities/ for more info.
type Capabilities struct {
	Audio                  bool `ocscapability:"audio"`
	Video                  bool `ocscapability:"video"`
	Chat                   bool `ocscapability:"chat"`
	GuestSignaling         bool `ocscapability:"guest-signaling"`
	EmptyGroupRoom         bool `ocscapability:"empty-group-room"`
	GuestDisplayNames      bool `ocscapability:"guest-display-names"`
	MultiRoomUsers         bool `ocscapability:"multi-room-users"`
	ChatV2                 bool `ocscapability:"chat-v2"`
	Favorites              bool `ocscapability:"favorites"`
	LastRoomActivity       bool `ocscapability:"last-room-activity"`
	NoPing                 bool `ocscapability:"no-ping"`
	SystemMessages         bool `ocscapability:"system-messages"`
	MentionFlag            bool `ocscapability:"mention-flag"`
	InCallFlags            bool `ocscapability:"in-call-flags"`
	InviteByMail           bool `ocscapability:"invite-by-mail"`
	NotificationLevels     bool `ocscapability:"notification-levels"`
	InviteGroupsAndMails   bool `ocscapability:"invite-groups-and-mails"`
	LockedOneToOneRooms    bool `ocscapability:"locked-one-to-one-rooms"`
	ReadOnlyRooms          bool `ocscapability:"read-only-rooms"`
	ChatReadMarker         bool `ocscapability:"chat-read-marker"`
	WebinaryLobby          bool `ocscapability:"webinary-lobby"`
	StartCallFlag          bool `ocscapability:"start-call-flag"`
	ChatReplies            bool `ocscapability:"chat-replies"`
	CirclesSupport         bool `ocscapability:"circles-support"`
	AttachmentsAllowed     bool `ocscapability:"config => attachments => allowed"`
	AttachmentsFolder      bool `ocscapability:"config => attachments => folder"`
	ConversationsCanCreate bool `ocscapability:"config => conversations => can-create"`
	ForceMute              bool `ocscapability:"force-mute"`
	ConversationV2         bool `ocscapability:"conversation-v2"`
	ChatReferenceId        bool `ocscapability:"chat-reference-id"`
}

type capabilitiesRequest struct {
	XMLName      xml.Name `xml:"ocs"`
	Capabilities []string `xml:"ocs>data>capabilities>spreed>features>element"`
}

func (t *TalkUser) RequestClient(client request.Client) *request.Client {
	if client.Header == nil {
		client.Header = make(map[string]string)
	}
	if client.Header["OCS-APIRequest"] == "" {
		client.Header["OCS-APIRequest"] = "true"
	}
	if client.Header["Accept"] == "" {
		client.Header["Accept"] = "application/json"
	}
	client.BasicAuth = request.BasicAuth{
		Username: t.User,
		Password: t.Pass,
	}
	return &client
}

func (t *TalkUser) Capabilities() (*Capabilities, error) {
	if t.capabilities != nil {
		return t.capabilities, nil
	}

	client := t.RequestClient(request.Client{
		URL: ocsCapabilitiesEndpoint,
	})
	res, err := client.Do()
	if err != nil {
		return nil, err
	}

	capabilities := &capabilitiesRequest{}
	err = xml.Unmarshal(res.Data, capabilities)
	if err != nil {
		return nil, err
	}

	tr := &Capabilities{}

	c := reflect.ValueOf(tr)
	for i := 0; i < c.NumField(); i++ {
		field := c.Field(i)
		tag := field.Type().Field(0).Tag.Get("ocscapability")

		for _, capability := range capabilities.Capabilities {
			if capability == tag && field.CanSet() {
				field.SetBool(true)
			}
		}
	}

	t.capabilities = tr
	return tr, nil
}
