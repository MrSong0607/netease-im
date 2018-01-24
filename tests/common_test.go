package tests

import (
	"testing"
	"time"
)

func TestTimeStamp(t *testing.T) {
	t.Log(time.Now().Unix())
}
