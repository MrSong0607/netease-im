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
	//EventTypeMediaDuration 汇报实时音视频通话时长、白板事件时长的消息
	EventTypeMediaDuration = "5"
	//EventTypeMediaInfo 汇报音视频/白板文件的大小、下载地址等消息
	EventTypeMediaInfo = "6"
	//EventTypeP2PMessageRecall 单聊消息撤回抄送
	EventTypeP2PMessageRecall = "7"
	//EventTypeGroupMessageRecall 群聊消息撤回抄送
	EventTypeGroupMessageRecall = "8"
	//EventTypeChatRoomInOut 汇报主播或管理员进出聊天室事件消息
	EventTypeChatRoomInOut = "9"
	//EventTypeECPCallback 汇报专线电话通话结束回调抄送的消息
	EventTypeECPCallback = "10"
	//EventTypeSMSCallback 汇报短信回执抄送的消息
	EventTypeSMSCallback = "11"
	//EventTypeSMSReply 汇报短信上行消息
	EventTypeSMSReply = "12"
	//EventTypeAvRoomInOut 汇报用户进出音视频/白板房间的消息
	EventTypeAvRoomInOut = "13"
	//EventTypeChatRoomQueueOperate 汇报聊天室队列操作的事件消息
	EventTypeChatRoomQueueOperate = "14"
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
