package tests

import (
	"strconv"
	"testing"
	"time"
)

func TestQueryMessage(t *testing.T) {
	res, err := client.QueryMessage("195371", "179495", "0", strconv.FormatInt(time.Now().UnixNano(), 10), 100, 0, "")
	if err != nil {
		t.Error(err)
	}
	// t.Log(res)
	for _, val := range res {
		t.Log(val)
	}
}
