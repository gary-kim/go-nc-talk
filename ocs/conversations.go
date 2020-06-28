package ocs

type Type int

const (
	ConversationOneToOne Type = 1
	ConversationGroup Type = 2
	ConversationPublic Type = 3
	ConversationChangelog Type = 4
)

type Conversations struct {
	Token string `json:"token"`
	Type Type `json:"type"`
	Name string `json:"name"`
	DisplayName string `json:"displayName"`
	ParticipantType int `json:"participantType"`

}
