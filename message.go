package netease

import (
	"encoding/json"
	"errors"
	"strconv"

	jsoniter "github.com/json-iterator/go"
)

const (
	sendMsgPoint      = neteaseBaseURL + "/msg/sendMsg.action"
	sendBatchMsgPoint = neteaseBaseURL + "/msg/sendBatchMsg.action"
)

//SendTextMessage 发送文本消息,消息内容最长5000
func (c *ImClient) SendTextMessage(fromID, toID string, msg *TextMessage, opt *ImSendMessageOption) error {
	bd, err := jsonTool.MarshalToString(msg)
	if err != nil {
		return err
	}
	return c.SendMessage(fromID, toID, bd, 0, 0, opt)
}

//SendBatchTextMessage 批量发送文本消息
func (c *ImClient) SendBatchTextMessage(fromID string, toIDs []string, msg *TextMessage, opt *ImSendMessageOption) error {
	bd, err := jsonTool.MarshalToString(msg)
	if err != nil {
		return err
	}

	return c.SendBatchMessage(fromID, bd, toIDs, 0, opt)
}

//SendMessage 发送普通消息
/**
 * @param fromID 发送者accid，用户帐号，最大32字符，必须保证一个APP内唯一
 * @param toID ope==0是表示accid即用户id，ope==1表示tid即群id
 * @param ope 0：点对点个人消息，1：群消息（高级群），其他返回414
 * @param msgType 0 表示文本消息,1 表示图片，2 表示语音，3 表示视频，4 表示地理位置信息，6 表示文件，100 自定义消息类型（特别注意，对于未对接易盾反垃圾功能的应用，该类型的消息不会提交反垃圾系统检测）
 * @param body 最大长度5000字符，为一个JSON串
 */
func (c *ImClient) SendMessage(fromID, toID, body string, ope, msgType int, opt *ImSendMessageOption) error {
	param := map[string]string{"from": fromID}

	param["ope"] = strconv.Itoa(ope)
	param["to"] = toID
	param["type"] = strconv.Itoa(msgType)
	param["body"] = body

	if opt != nil {
		param["antispam"] = strconv.FormatBool(opt.Antispam)

		if opt.AntispamCustom != nil {
			param["antispamCustom"], _ = jsonTool.MarshalToString(opt.AntispamCustom)
		}

		if opt.Option != nil {
			param["option"], _ = jsonTool.MarshalToString(opt.Option)
		}

		if len(opt.Pushcontent) > 0 {
			param["pushcontent"] = opt.Pushcontent
		}

		if len(opt.Payload) > 0 {
			param["payload"] = opt.Payload
		}

		if len(opt.Extension) > 0 {
			param["ext"] = opt.Extension
		}

		if opt.ForcePushList != nil {
			param["forcepushlist"], _ = jsonTool.MarshalToString(opt.ForcePushList)
		}

		if len(opt.ForcePushContent) > 0 {
			param["forcepushcontent"] = opt.ForcePushContent
		}
		param["forcepushall"] = strconv.FormatBool(opt.ForcePushAll)
		if len(opt.Bid) > 0 {
			param["bid"] = opt.Bid
		}
	}
	client := c.client.R()
	client.SetFormData(param)

	resp, err := client.Post(sendMsgPoint)

	var jsonRes map[string]*json.RawMessage
	err = jsoniter.Unmarshal(resp.Body(), &jsonRes)
	if err != nil {
		return err
	}

	var code int
	err = json.Unmarshal(*jsonRes["code"], &code)
	if err != nil {
		return err
	}

	if code != 200 {
		return errors.New(string(resp.Body()))
	}

	return nil
}

//SendBatchMessage 批量发送点对点普通消息
/**
 * @param fromID 发送者accid，用户帐号，最大32字符，必须保证一个APP内唯一
 * @param toIDs ["aaa","bbb"]（JSONArray对应的accid，如果解析出错，会报414错误），限500人
 * @param msgType 0 表示文本消息,1 表示图片，2 表示语音，3 表示视频，4 表示地理位置信息，6 表示文件，100 自定义消息类型
 */
func (c *ImClient) SendBatchMessage(fromID, body string, toIDs []string, msgType int, opt *ImSendMessageOption) error {
	param := map[string]string{"fromAccid": fromID}

	to, err := jsonTool.MarshalToString(toIDs)
	if err != nil {
		return err
	}
	param["toAccids"] = to
	param["type"] = strconv.Itoa(msgType)
	param["body"] = body

	if opt != nil {
		if opt.Option != nil {
			param["option"], _ = jsonTool.MarshalToString(opt.Option)
		}

		if len(opt.ForcePushContent) > 0 {
			param["forcepushcontent"] = opt.ForcePushContent
		}

		if len(opt.Payload) > 0 {
			param["payload"] = opt.Payload
		}

		if len(opt.Extension) > 0 {
			param["ext"] = opt.Extension
		}

		if len(opt.Bid) > 0 {
			param["bid"] = opt.Bid
		}
	}
	client := c.client.R()
	client.SetFormData(param)

	resp, err := client.Post(sendBatchMsgPoint)

	var jsonRes map[string]*json.RawMessage
	err = jsoniter.Unmarshal(resp.Body(), &jsonRes)
	if err != nil {
		return err
	}

	var code int
	err = json.Unmarshal(*jsonRes["code"], &code)
	if err != nil {
		return err
	}

	if code != 200 {
		return errors.New(string(resp.Body()))
	}

	return nil
}
