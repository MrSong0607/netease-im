package tests

import (
	"testing"

	netease "github.com/MrSong0607/netease-im"
)

func TestSendTextMessage(t *testing.T) {
	msg := &netease.TextMessage{Message: "message test"}
	err := client.SendTextMessage("1", "3", msg, nil)
	if err != nil {
		t.Error(err)
	}
}
