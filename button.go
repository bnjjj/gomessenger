package gomessenger

import "net/http"

type MessageButton struct {
	Recipient ObjectWithId      `json:"recipient"`
	Message   MessageTextButton `json:"message"`
}

type MessageTextButton struct {
	Attachment AttachmentButton `json:"attachment"`
}

type AttachmentButton struct {
	Type    string        `json:"type"`
	Payload PayloadButton `json:"payload"`
}

type PayloadButton struct {
	TemplateType string   `json:"template_type"`
	Text         string   `json:"text"`
	Buttons      []Button `json:"buttons"`
}

type Button struct {
	Type    string `json:"type"`
	URL     string `json:"url,omitempty"`
	Payload string `json:"payload,omitempty"`
	Title   string `json:"title,omitempty"`
}

// Send a message with buttons
func (messenger *Messenger) SendButtonsMessage(recipientId string, text string, buttons []Button) (*http.Response, error) {
	messageData := MessageButton{
		Recipient: ObjectWithId{Id: recipientId},
		Message: MessageTextButton{
			Attachment: AttachmentButton{
				Type: "template",
				Payload: PayloadButton{
					TemplateType: "button",
					Text:         text,
					Buttons:      buttons,
				},
			},
		},
	}

	return messenger.CallSendAPI(messageData)
}
