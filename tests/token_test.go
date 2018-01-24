package tests

import (
	"testing"

	netease "github.com/MrSong0607/netease-im"
)

var client = netease.CreateImClient("36bb3190572f691d3b180fc099a1b4f1", "99c220190258", "http://127.0.0.1:8888")

func TestToken(t *testing.T) {
	tk, err := client.CreateImUser("2", "test2", "", "", "", "", "", "", "", "", 1)
	if err != nil {
		t.Error(err)
	}
	t.Log(tk)
}
