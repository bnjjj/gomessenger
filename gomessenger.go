// Gomessenger is a SDK to interact with Facebook Messenger Platform API in Golang
package gomessenger

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"
)

type Messenger struct {
	AccessToken string
}

type ObjectWithId struct {
	Id string `json:"id"`
}

// Create messenger with access token
func New(accessToken string) *Messenger {
	return &Messenger{
		AccessToken: accessToken,
	}
}

// Call the Send API. The message data goes in the body. If successful, we'll
// get the message id in a response
func (messenger *Messenger) CallSendAPI(messageData interface{}) (*http.Response, error) {
	url := "https://graph.facebook.com/v3.3/me/messages?access_token=" + messenger.AccessToken

	client := &http.Client{
		Timeout: time.Second * 10,
	}

	body := new(bytes.Buffer)
	json.NewEncoder(body).Encode(messageData)

	req, _ := http.NewRequest("POST", url, body)

	fmt.Println("REQUEST [POST]: " + url + " params : " + body.String())
	req.Header.Add("Content-Type", "application/json")
	resp, err := client.Do(req)

	if err != nil {
		fmt.Println("Error request : " + err.Error())
		return resp, err
	}

	json.NewEncoder(body).Encode(resp.Body)
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		fmt.Printf("Error http resp -> %v\n", resp)
		fmt.Println("Error http " + strconv.Itoa(resp.StatusCode) + " -> " + body.String())
	}

	return resp, err
}
