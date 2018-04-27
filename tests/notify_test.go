package tests

import (
	"bytes"
	"net/http"
	"os"
	"strconv"
	"testing"
	"time"

	netease "github.com/MrSong0607/netease-im"
)

func TestCheckSum(t *testing.T) {
	os.Setenv("GOCACHE", "off")
	cli := netease.CreateImClient("b2c60dbed0ae2d3c48e6c85664836dc9", "1ed04f7d7085", "")

	body := []byte(`{}`)
	req, _ := http.NewRequest("POST", "http://yunxinservice.com.cn/receiveMsg.action", bytes.NewReader(body))
	curTime := strconv.FormatInt(time.Now().UnixNano(), 10)
	md5 := netease.Md5HashToHexString(body)
	req.Header.Set("CurTime", curTime)
	req.Header.Set("MD5", md5)
	req.Header.Set("CheckSum", netease.ShaHashToHexStringFromString(cli.AppSecret+md5+curTime))
	t.Log("checksum:", cli.AppSecret+md5+curTime, "checksum-encoded:", netease.ShaHashToHexStringFromString(cli.AppSecret+md5+curTime))

	bd, err := cli.GetEventNotification(req)
	t.Log(string(bd), err)
}
