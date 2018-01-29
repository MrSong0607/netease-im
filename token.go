package netease

import (
	"encoding/json"
	"errors"
	"strconv"

	"github.com/json-iterator/go"
)

const (
	neteaseBaseURL    = "https://api.netease.im/nimserver"
	createImUserPoint = neteaseBaseURL + "/user/create.action"
	refreshTokenPoint = neteaseBaseURL + "/user/refreshToken.action"
)

//CreateImUser 创建网易云通信ID
/**
 * @param accid 网易云通信ID，最大长度32字符，必须保证一个APP内唯一（只允许字母、数字、半角下划线_、@、半角点以及半角-组成，不区分大小写，会统一小写处理，请注意以此接口返回结果中的accid为准）。
 * @param name 网易云通信ID昵称，最大长度64字符，用来PUSH推送时显示的昵称
 * @param props json属性，第三方可选填，最大长度1024字符
 * @param icon 网易云通信ID头像URL，第三方可选填，最大长度1024
 * @param token 网易云通信ID可以指定登录token值，最大长度128字符，并更新，如果未指定，会自动生成token，并在创建成功后返回
 * @param sign 用户签名，最大长度256字符
 * @param email 用户email，最大长度64字符
 * @param birth 用户生日，最大长度16字符
 * @param mobile 用户mobile，最大长度32字符
 * @param gender 用户性别，0表示未知，1表示男，2女表示女，其它会报参数错误
 * @param ex 用户名片扩展字段，最大长度1024字符，用户可自行扩展，建议封装成JSON字符串
 */
func (c *ImClient) CreateImUser(u *ImUser) (*TokenInfo, error) {
	param := map[string]string{"accid": u.ID}

	if len(u.Name) > 0 {
		param["name"] = u.Name
	}
	if len(u.Propertys) > 0 {
		param["props"] = u.Propertys
	}
	if len(u.IconURL) > 0 {
		param["icon"] = u.IconURL
	}
	if len(u.Token) > 0 {
		param["token"] = u.Token
	}
	if len(u.Sign) > 0 {
		param["sign"] = u.Sign
	}
	if len(u.Email) > 0 {
		param["email"] = u.Email
	}
	if len(u.Birthday) > 0 {
		param["birth"] = u.Birthday
	}
	if len(u.Mobile) > 0 {
		param["mobile"] = u.Mobile
	}
	if len(u.Extension) > 0 {
		param["ex"] = u.Extension
	}
	if u.Gender == 1 || u.Gender == 2 {
		param["gender"] = strconv.Itoa(u.Gender)
	}

	client := c.client.R()
	client.SetFormData(param)

	resp, err := client.Post(createImUserPoint)
	if err != nil {
		return nil, err
	}

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
		var msg string
		json.Unmarshal(*jsonRes["desc"], &msg)
		return nil, errors.New(msg)
	}

	tk := &TokenInfo{}
	err = jsoniter.Unmarshal(*jsonRes["info"], tk)
	if err != nil {
		return nil, err
	}

	return tk, nil
}

//RefreshToken 更新并获取新token
/**
 * @param accid 网易云通信ID，最大长度32字符，必须保证一个APP内唯一
 */
func (c *ImClient) RefreshToken(accid string) (*TokenInfo, error) {
	if len(accid) == 0 {
		return nil, errors.New("必须指定网易云通信ID")
	}

	param := map[string]string{"accid": accid}
	client := c.client.R()
	client.SetFormData(param)

	resp, err := client.Post(refreshTokenPoint)
	if err != nil {
		return nil, err
	}

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
		var msg string
		json.Unmarshal(*jsonRes["desc"], &msg)
		return nil, errors.New(msg)
	}

	tk := &TokenInfo{}
	err = jsoniter.Unmarshal(*jsonRes["info"], tk)
	if err != nil {
		return nil, err
	}

	return tk, nil
}
