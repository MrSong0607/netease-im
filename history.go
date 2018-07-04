package netease

import (
	"encoding/json"
	"errors"
	"strconv"

	jsoniter "github.com/json-iterator/go"
)

const (
	queryMessagePoint = neteaseBaseURL + "/history/querySessionMsg.action"
)

//HistoryMessage .
type HistoryMessage struct {
	From string      `json:"from"`
	ID   int64       `json:"msgid"`
	Time int64       `json:"sendtime"`
	Type int         `json:"type"`
	Body interface{} `json:"body"`
}

//QueryMessage 查询存储在网易云通信服务器中的单人聊天历史消息，只能查询在保存时间范围内的消息
/*
*1.跟据时间段查询点对点消息，每次最多返回100条；
*2.不提供分页支持，第三方需要跟据时间段来查询。
* @param fromID	发送者accid
* @param toID	接收者accid
* @param beginTime	开始时间，ms
* @param endTime	截止时间，ms
* @param limit		本次查询的消息条数上限(最多100条),小于等于0，或者大于100，会提示参数错误
* @param reverse	1按时间正序排列，2按时间降序排列。其它返回参数414错误.默认是按降序排列
* @param tp	查询指定的多个消息类型，类型之间用","分割，不设置该参数则查询全部类型消息格式 示例： 0,1,2,3 类型支持： 1:图片，2:语音，3:视频，4:地理位置，5:通知，6:文件，10:提示，11:Robot，100:自定义
 */
func (c *ImClient) QueryMessage(fromID, toID, beginTime, endTime string, limit, reverse int, tp string) ([]HistoryMessage, error) {
	param := map[string]string{"from": fromID}

	param["to"] = toID
	param["begintime"] = beginTime
	param["endtime"] = endTime
	param["limit"] = strconv.Itoa(limit)

	if reverse != 0 {
		param["reverse"] = strconv.Itoa(reverse)
	}

	if len(tp) > 0 {
		param["type"] = tp
	}

	client := c.client.R()
	c.setCommonHead(client)
	client.SetFormData(param)

	resp, err := client.Post(queryMessagePoint)

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

	var msgs []HistoryMessage
	err = jsoniter.Unmarshal(*jsonRes["msgs"], &msgs)
	if err != nil {
		return nil, err
	}

	return msgs, nil
}
