package gomessenger

import "net/http"

type MessageGeneric struct {
	Recipient ObjectWithId       `json:"recipient"`
	Message   MessageTextGeneric `json:"message"`
}

type MessageTextGeneric struct {
	Attachment AttachmentGeneric `json:"attachment"`
}

type AttachmentGeneric struct {
	Type    string         `json:"type"`
	Payload PayloadGeneric `json:"payload"`
}

type PayloadGeneric struct {
	TemplateType string           `json:"template_type"`
	Elements     []ElementGeneric `json:"elements"`
}

type ElementGeneric struct {
	Title    string          `json:"title"`
	Subtitle string          `json:"subtitle"`
	ImageUrl string          `json:"image_url"`
	Buttons  []ButtonGeneric `json:"buttons"`
}

type ButtonGeneric struct {
	Type    string `json:"type"`
	Payload string `json:"payload"`
	Title   string `json:"title"`
}

// Send a generic message using the Send API.
func (messenger *Messenger) SendMessageGeneric(recipientId string, messageElt []ElementGeneric) (*http.Response, error) {
	messageData := MessageGeneric{
		Recipient: ObjectWithId{Id: recipientId},
		Message: MessageTextGeneric{
			Attachment: AttachmentGeneric{
				Type: "template",
				Payload: PayloadGeneric{
					TemplateType: "generic",
					Elements:     messageElt,
				},
			},
		},
	}

	return messenger.CallSendAPI(messageData)
}
