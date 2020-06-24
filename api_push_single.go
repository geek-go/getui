package getui

import (
	"encoding/json"
)

type PushSingleParam struct {
	Message      *Message      `json:"message"`
	Notification *Notification `json:"notification,omitempty"`
	Link         *Link         `json:"link,omitempty"`
	Notypopload  *NotyPopload  `json:"notypopload,omitempty"`
	Transmission *Transmission `json:"transmission,omitempty"`
	PushInfo     *PushInfo     `json:"push_info,omitempty"`
	Cid          string        `json:"cid,omitempty"`
	Alias        string        `json:"alias,omitempty"`
	RequestId    string        `json:"requestid"`
}

type PushSingleResult struct {
	Result string `json:"result"` //ok 鉴权成功
	TaskId string `json:"taskid"` //任务标识号
	Desc   string `json:"desc"`   //错误信息描述
	Status string `json:"status"` //推送结果successed_offline 离线下发successed_online 在线下发successed_ignore 非活跃用户不下发
}

//单推
func PushSingle(appId string, authToken string, param *PushSingleParam) (*PushSingleResult, error) {

	url := API_URL + appId + "/push_single"
	bodyByte, err := MakeReqBody(param)
	if err != nil {
		return nil, err
	}

	result, err := SendPost(url, authToken, bodyByte)
	if err != nil {
		return nil, err
	}

	var pushSingleResult *PushSingleResult
	if err := json.Unmarshal([]byte(result), &pushSingleResult); err != nil {
		return nil, err
	}

	return pushSingleResult, err
}
