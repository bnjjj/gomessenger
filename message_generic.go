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
	TemplateType     string           `json:"template_type"`
	ImageAspectRatio string           `json:"image_aspect_ratio,omitempty"`
	Elements         []ElementGeneric `json:"elements"`
}

type ElementGeneric struct {
	Title    string   `json:"title"`
	Subtitle string   `json:"subtitle"`
	ImageUrl string   `json:"image_url"`
	Buttons  []Button `json:"buttons"`
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
