package gomessenger

import "net/http"

type MessageQuickReply struct {
	Recipient ObjectWithId             `json:"recipient"`
	Message   MessageContentQuickReply `json:"message"`
}

type MessageContentQuickReply struct {
	Text         string       `json:"text"`
	QuickReplies []QuickReply `json:"quick_replies"`
}

type QuickReply struct {
	ContentType string `json:"content_type"`
	Title       string `json:"title,omitempty"`
	ImageUrl    string `json:"image_url,omitempty"`
	Payload     string `json:"payload,omitempty"`
}

// Send a text message using the Send API.
func (messenger *Messenger) SendQuickReply(recipientId string, text string, quickReplies []QuickReply) (*http.Response, error) {
	messageData := MessageQuickReply{
		Recipient: ObjectWithId{Id: recipientId},
		Message: MessageContentQuickReply{
			Text:         text,
			QuickReplies: quickReplies,
		},
	}

	return messenger.CallSendAPI(messageData)
}
