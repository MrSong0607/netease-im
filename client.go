package netease

import (
	"strconv"
	"sync"
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

	mutex  *sync.Mutex
	client *resty.Client
}

//CreateImClient  创建im客户端，proxy留空表示不使用代理
func CreateImClient(appkey, appSecret, httpProxy string) *ImClient {
	c := &ImClient{AppKey: appkey, AppSecret: appSecret, Nonce: RandStringBytesMaskImprSrc(64), mutex: new(sync.Mutex)}
	c.client = resty.New()
	if len(httpProxy) > 0 {
		c.client.SetProxy(httpProxy)
	}

	c.client.SetHeader("Accept", "application/json;charset=utf-8")
	c.client.SetHeader("Content-Type", "application/x-www-form-urlencoded;charset=utf-8;")
	c.client.SetHeader("AppKey", c.AppKey)
	c.client.SetHeader("Nonce", c.Nonce)

	return c
}

func (c *ImClient) setCommonHead(req *resty.Request) {
	c.mutex.Lock() //多线程并发访问map导致panic
	defer c.mutex.Unlock()

	timeStamp := strconv.FormatInt(time.Now().Unix(), 10)
	req.SetHeader("CurTime", timeStamp)
	req.SetHeader("CheckSum", ShaHashToHexStringFromString(c.AppSecret+c.Nonce+timeStamp))
}
