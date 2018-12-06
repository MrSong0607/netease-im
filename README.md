## 使用方法
#### 安装:
`go get -u github.com/MrSong0607/netease-im`

#### 导入:
`import netease "github.com/MrSong0607/netease-im"`

#### 使用:
##### 获取token：
```
client := netease.CreateImClient("AppKey", "AppSecret", "")
user := &netease.ImUser{ID: "3", Name: "test3", Gender: 1}
tk, err := client.CreateImUser(user)
```
##### 发送文本消息
```
msg := &netease.TextMessage{Message: "message test"}
err := client.SendTextMessage("1", "3", msg, nil)
```
##### 发送图片
```
msg := &netease.ImageMessage{URL: "https://golang.org/doc/gopher/frontpage.png", Md5: "可以填任意md5", Extension: "png"}
err := client.SendBatchImageMessage("1", []string{"3"}, msg, nil)
```
##### 发送语音
```
msg := &netease.VoiceMessage{URL: "audio url", Md5: "可以填任意md5", Duration: 10, Extension: "aac"}
err := client.SendBatchVoiceMessage("1", []string{"3"}, msg, nil)
```
##### 发送视频
```
msg := &netease.VideoMessage{URL: "video file url", Md5: "可以填任意md5", Extension: "mp4"}
err := client.SendBatchVideoMessage("1", []string{"3"}, msg, nil)
```
## 已实现功能
* [ ] 通信服务
    - [x] 获取IM通信token
    - [x] 更新并获取新token
    - [x] 发送文本消息
    - [x] 发送图片
    - [x] 发送视频
    - [x] 批量发送文本消息
    - [x] 批量发送点对点自定义系统通知	
    - [x] 查询单聊历史消息
    - [x] 聊天室
    - [x] 消息抄送

##  License
netease-im 使用[GPL](https://www.gnu.org/licenses/gpl-3.0.en.html)开源协议