package gomessenger

import "net/http"

type AccountLinking struct {
	Recipient ObjectWithId          `json:"recipient"`
	Message   MessageAccountLinking `json:"message"`
}

type MessageAccountLinking struct {
	Attachment AttachmentAccountLinking `json:"attachment"`
}

type AttachmentAccountLinking struct {
	Type    string                `json:"type"`
	Payload PayloadAccountLinking `json:"payload"`
}

type PayloadAccountLinking struct {
	TemplateType string                  `json:"template_type"`
	Text         string                  `json:"text"`
	Buttons      []ButtonsAccountLinking `json:"buttons"`
}

type ButtonsAccountLinking struct {
	Type string `json:"type"`
	Url  string `json:"url"`
}

// Send a message with the account linking call-to-action
func (messenger *Messenger) SendAccountLinking(recipientId string, url string, text string) (*http.Response, error) {
	messageData := AccountLinking{
		Recipient: ObjectWithId{Id: recipientId},
		Message: MessageAccountLinking{
			Attachment: AttachmentAccountLinking{
				Type: "template",
				Payload: PayloadAccountLinking{
					TemplateType: "button",
					Text:         text,
					Buttons: []ButtonsAccountLinking{
						{Type: "account_link", Url: url},
					},
				},
			},
		},
	}

	return messenger.CallSendAPI(messageData)
}
