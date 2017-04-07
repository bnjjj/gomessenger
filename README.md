# Gomessenger [![GoDoc](https://godoc.org/github.com/bnjjj/gomessenger?status.svg)](http://godoc.org/github.com/bnjjj/gomessenger)

## Gomessenger is a SDK to interact with Facebook Messenger API in Golang

## Disclaimer

The Facebook Messenger API isn't completely implemented. It works for my case ([Talk to my car](http://www.talk-to-my-car.com) but I need few features so it's up to you to add your own feature. It's very simple, just add your struct and the call with the function `CallSendApi()` !

### Usage

```golang
messenger := gomessenger.New("myAccessToken")

messenger.SendMessageText("senderID", "my response")

messenger.SendMessageGeneric("senderID", []ElementGeneric{})

messenger.SendAccountLinking("senderID", "https://myurl.com/authorize", "my text to login")
```

## References

It's the package used for [Talk to my car](http://www.talk-to-my-car.com)

## Contributions

Feel free to contribute and extend this package and if you have bugs or if you want more specs make an issue. Have fun !

-------------

Made by [Coenen Benjamin](https://twitter.com/BnJ25) with love