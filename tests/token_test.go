package tests

import (
	"testing"

	netease "github.com/MrSong0607/netease-im"
)

var client = netease.CreateImClient("36bb3190572f691d3b180fc099a1b4f1", "99c220190258", "")

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
	t.Log(tk)
}
