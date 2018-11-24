package tests

import (
	// "os"
	"testing"
)

func TestGetRoomInfo(t *testing.T) {
	// os.Setenv("HTTP_PROXY", "http://127.0.0.1:8889")
	// os.Setenv("HTTPS_PROXY", "http://127.0.0.1:8889")
	t.Log(client.GetRoomInfo("6312415970197875"))
}
func TestDeleteRoom(t *testing.T) {
	// os.Setenv("HTTP_PROXY", "http://127.0.0.1:8889")
	// os.Setenv("HTTPS_PROXY", "http://127.0.0.1:8889")
	t.Log(client.DeleteRoom("6315744469922172"))
}
