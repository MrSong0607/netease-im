package tests

import (
	"testing"
)

func TestQueryMessage(t *testing.T) {
	res, err := client.QueryMessage("171013", "", "0", "1531128348000", 100, 0, "")
	if err != nil {
		t.Error(err)
	}
	t.Log(res)
}
