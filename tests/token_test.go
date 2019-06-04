package tests

import (
	"encoding/json"
	"os"
	"testing"

	netease "github.com/MrSong0607/netease-im"
)

var client = netease.CreateImClient("", "", "")

func init() {
	os.Setenv("GOCACHE", "off")
}

func TestToken(t *testing.T) {
	user := &netease.ImUser{ID: "test2", Name: "test3", Gender: 1}
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

func Benchmark_SyncMap(b *testing.B) {
	netease.CreateImClient("", "", "")
}
