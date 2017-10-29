package gomessenger

import "net/http"

type SenderAction struct {
	Recipient    ObjectWithId `json:"recipient"`
	SenderAction string       `json:"sender_action"`
}

var (
	ActionTypingOn  = "typing_on"
	ActionTypingOff = "typing_off"
	ActionMarkSeen  = "mark_seen"
)

// SendSenderAction send an action to sender using the Send API.
func (messenger *Messenger) SendSenderAction(recipientId string, actionType string) (*http.Response, error) {
	messageData := SenderAction{
		Recipient:    ObjectWithId{Id: recipientId},
		SenderAction: actionType,
	}

	return messenger.CallSendAPI(messageData)
}
