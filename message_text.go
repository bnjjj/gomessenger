package gomessenger

import "net/http"

type Message struct {
	Recipient ObjectWithId `json:"recipient"`
	Message   MessageText  `json:"message"`
}

type MessageText struct {
	Text string `json:"text"`
}

// Send a text message using the Send API.
func (messenger *Messenger) SendMessageText(recipientId string, messageText string) (*http.Response, error) {
	messageData := Message{
		Recipient: ObjectWithId{Id: recipientId},
		Message:   MessageText{Text: messageText},
	}

	return messenger.CallSendAPI(messageData)
}
