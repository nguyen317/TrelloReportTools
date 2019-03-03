package controllers

import (
	"github.com/adlio/trello"
)

const (
	keyApp = "fa6b1a601cfd6559cc134d0055507cc2"
	token  = "e0a44f959cde9dfb0883e2865d5632232b0b3ac93900263d22d6e7f84a1d0017"
)

var client *trello.Client
var user *trello.Member

func init() {
	var err error

	client = trello.NewClient(keyApp, token)
	if err != nil {
		// Handle error
	}
}
