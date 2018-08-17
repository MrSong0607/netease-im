package netease

import (
	"encoding/json"
	"errors"
	"strconv"

	jsoniter "github.com/json-iterator/go"
)

const (
	broadcastMessagePoint = neteaseBaseURL + "/msg/broadcastMsg.action"
)

//BroadcastResult 广播推送结果
type BroadcastResult struct {
	BroadcastID int64    `json:"broadcastId"`
	ExpireTime  int64    `json:"expireTime"`
	Body        string   `json:"body"`
	CreateTime  int64    `json:"createTime"`
	IsOffline   bool     `json:"isOffline"`
	TargetOs    []string `json:"targetOs"`
}

//BroadcastMsg 发送广播消息
/*
 * *广播消息，可以对应用内的所有用户发送广播消息，广播消息目前暂不支持第三方推送（APNS、小米、华为等）
 * *广播消息支持离线存储，并可以自定义设置离线存储的有效期，最多保留最近100条离线广播消息
 * *此接口受频率控制，一个应用一分钟最多调用10次，一天最多调用1000次，超过会返回416状态码
 * @param body 广播消息内容，最大4096字符
 * @param fromID(可选) 发送者accid, 用户帐号，最大长度32字符，必须保证一个APP内唯一
 * @param isOffline(可选) 是否存离线，true或false，默认false
 * @param targetOs(可选) 目标客户端，默认所有客户端,"ios","aos","pc","web","mac"
 */
func (c *ImClient) BroadcastMsg(body, fromID string, isOffline *bool, targetOs []string) (msgs *BroadcastResult, err error) {
	param := map[string]string{"body": body}

	if len(fromID) > 0 {
		param["from"] = fromID
	}
	if isOffline != nil {
		param["isOffline"] = strconv.FormatBool(*isOffline)
	}
	if len(targetOs) > 0 {
		if param["targetOs"], err = jsonTool.MarshalToString(targetOs); err != nil {
			return nil, err
		}
	}

	client := c.client.R()
	c.setCommonHead(client)
	client.SetFormData(param)

	resp, err := client.Post(broadcastMessagePoint)

	var jsonRes map[string]*json.RawMessage
	err = jsoniter.Unmarshal(resp.Body(), &jsonRes)
	if err != nil {
		return nil, err
	}

	var code int
	err = json.Unmarshal(*jsonRes["code"], &code)
	if err != nil {
		return nil, err
	}

	if code != 200 {
		return nil, errors.New(string(resp.Body()))
	}

	msgs = &BroadcastResult{}
	err = jsoniter.Unmarshal(*jsonRes["msg"], msgs)
	if err != nil {
		return nil, err
	}

	return msgs, nil
}
