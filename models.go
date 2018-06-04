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

//VoiceMessage 语音消息
type VoiceMessage struct {
	Duration  uint   `json:"dur"`  //语音持续时长ms
	Md5       string `json:"md5"`  //语音文件md5
	URL       string `json:"url"`  //生成的url
	Extension string `json:"ext"`  //语音消息格式，只能是aac格式
	Size      uint   `json:"size"` //语音文件大小
}

//VideoMessage 视频消息
type VideoMessage struct {
	Duration  uint   `json:"dur"`  //视频持续时长ms
	Md5       string `json:"md5"`  //视频文件md5
	URL       string `json:"url"`  //生成的url
	Width     uint   `json:"w"`    //宽
	Height    uint   `json:"h"`    //高
	Extension string `json:"ext"`  //视频格式
	Size      uint   `json:"size"` //视频文件大小
}

//LocationMessage 位置信息
type LocationMessage struct {
	Title     string  `json:"title"` //地理位置title
	Longitude float64 `json:"lng"`   //经度
	Latitude  float64 `json:"lat"`   //纬度
}

//FileMessage 文件消息
type FileMessage struct {
	Name      string `json:"name"` //文件名
	Md5       string `json:"md5"`  //图片文件md5
	URL       string `json:"url"`  //生成的url
	Extension string `json:"ext"`  //语音消息格式，只能是aac格式
	Size      uint   `json:"size"` //语音文件大小
}

//LoginEventCopyInfo 登录事件消息抄送
type LoginEventCopyInfo struct {
	EventType  string `json:"eventType"`  //值为2，表示是登录事件的消息
	AcctID     string `json:"accid"`      //发生登录事件的用户帐号，字符串类型
	IPAdrees   string `json:"clientIp"`   //登录时的ip地址
	ClientType string `json:"clientType"` //客户端类型： AOS、IOS、PC、WINPHONE、WEB、REST，字符串类型
	Code       string `json:"code"`       //登录事件的返回码，可转为Integer类型数据
	SdkVersion string `json:"sdkVersion"` //当前sdk的版本信息，字符串类型
	Time       string `json:"timestamp"`  //登录事件发生时的时间戳，可转为Long型数据
}

//SenssionCopyInfo 会话类型信息抄送
type SenssionCopyInfo struct {
	EventType      string `json:"eventType"`      //值为1，表示是会话类型的消息
	ConvType       string `json:"convType"`       //会话具体类型：PERSON（二人会话数据）、TEAM（群聊数据）、	CUSTOM_PERSON（个人自定义系统通知）、CUSTOM_TEAM（群组自定义系统通知），字符串类型
	To             string `json:"to"`             //若convType为PERSON或CUSTOM_PERSON，则to为消息接收者的用户账号，字符串类型；若convType为TEAM或CUSTOM_TEAM，则to为tid，即群id，可转为Long型数据
	FromAccount    string `json:"fromAccount"`    //消息发送者的用户账号，字符串类型
	FromClientType string `json:"fromClientType"` //发送客户端类型： AOS、IOS、PC、WINPHONE、WEB、REST，字符串类型
	FromDeviceID   string `json:"fromDeviceId"`   //发送设备id，字符串类型
	FromNick       string `json:"fromNick"`       //发送方昵称，字符串类型
	MsgTimestamp   string `json:"msgTimestamp"`   //消息发送时间，字符串类型
	MsgType        string `json:"msgType"`        //会话具体类型PERSON、TEAM对应的通知消息类型:EXT、PICTURE、AUDIO、VIDEO、LOCATION 、NOTIFICATION、FILE、 //文件消息NETCALL_AUDIO、 //网络电话音频聊天 	NETCALL_VEDIO、 //网络电话视频聊天 	DATATUNNEL_NEW、 //新的数据通道请求通知 	TIPS、 //提示类型消息 	CUSTOM //自定义消息		会话具体类型CUSTOM_PERSON对应的通知消息类型：	FRIEND_ADD、 //加好友	FRIEND_DELETE、 //删除好友	CUSTOM_P2P_MSG、 //个人自定义系统通知		会话具体类型CUSTOM_TEAM对应的通知消息类型：	TEAM_APPLY、 //申请入群	TEAM_APPLY_REJECT、 //拒绝入群申请	TEAM_INVITE、 //邀请进群	TEAM_INVITE_REJECT、 //拒绝邀请	TLIST_UPDATE、 //群信息更新 	CUSTOM_TEAM_MSG、 //群组自定义系统通知
	Body           string `json:"body"`           //消息内容，字符串类型
	Attach         string `json:"attach"`         //附加消息，字符串类型
	MsgidClient    string `json:"msgidClient"`    //客户端生成的消息id，仅在convType为PERSON或TEAM含此字段，字符串类型
	MsgidServer    string `json:"msgidServer"`    //服务端生成的消息id，可转为Long型数据
	ResendFlag     string `json:"resendFlag"`     //重发标记：0不是重发, 1是重发。仅在convType为PERSON或TEAM时含此字段，可转为Integer类型数据
	CustomSafeFlag string `json:"customSafeFlag"` //自定义系统通知消息是否存离线:0：不存，1：存。仅在convType为CUSTOM_PERSON或CUSTOM_TEAM时含此字段，可转为Integer类型数据
	CustomApnsText string `json:"customApnsText"` //自定义系统通知消息推送文本。仅在convType为CUSTOM_PERSON或CUSTOM_TEAM时含此字段，字符串类型
	TMembers       string `json:"tMembers"`       //跟本次群操作有关的用户accid，仅在convType为TEAM或CUSTOM_TEAM时含此字段，字符串类型
	Ext            string `json:"ext"`            //消息扩展字段
	Antispam       string `json:"antispam"`       //标识是否被反垃圾，仅在被反垃圾时才有此字段，可转为Boolean类型数据
	YidunRes       string `json:"yidunRes"`       //易盾反垃圾的原始处理细节，只有接入了相关功能易盾反垃圾的应用才会有这个字段。
}

