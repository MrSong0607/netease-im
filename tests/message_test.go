package tests

import (
	"os"
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

func TestBroadcastMsg(t *testing.T) {
	os.Setenv("GOCACHE", "off")
	t.Log(client.BroadcastMsg("好久不见了呢，我在这里等你哦", "", nil, nil))
}
