package getui

//点开通知打开应用模板
//在通知栏显示一条含图标、标题等的通知，用户点击后，会激活您的应用
//http://docs.getui.com/getui/server/rest/template/?id=doc-title-0
type Notification struct {
	TransmissionType    bool        `json:"transmission_type,omitempty"`    //收到消息是否立即启动应用，true为立即启动，false则广播等待启动，默认是否  可选
	TransmissionContent string      `json:"transmission_content,omitempty"` //透传内容  可选
	DurationBegin       string      `json:"duration_begin,omitempty"`       //设定展示开始时间，格式为yyyy-MM-dd HH:mm:ss  可选
	DurationEnd         string      `json:"duration_end,omitempty"`         //设定展示结束时间，格式为yyyy-MM-dd HH:mm:ss  可选
	Style               interface{} `json:"style"`                          //通知栏消息布局样式(0 系统样式 1 个推样式) 默认为0  可选
}

//开通知打开网页模板
type Link struct {
	Url           string      `json:"url,omitempty"`            //	打开网址 可选
	DurationBegin string      `json:"duration_begin,omitempty"` //	设定展示开始时间，格式为yyyy-MM-dd HH:mm:ss  可选
	DurationEnd   string      `json:"duration_end,omitempty"`   //设定展示结束时间，格式为yyyy-MM-dd HH:mm:ss  可选
	Style         interface{} `json:"style,omitempty"`          //	通知栏消息布局样式(0 系统样式 1 个推样式) 默认为0  可选
}

//点击通知弹窗下载模板
type NotyPopload struct {
	NotyIcon    string `json:"notyicon"`    //通知栏图标
	NotyTitle   string `json:"notytitle"`   //	通知标题
	NotyContent string `json:"notycontent"` //	通知内容
	//	LogoUrl       string `json:"logourl,omitempty"`        //	通知的网络图标地址 可选
	PopTitle      string      `json:"poptitle"`                 //	弹出框标题
	PopContent    string      `json:"popcontent"`               //弹出框内容
	PopImage      string      `json:"popimage"`                 //弹出框图标
	PopButton1    string      `json:"PopButton1"`               //弹出框左边按钮名称
	PopButton2    string      `json:"PopButton2"`               //弹出框左边按钮名称
	LoadUrl       string      `json:"loadurl"`                  //	下载文件地址
	LoadIcon      string      `json:"loadicon,omitempty"`       // 下载图标  可选
	LoadTitle     string      `json:"loadtitle,omitempty"`      //	下载标题  可选
	IsAutoinstall bool        `json:"is_autoinstall,omitempty"` //	是否自动安装，默认值false 可选
	IsActived     bool        `json:"is_actived,omitempty"`     //安装完成后是否自动启动应用程序，默认值false 可选
	AndroidMark   string      `json:"androidmark,omitempty"`    //安卓标识 可选
	SymbianMark   string      `json:"symbianmark,omitempty"`    //塞班标识 可选
	IphoneMark    string      `json:"iphonemark,omitempty"`     //苹果标志 可选
	DurationBegin string      `json:"duration_begin,omitempty"` //	设定展示开始时间，格式为yyyy-MM-dd HH:mm:ss 可选
	DurationEnd   string      `json:"duration_end,omitempty"`   //设定展示结束时间，格式为yyyy-MM-dd HH:mm:ss  可选
	NotifyStyle   int         `json:"notify_style,omitempty"`   //	通知栏消息布局样式(0 系统样式 1 个推样式) 默认为0  可选
	Style         interface{} `json:"style,omitempty"`          //通知栏消息布局样式(0 系统样式 1 个推样式) 默认为0  可选
}

//透传消息模板
//透传消息是指消息传递到客户端只有消息内容，展现的形式由客户端自行定义。客户端可自定义通知的展现形式，可以自定义通知到达后的动作，或者不做任何展现。IOS推送也使用该模板
// http://docs.getui.com/getui/server/rest/template/?id=doc-title-3
type Transmission struct {
	TransmissionType    bool    `json:"transmission_type"`        //	收到消息是否立即启动应用，true为立即启动，false则广播等待启动，默认是否  可选
	TransmissionContent string  `json:"transmission_content"`     //透传内容
	DurationBegin       string  `json:"duration_begin,omitempty"` //	设定展示开始时间，格式为yyyy-MM-dd HH:mm:ss  可选
	DurationEnd         string  `json:"duration_end,omitempty"`   //设定展示结束时间，格式为yyyy-MM-dd HH:mm:ss  可选
	Notify              *Notify `json:"notify,omitempty"`         //多厂商推送透传消息带通知
}
