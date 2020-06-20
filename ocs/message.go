package ocs

// MessageType describes what kind of message a returned Nextcloud Talk message is
type MessageType string

const (
	// MessageComment is a Nextcloud Talk message that is a comment
	MessageComment MessageType = "comment"

	// MessageSystem is a Nextcloud Talk message that is a system
	MessageSystem MessageType = "system"

	// MessageCommand is a Nextcloud Talk message that is a command
	MessageCommand MessageType = "command"
)

// TalkRoomMessage describes the data part of a ocs response for a Talk room message
type TalkRoomMessage struct {
	Message          string      `json:"message"`
	ID               int         `json:"id"`
	ActorID          string      `json:"actorId"`
	ActorDisplayName string      `json:"actorDisplayName"`
	SystemMessage    string      `json:"systemMessage"`
	Timestamp        int         `json:"timestamp"`
	MessageType      MessageType `json:"messageType"`
}

// OCSTalkRoomMessage describes an ocs response for a Talk room message
type OCSTalkRoomMessage struct {
	ocs
	TalkRoomMessage []TalkRoomMessage `json:"data"`
}
