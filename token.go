package netease

import (
	"strconv"

	"github.com/json-iterator/go"
)

const (
	neteaseBaseURL    = "https://api.netease.im/nimserver"
	createImUserPoint = neteaseBaseURL + "/user/create.action"
)

//CreateImUser .
/**
 * 创建网易云通信ID
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
func (c *ImClient) CreateImUser(accid, name, props, icon, token, sign, email, birth, mobile, ex string, gender int) (*TokenInfo, error) {
	param := map[string]string{"accid": accid}

	if len(name) > 0 {
		param["name"] = name
	}
	if len(props) > 0 {
		param["props"] = props
	}
	if len(icon) > 0 {
		param["icon"] = icon
	}
	if len(token) > 0 {
		param["token"] = token
	}
	if len(sign) > 0 {
		param["sign"] = sign
	}
	if len(email) > 0 {
		param["email"] = email
	}
	if len(birth) > 0 {
		param["birth"] = birth
	}
	if len(mobile) > 0 {
		param["mobile"] = mobile
	}
	if len(ex) > 0 {
		param["ex"] = ex
	}
	if gender == 1 || gender == 2 {
		param["gender"] = strconv.Itoa(gender)
	}

	client := c.genRestClient()
	client.SetFormData(param)

	resp, err := client.Post(createImUserPoint)
	if err != nil {
		return nil, err
	}

	r := &APIResult{}
	err = jsoniter.Unmarshal(resp.Body(), r)
	if err != nil {
		return nil, err
	}

	tk := &TokenInfo{}
	if val, ok := r.Info["token"]; ok {
		tk.Token = val.(string)
	}

	if val, ok := r.Info["accid"]; ok {
		tk.Accid = val.(string)
	}

	if val, ok := r.Info["name"]; ok {
		tk.Name = val.(string)
	}

	return tk, nil
}
