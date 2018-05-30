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

func TestSendBatchTextMessage(t *testing.T) {
	msg := &netease.TextMessage{Message: "message test"}
	str, err := client.SendBatchTextMessage("1", []string{"2", "3"}, msg, nil)
	t.Log(str)
	if err != nil {
		t.Error(err)
	}
}

func TestSendBatchAttachMessage(t *testing.T) {
	err := client.SendBatchAttachMsg("1", "{'msg':'test'}", []string{"2", "3"}, nil)
	if err != nil {
		t.Error(err)
	}
}
