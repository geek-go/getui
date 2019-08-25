package getui

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

//post请求
func SendPost(url string, auth_token string, bodyByte []byte) (string, error) {
	//创建客户端实例
	client := &http.Client{
		Timeout: 10 * time.Second,
	}

	body := bytes.NewBuffer(bodyByte)

	//fmt.Println(body)

	//创建请求实例
	req, err := http.NewRequest("POST", url, body)
	if err != nil {
		return "", err
	}

	req.Header.Add("authtoken", auth_token)
	req.Header.Add("Charset", "UTF-8")
	req.Header.Add("Content-Type", "application/json")

	//发起请求
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}

	defer resp.Body.Close()

	//读取响应
	result, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		fmt.Println("request getui fail.", resp)
		return "", err
	}

	return string(result), nil
}

//生成请求参数对应的JSON
func MakeReqBody(parmar interface{}) ([]byte, error) {

	body, err := json.Marshal(parmar)
	if err != nil {
		return nil, err
	}

	return body, nil
}
