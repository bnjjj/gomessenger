package gomessenger

type Message struct {
	Recipient ObjectWithId `json:"recipient"`
	Message   MessageText  `json:"message"`
}

type MessageText struct {
	Text string `json:"text"`
}

type Attachment struct {
	Type    string  `json:"type"`
	Payload Payload `json:"payload"`
}

type Payload struct {
	URL string `json:"url"`
}

// Send a text message using the Send API.
func (messenger *Messenger) SendMessageText(recipientId string, messageText string) {
	messageData := Message{
		Recipient: ObjectWithId{Id: recipientId},
		Message:   MessageText{Text: messageText},
	}

	messenger.CallSendAPI(messageData)
}
