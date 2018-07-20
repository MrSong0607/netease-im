package tests

import (
	"encoding/json"
	"testing"

	netease "github.com/MrSong0607/netease-im"
)

var client = netease.CreateImClient("", "", "http://127.0.0.1:8889")

func TestToken(t *testing.T) {
	user := &netease.ImUser{ID: "test1", Name: "test3", Gender: 1}
	tk, err := client.CreateImUser(user)
	if err != nil {
		t.Error(err)
	}
	t.Log(tk)
}

func TestRefreshToken(t *testing.T) {
	tk, err := client.RefreshToken("7")
	if err != nil {
		t.Error(err)
	}
	b, err := json.Marshal(tk)
	t.Log(string(b), err)
}
