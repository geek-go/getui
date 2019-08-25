package getui

//@see http://docs.getui.com/getui/server/rest/explain/
type Message struct {
	AppKey            string `json:"appkey"`                        //	注册应用时生成的appkey
	IsOffline         bool   `json:"is_offline,omitempty"`          //	是否离线推送 可选 默认true
	OfflineExpireTime int    `json:"offline_expire_time,omitempty"` //	消息离线存储有效期，单位：ms 默认24小时
	PushNetWorkType   int    `json:"push_network_type,omitempty"`   //	选择推送消息使用网络类型，0：不限制，1：wifi 默认0
	MsgType           string `json:"msgtype"`                       //	消息应用类型，可选项：notification、link、notypopload、transmission。选了MsgType，那么必须有对应的模板，否则失败
}

type messageType struct {
	Notification string
	Link         string
	Notypopload  string
	Transmission string
}

//消息类型
var MsgType = &messageType{
	"notification",
	"link",
	"notypopload",
	"transmission",
}

//获取Message实例
func GetMessage() *Message {
	message := &Message{
		IsOffline:         true,
		OfflineExpireTime: 24 * 60 * 60 * 1000,
		PushNetWorkType:   0,
	}
	return message
}

//json串，当手机为ios，并且为离线的时候
type PushInfo map[string]interface{}

//push_info.Alert
type Alert struct {
	Title string `json:"title"` //通知标题
	Body  string `json:"body"`  //通知内容
}

//push_info.Apns
type Apns struct {
	Alert            *Alert `json:"alert"`               //
	AutoBadge        string `json:"autoBadge,omitempty"` //用于计算icon上显示的数字，还可以实现显示数字的自动增减，如“+1”、 “-1”、 “1” 等，计算结果将覆盖badge
	Sound            string `json:"sound,omitempty"`     //通知铃声文件名，无声设置为“com.gexin.ios.silence”
	ContentAvailable int    `json:"content-available"`   //推送直接带有透传数据
	Category         string `json:"category,omitempty"`  //在客户端通知栏触发特定的action和button显示
}

//多厂商推送透传消息带通知
type Notify struct {
	Title   string `json:"title"`
	Content string `json:"content"`
	Intent  string `json:"intent"`
	Type    string `json:"type"`
}
