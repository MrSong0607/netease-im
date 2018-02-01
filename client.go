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

	client *resty.Client
}

//CreateImClient  创建im客户端，proxy留空表示不使用代理
func CreateImClient(appkey, appSecret, httpProxy string) *ImClient {
	c := &ImClient{AppKey: appkey, AppSecret: appSecret, Nonce: RandStringBytesMaskImprSrc(64)}
	c.client = resty.New()
	if len(httpProxy) > 0 {
		c.client.SetProxy(httpProxy)
	}

	c.client.SetHeader("Accept", "application/json;charset=utf-8")
	c.client.SetHeader("Content-Type", "application/x-www-form-urlencoded;charset=utf-8;")
	c.client.SetHeader("AppKey", c.AppKey)
	c.client.SetHeader("Nonce", c.Nonce)

	timeStamp := strconv.FormatInt(time.Now().Unix(), 10)
	c.client.SetHeader("CurTime", timeStamp)
	c.client.SetHeader("CheckSum", ShaHashToHexStringFromString(c.AppSecret+c.Nonce+timeStamp))

	return c
}
