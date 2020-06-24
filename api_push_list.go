package getui

import (
	"encoding/json"
)

//http://docs.getui.com/getui/server/rest/push/#doc-title-1
type PushListParam struct {
	Cid        []string `json:"cid"` //cid为cid list，与alias list二选一
	Alias      []string `json:"alias"`
	Taskid     string   `json:"taskid"`
	NeedDetail bool     `json:"need_detail"`
}

type PushListResult struct {
	Result       string            `json:"result"`
	Taskid       string            `json:"taskid"`
	Desc         string            `json:"desc"`
	CidDetails   map[string]string `json:"cid_details"`
	AliasDetails map[string]string `json:"alias_details"`
}

//群推
func PushList(appId string, authToken string, param *PushListParam) (*PushListResult, error) {
	url := API_URL + appId + "/push_list"
	bodyByte, err := MakeReqBody(param)
	if err != nil {
		return nil, err
	}

	result, err := SendPost(url, authToken, bodyByte)
	if err != nil {
		return nil, err
	}

	//fmt.Println("PushList result: " + result)

	var pushListResult *PushListResult
	if err := json.Unmarshal([]byte(result), &pushListResult); err != nil {
		return nil, err
	}

	return pushListResult, err
}
