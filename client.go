package netease

import (
	"strconv"
	"time"

	"github.com/go-resty/resty"
	jsoniter "github.com/json-iterator/go"
)

var jsonTool = jsoniter.ConfigCompatibleWithStandardLibrary

//ImClient .
type ImClient struct {
	AppKey    string
	AppSecret string
	Nonce     string
}

//CreateImClient .
func CreateImClient(appkey, appSecret string) *ImClient {
	return &ImClient{AppKey: appkey, AppSecret: appSecret, Nonce: RandStringBytesMaskImprSrc(64)}
}

func (c *ImClient) genRestClient() *resty.Request {
	client := resty.R()
	client.SetHeader("Content-Type", "application/x-www-form-urlencoded;charset=utf-8;")
	client.SetHeader("AppKey", c.AppKey)
	client.SetHeader("Nonce", c.Nonce)

	timeStamp := strconv.FormatInt(time.Now().UnixNano(), 10)
	client.SetHeader("CurTime", timeStamp)
	client.SetHeader("CheckSum", ShaHashToHexStringFromString(c.AppSecret+c.Nonce+timeStamp))

	return client
}