//AudioCopyInfo 音视频/白板时长消息抄送
type AudioCopyInfo struct {
	ChannelID  string `json:"channelId"`  //通道号
	Createtime string `json:"createtime"` //音视频通话/白板开始的事件, 可转为13位时间戳
	Duration   string `json:"duration"`   //此通通话/白板的通话时长，精确到秒，可转为Integer类型
	EventType  string `json:"eventType"`  //为5，表示是实时音视频/白板时长类型事件
	Live       string `json:"live"`       //是否是互动直播的音视频，0：否，1：是
	Members    string `json:"members"`    //表示通话/白板的参与者：accid为用户帐号；如果是通话的发起者的话，caller字段为true，否则无caller字段；duration表示对应accid用户的单方时长，其中白板消息暂无此单方时长的统计
	Status     string `json:"status"`     //通话/白板状态:SUCCESS：表示正常挂断;TIMEOUT：表示超时;SINGLE_PARTICIPATE：表示只有一个参与者;UNKNOWN：表示未知状态
	Type       string `json:"type"`       //类型：AUDIO：表示音频通话；VEDIO：表示视频通话；DataTunnel：表示白板事件
	Ext        string `json:"ext"`        //音视频发起时的自定义字段，可选，由用户指定
	Running    bool   `json:"running"`    //若为true表示超长时长通话的过程中的抄送，缺省或者false表示普通时长通话的抄送或者超长时长通话的最后一次抄送
}

//AudioDownloadCopyInfo 音视频/白板文件下载信息抄送
type AudioDownloadCopyInfo struct {
	EventType string `json:"eventType"` //值为6，表示是音视频/白板文件下载信息类型的消息
	// 可转为JSONArray，其中的字段释义如下：
	// caller：是否是此通通话的发起者，若是则为true，若不是则没有此字段，可转为Boolean值
	// channelid：通道号，可转为Long值
	// filename：文件名，直接存储，混合录制文件filename带有"-mix"标记
	// md5：文件的md5值
	// size：文件大小，单位为字符，可转为Long值
	// type：文件的类型（扩展名），包括：实时音频录制文件(aac)、白板录制文件(gz)、实时视频录制文件(mp4)、互动直播视频录制文件(flv)
	// url：文件的下载地址，请不要解析该字段
	// user：用户帐号，若该文件为混合录制文件，则该字段为"0"
	// mix：是否为混合录制文件，true：混合录制文件；false：单人录制文件
	// vid：点播文件id，注意白板录制文件(gz)无此字段
	FileInfo string `json:"fileinfo"`
}

//FileDownloadInfo 单个文件下载信息
type FileDownloadInfo struct {
	Caller    bool   `json:"caller"`    //是否是此通通话的发起者，若是则为true，若不是则没有此字段，可转为Boolean值
	ChannelID string `json:"channelid"` //通道号
	Filename  string `json:"filename"`  //文件名，直接存储，混合录制文件filename带有"-mix"标记
	Md5       string `json:"md5"`       //文件的md5值
	Mix       bool   `json:"mix"`       //是否为混合录制文件，true：混合录制文件；false：单人录制文件
	Size      string `json:"size"`      //size：文件大小，单位为字符，可转为Long值
	Type      string `json:"type"`      //文件的类型（扩展名），包括：实时音频录制文件(aac)、白板录制文件(gz)、实时视频录制文件(mp4)、互动直播视频录制文件(flv)
	Vid       string `json:"vid"`       //点播文件id，注意白板录制文件(gz)无此字段
	URL       string `json:"url"`       //文件的下载地址，请不要解析该字段
	User      string `json:"user"`      //用户帐号，若该文件为混合录制文件，则该字段为"0"
}
