package gomessenger

import "net/http"

// Send a message with the account linking call-to-action
func (messenger *Messenger) SendAccountLinking(recipientId string, url string, text string) (*http.Response, error) {
	return messenger.SendButtonsMessage(recipientId, text, []Button{
		Button{
			Type: "account_link",
			URL:  url,
		},
	})
}
