package tests

import (
	"os"
	"testing"
)

func TestGetRoomInfo(t *testing.T) {
	os.Setenv("HTTP_PROXY", "http://127.0.0.1:8889")
	os.Setenv("HTTPS_PROXY", "http://127.0.0.1:8889")
	t.Log(client.GetRoomInfo("6259031888023828"))
}
func TestDeleteRoom(t *testing.T) {
	os.Setenv("HTTP_PROXY", "http://127.0.0.1:8889")
	os.Setenv("HTTPS_PROXY", "http://127.0.0.1:8889")
	t.Log(client.DeleteRoom("6259031888023828"))
}
