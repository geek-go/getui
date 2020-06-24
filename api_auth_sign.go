package getui

import (
	"encoding/json"
)

type SignParam struct {
	Sign      string `json:"sign"`
	Timestamp string `json:"timestamp"`
	AppKey    string `json:"appkey"`
}

//token
type SignResult struct {
	Result    string `json:"result"`
	AuthToken string `json:"auth_token"`
}

//获取Auth签名
//http://docs.getui.com/getui/server/rest/other_if/
func GetGeTuiToken(appID string, appKey string, masterSecret string) (string, error) {

	signStr, timestamp := Signature(appKey, masterSecret)

	param := &SignParam{
		Sign:      signStr,
		Timestamp: timestamp,
		AppKey:    appKey,
	}

	bodyByte, err := json.Marshal(param)
	if err != nil {
		return "", err
	}

	url := API_URL + appID + "/auth_sign"
	result, err := SendPost(url, "", bodyByte)
	if err != nil {
		return "", err
	}

	tokenResult := new(SignResult)
	if err := json.Unmarshal([]byte(result), &tokenResult); err != nil {
		return "", err
	}

	return tokenResult.AuthToken, nil
}
