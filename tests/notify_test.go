package tests

import (
	"bytes"
	"net/http"
	"strconv"
	"testing"
	"time"

	netease "github.com/MrSong0607/netease-im"
)

func TestCheckSum(t *testing.T) {
	body := []byte(`{"attach":"thisisattach","body":"hello","convType":"PERSON","eventType":"1","fromAccount":"111","fromClientType":"IOS","fromDeviceId":"thisisfromdeviceid","fromNick":"mike","msgTimestamp":"1441977355557","msgType":"TEXT","msgidClient":"1234567","msgidServer":"3456789","resendFlag":"0","to":"222"}`)
	req, _ := http.NewRequest("POST", "http://yunxinservice.com.cn/receiveMsg.action", bytes.NewReader(body))
	curTime := strconv.FormatInt(time.Now().UnixNano(), 10)
	md5 := netease.ShaHashToHexString(body)
	req.Header.Set("CurTime", curTime)
	req.Header.Set("MD5", md5)
	req.Header.Set("CheckSum", netease.ShaHashToHexStringFromString(client.AppSecret+md5+curTime))

	err := client.GetEventNotification(req)
	t.Log(err)
}
