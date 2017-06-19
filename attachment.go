package gomessenger

import "net/http"

type AttachmentMessage struct {
	Recipient ObjectWithId      `json:"recipient"`
	Message   MessageAttachment `json:"message"`
}

type MessageAttachment struct {
	Attachment AttachmentUpload `json:"attachment"`
}

type AttachmentUpload struct {
	Type    string     `json:"type"`
	Payload PayloadUrl `json:"payload"`
}

type PayloadUrl struct {
	URL        string `json:"url"`
	IsReusable bool   `json:"is_reusable,omitempty"`
}

// Send a message with attachment using the Send API.
func (messenger *Messenger) SendAttachment(recipientId, attachmentType, url string) (*http.Response, error) {
	messageData := AttachmentMessage{
		Recipient: ObjectWithId{Id: recipientId},
		Message: MessageAttachment{
			Attachment: AttachmentUpload{
				Type: attachmentType,
				Payload: PayloadUrl{
					URL:        url,
					IsReusable: true,
				},
			},
		},
	}

	return messenger.CallSendAPI(messageData)
}
