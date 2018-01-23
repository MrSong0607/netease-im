package netease

//APIResult .
type APIResult struct {
	Code string                 `json:"code"`
	Info map[string]interface{} `json:"info"`
}

//TokenInfo 云通信Token
type TokenInfo struct {
	Token string `json:"token"`
	Accid string `json:"accid"`
	Name  string `json:"name"`
}
