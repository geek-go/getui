package getui

import (
	"encoding/json"
)

type Condition []AppCondition

type AppCondition struct {
	Key     string   `json:"key"`      //筛选条件类型名称(省市region,手机类型phonetype,用户标签tag)
	Values  []string `json:"values"`   //筛选参数
	OptType int      `json:"opt_type"` //筛选参数的组合，0:取参数并集or，1：交集and，2：相当与not in {参数1，参数2，....}
}

//手机类型
const PHONE_TYPE = "phonetype"

//地区
const REGION = "region"

//自定义tag
const TAG = "tag"

type PushAppParam struct {
	Message       *Message      `json:"message"`
	Notification  *Notification `json:"notification"`
	Link          *Link         `json:"link,omitempty"`
	Notypopload   *NotyPopload  `json:"notypopload,omitempty"`
	Transmission  *Transmission `json:"transmission,omitempty"`
	PushInfo      *PushInfo     `json:"push_info,omitempty"`
	Condition     *Condition    `json:"condition,omitempty"`
	RequestId     string        `json:"requestid"`                //请求唯一标识
	Speed         int64         `json:"speed,omitempty"`          //推送速度控制。设置群推接口的推送速度，单位为条/秒，例如填写100，则为100条/秒。仅对指定应用群推接口有效。
	TaskName      string        `json:"task_name,omitempty"`      //任务别名
	DurationBegin string        `json:"duration_begin,omitempty"` //设定展示开始时间，格式为yyyy-MM-dd HH:mm:ss
	DurationEnd   string        `json:"duration_end,omitempty"`   //设定展示结束时间，格式为yyyy-MM-dd HH:mm:ss
}

type PushAppResult struct {
	Result string `json:"result"`
	Taskid string `json:"taskid"`
	Desc   string `json:"desc"`
}

//全推,http://docs.getui.com/getui/server/rest/push/#doc-title-2
func PushApp(appId string, authToken string, param *PushAppParam) (*PushAppResult, error) {
	url := API_URL + appId + "/push_app"
	bodyByte, err := MakeReqBody(param)
	if err != nil {
		return nil, err
	}

	result, err := SendPost(url, authToken, bodyByte)
	if err != nil {
		return nil, err
	}

	var pushAppResult *PushAppResult
	if err := json.Unmarshal([]byte(result), &pushAppResult); err != nil {
		return nil, err
	}

	return pushAppResult, err
}
