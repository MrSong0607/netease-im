# netease-im
netease-im 是用GO语言实现的网易云信的服务端API封装，目前实现了常用的大部分功能，如有其他的需要或者功能失效，可以提issue告知

![](https://img.shields.io/badge/language-golang-blue.svg)
[![996.icu](https://img.shields.io/badge/link-996.icu-red.svg)](https://996.icu)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)
[![LICENSE](https://img.shields.io/badge/license-Anti%20996-blue.svg)](https://github.com/996icu/996.ICU/blob/master/LICENSE)

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
netease-im 使用[MIT](https://opensource.org/licenses/MIT)开源协议