package netease

import (
	"errors"
	"io/ioutil"
	"net/http"
)

const (
	//EventTypeConversation 表示CONVERSATION消息，即会话类型的消息（目前包括P2P聊天消息，群组聊天消息，群组操作，好友操作）
	EventTypeConversation = "1"
	//EventTypeLogin 表示LOGIN消息，即用户登录事件的消息
	EventTypeLogin = "2"
	//EventTypeLogout  表示LOGOUT消息，即用户登出事件的消息
	EventTypeLogout = "3"
	//EventTypeChatRoom 表示CHATROOM消息，即聊天室中聊天的消息
	EventTypeChatRoom = "4"
)

//GetEventNotification .
func (c *ImClient) GetEventNotification(req *http.Request) error {
	if req == nil {
		return errors.New("request 参数不能为空")
	}

	checkSum := req.Header.Get("CheckSum")
	md5 := req.Header.Get("MD5")
	curTime := req.Header.Get("CurTime")
	bd, err := ioutil.ReadAll(req.Body)
	if err != nil {
		return err
	}

	if ShaHashToHexString(bd) != md5 {
		return errors.New("消息抄送内容被劫持")
	}

	if checkSum != ShaHashToHexStringFromString(c.AppSecret+md5+curTime) {
		return errors.New("CheckSum校验失败")
	}
	return nil
}
