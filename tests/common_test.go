package tests

import (
	"testing"
	"time"
)

func TestTimeStamp(t *testing.T) {
	t.Log(time.Now().Unix())
}

func TestArray(t *testing.T) {
	var a []string
	t.Log(a)
	t.Log(len(a))

	a = nil
	t.Log(a)
	t.Log(len(a))
}
