# netease-im
本项目是用GO语言实现的网易云信的服务端API封装

### 使用方法
#### 安装:
`go get github.com/MrSong0607/netease-im`

#### 导入:
`netease "github.com/MrSong0607/netease-im"`

#### 使用:
```
    client := netease.CreateImClient("AppKey", "AppSecret", "")
	tk, err := client.CreateImUser("2", "test2", "", "", "", "", "", "", "", "", 1)
	if err != nil {
		t.Error(err)
	}
	t.Log(tk)
```

## 已实现功能
* [ ] 通信服务
    - [x] 获取IM通信token