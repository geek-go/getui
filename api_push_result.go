package getui

import (
	"encoding/json"
)

type PushResultParam struct {
	Taskidlist []string `json:"taskIdList"` //查询的任务结果列表
}

type GtFeedBack struct {
	Feedback  int    `json:"feedback"`
	Displayed int    `json:"displayed"`
	Result    string `json:"result"`
	Sent      int    `json:"sent"`
	Clicked   int    `json:"clicked"`
}

type PushResult struct {
	Result string `json:"result"`
	Data   []struct {
		Taskid     string `json:"taskId"`
		MsgTotal   int    `json:"msgTotal"`
		MsgProcess int    `json:"msgProcess"`
		ClickNum   int    `json:"clickNum"`
		PushNum    int    `json:"pushNum"`
		APN        string `json:"APN,omitempty"`
		GT         string `json:"GT"`
	} `json:"data"`
}

//获取推送结果接口
func GetPushResult(appId string, auth_token string, param *PushResultParam) (*PushResult, error) {

	url := API_URL + appId + "/push_result"
	bodyByte, err := MakeReqBody(param)
	if err != nil {
		return nil, err
	}

	result, err := SendPost(url, auth_token, bodyByte)
	if err != nil {
		return nil, err
	}

	//fmt.Println(result)

	var pushResult *PushResult
	if err := json.Unmarshal([]byte(result), &pushResult); err != nil {
		return nil, err
	}

	return pushResult, err
}
