# netease-im
本项目是用GO语言实现的网易云信的服务端API封装

### 使用方法
#### 安装:
`go get github.com/MrSong0607/netease-im`

#### 导入:
`import netease "github.com/MrSong0607/netease-im"`

#### 使用:
##### 获取token：
```
client := netease.CreateImClient("AppKey", "AppSecret", "")
user := &netease.ImUser{ID: "3", Name: "test3", Gender: 1}
tk, err := client.CreateImUser(user)
if err != nil {
	t.Error(err)
}
t.Log(tk)
```
##### 发送文本消息
```
msg := &netease.TextMessage{Message: "message test"}
err := client.SendTextMessage("1", "3", msg, nil)
```

## 已实现功能
* [ ] 通信服务
    - [x] 获取IM通信token
    - [x] 发送文本消息