package netease

//TokenInfo 云通信Token
type TokenInfo struct {
	Token string `json:"token"`
	Accid string `json:"accid"`
	Name  string `json:"name"`
}

//ImUser 云通信用户
type ImUser struct {
	ID        string //网易云通信ID，最大长度32字符，必须保证一个APP内唯一（只允许字母、数字、半角下划线_、@、半角点以及半角-组成，不区分大小写，会统一小写处理，请注意以此接口返回结果中的accid为准）。
	Name      string //网易云通信ID昵称，最大长度64字符，用来PUSH推送时显示的昵称
	Propertys string
	IconURL   string //网易云通信ID头像URL，第三方可选填，最大长度1024
	Token     string
	Sign      string //用户签名，最大长度256字符
	Email     string //用户email，最大长度64字符
	Birthday  string //用户生日，最大长度16字符
	Mobile    string //用户mobile，最大长度32字符
	Gender    int    //用户性别，0表示未知，1表示男，2女表示女
	Extension string //用户名片扩展字段，最大长度1024字符
}

//ImSendMessageOption .
type ImSendMessageOption struct {
	Antispam         bool            //对于对接了易盾反垃圾功能的应用，本消息是否需要指定经由易盾检测的内容（antispamCustom）。true或false, 默认false。只对消息类型为：100 自定义消息类型 的消息生效。
	AntispamCustom   *AntiSpamCustom //在antispam参数为true时生效。
	Option           *MessageOption  //发消息时特殊指定的行为选项
	Pushcontent      string          //ios推送内容，不超过150字符，option选项中允许推送（push=true），此字段可以指定推送内容
	Payload          string          //ios 推送对应的payload,必须是JSON,不能超过2k字符
	Extension        string          //开发者扩展字段，长度限制1024字符
	ForcePushList    []string        //发送群消息时的强推（@操作）用户列表，格式为JSONArray，如["accid1","accid2"]。若forcepushall为true，则forcepushlist为除发送者外的所有有效群成员
	ForcePushContent string          //发送群消息时，针对强推（@操作）列表forcepushlist中的用户，强制推送的内容
	ForcePushAll     bool            //发送群消息时，强推（@操作）列表是否为群里除发送者外的所有有效成员，true或false，默认为false
	Bid              string          //可选，反垃圾业务ID，实现“单条消息配置对应反垃圾”，若不填则使用原来的反垃圾配置
}

//ImSendAttachMessageOption .
type ImSendAttachMessageOption struct {
	Pushcontent string         //iOS推送内容，第三方自己组装的推送内容,不超过150字符
	Payload     string         //ios 推送对应的payload,必须是JSON,不能超过2k字符
	Sound       string         //如果有指定推送，此属性指定为客户端本地的声音文件名，长度不要超过30个字符，如果不指定，会使用默认声音
	Save        int            //1表示只发在线，2表示会存离线，其他会报414错误。默认会存离线
	Option      *MessageOption //发消息时特殊指定的行为选项
}

//AntiSpamCustom 自定义的反垃圾检测内容, JSON格式，不能超过5000字符
type AntiSpamCustom struct {
	Type int    `json:"type"` //1：文本，2：图片。
	Data string `json:"data"` // 文本内容or图片地址
}

//MessageOption 发消息时特殊指定的行为选项
type MessageOption struct {
	Roam         *bool `json:"roam,omitempty"`         //该消息是否需要漫游，默认true（需要app开通漫游消息功能）
	History      *bool `json:"history,omitempty"`      //该消息是否存云端历史，默认true
	Sendersync   *bool `json:"sendersync,omitempty"`   //该消息是否需要发送方多端同步，默认true
	Push         *bool `json:"push,omitempty"`         //该消息是否需要APNS推送或安卓系统通知栏推送，默认true
	Route        *bool `json:"route,omitempty"`        //该消息是否需要抄送第三方；默认true (需要app开通消息抄送功能)
	Badge        *bool `json:"badge,omitempty"`        //该消息是否需要计入到未读计数中，默认true
	NeedPushNick *bool `json:"needPushNick,omitempty"` //推送文案是否需要带上昵称，不设置该参数时默认true
	Persistent   *bool `json:"persistent,omitempty"`   //是否需要存离线消息，不设置该参数时默认true
}

//TextMessage 文本消息
type TextMessage struct {
	Message string `json:"msg"`
}

//ImageMessage 图片消息
type ImageMessage struct {
	Name      string `json:"name"` //图片name
	Md5       string `json:"md5"`  //图片文件md5
	URL       string `json:"url"`  //生成的url
	Extension string `json:"ext"`  //图片后缀
	Width     uint   `json:"w"`    //宽
	Height    uint   `json:"h"`    //高
	Size      uint   `json:"size"` //图片大小
}
