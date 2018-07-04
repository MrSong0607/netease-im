package tests

import (
	"os"
	"testing"
)

func TestQueryMessage(t *testing.T) {
	os.Setenv("HTTP_PROXY", "http://127.0.0.1:8889")

	res, err := client.QueryMessage("170323", "188090", "0", "1530687679000", 100, 0, "")
	if err != nil {
		t.Error(err)
	}
	t.Log(res)
}
