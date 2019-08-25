package getui

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"testing"
	"time"
)

var Cfg = &GetuiConfig{
	AppId:        "",
	AppSecret:    "",
	AppKey:       "",
	MasterSecret: "",
}

//测试单推
func TestGeTui_SendByCid(t *testing.T) {
	igetui, err := NewGeTui(Cfg)
	if err != nil {
		t.Error(err)
		os.Exit(1)
	}

	cid := "b1f0c722b141e4ee718f2b386b39c683"
	payLoad := Payload{"这是测试title", "这是测试内容", "1", ""}
	err = igetui.SendByCid(cid, &payLoad)
	if err != nil {
		t.Error(err)
	} else {
		t.Log("ok")
	}
}

//测试群推
func TestGeTui_SendByCids(t *testing.T) {
	igetui, err := NewGeTui(Cfg)
	if err != nil {
		t.Error(err)
		os.Exit(1)
	}

	cids := []string{"b1f0c722b141e4ee718f2b386b39c683"}
	payLoad := Payload{"这是测试title", "这是测试内容", "1", ""}
	err = igetui.SendByCids(cids, &payLoad)
	if err != nil {
		t.Error(err)
	} else {
		t.Log("ok")
	}
}

//测试全推
func TestGeTui_SendAll(t *testing.T) {
	igetui, err := NewGeTui(Cfg)
	if err != nil {
		t.Error(err)
		os.Exit(1)
	}

	payLoad := Payload{"这是测试title", "这是测试内容", "1", ""}
	err = igetui.SendAll(&payLoad)
	if err != nil {
		t.Error(err)
	} else {
		t.Log("ok")
	}
}

//配置
type GetuiConfig struct {
	AppId        string `toml:"app_id"`
	AppKey       string `toml:"app_key"`
	AppSecret    string `toml:"app_secret"`
	MasterSecret string `toml:"master_secret"`
}

//消息payload，根据业务自定义
type Payload struct {
	PushTitle    string `json:"push_title"`
	PushBody     string `json:"push_body"`
	IsShowNotify string `json:"is_show_notify"`
	Ext          string `json:"ext"`
}

//个推
type GeTuiPush struct {
	Config *GetuiConfig
}

//获取个推实例
func NewGeTui(config *GetuiConfig) (*GeTuiPush, error) {

	if config.AppId == "" || config.AppSecret == "" || config.AppKey == "" {
		return nil, errors.New("请检查配置")
	}

	gt := &GeTuiPush{
		Config: config,
	}

	return gt, nil
}

//透传模板
func IGtTransmissionTemplate(payload *Payload) (*Transmission, *PushInfo, error) {

	payloadByte, err := json.Marshal(payload)
	if err != nil {
		return nil, nil, err
	}

	//notify:多厂商推送透传消息带通知配置
	//notify := &getui.Notify{
	//	Title:   payload.PushTitle,
	//	Content: payload.GetPushBody(),
	//	Intent:  "",
	//	Type:    "1",
	//}

	//实例化透传模板
	template := &Transmission{
		TransmissionType:    false,
		TransmissionContent: string(payloadByte), //安卓使用
		//Notify:              notify,
	}

	//设置APNS
	apn := Apns{
		Category: "ACTIONABLE",
	}
	if payload.IsShowNotify == "1" {
		alertmsg := &Alert{}
		alertmsg.Title = payload.PushTitle
		alertmsg.Body = payload.PushBody

		apn.Alert = alertmsg
		apn.Sound = ""
		apn.AutoBadge = "+1" //角标
		apn.ContentAvailable = 0
	} else { //静默推送
		apn.Sound = "com.gexin.ios.silence"
		apn.AutoBadge = "+0" //角标
		apn.ContentAvailable = 1
	}

	//pushInfo
	pushInfo := PushInfo{}
	pushInfo["aps"] = apn
	pushInfo["payload"] = string(payloadByte) //iOS使用

	return template, &pushInfo, nil
}

//根据用户cid推送
func (g *GeTuiPush) SendByCid(cid string, payload *Payload) error {

	//获取签名
	token, _ := GetGeTuiToken(g.Config.AppId, g.Config.AppKey, g.Config.MasterSecret)

	//消息体
	message := GetMessage()
	message.AppKey = g.Config.AppKey
	message.MsgType = MsgType.Transmission

	//推送模板
	template, pushInfo, err := IGtTransmissionTemplate(payload)
	if err != nil {
		return err
	}

	pushSingleParam := &PushSingleParam{
		Message:      message,
		Transmission: template,
		Cid:          cid,
		PushInfo:     pushInfo,
		RequestId:    time.Now().Format("20160102150405"),
	}

	res, err := PushSingle(g.Config.AppId, token, pushSingleParam)
	if err != nil {
		return err
	}

	fmt.Println(res)

	return nil
}

//根据用户cids批量推送
func (g *GeTuiPush) SendByCids(cids []string, payload *Payload) error {

	//获取签名
	token, _ := GetGeTuiToken(g.Config.AppId, g.Config.AppKey, g.Config.MasterSecret)

	//消息体
	message := GetMessage()
	message.AppKey = g.Config.AppKey
	message.MsgType = MsgType.Transmission

	//推送模板
	template, pushInfo, err := IGtTransmissionTemplate(payload)
	if err != nil {
		return err
	}

	//a. 先调用save_list_body保存消息共同体
	saveListBodyParam := &SaveListBodyParam{
		Message:      message,
		Transmission: template,
		PushInfo:     pushInfo, //必须
	}

	res, err := SaveListBody(g.Config.AppId, token, saveListBodyParam)
	if err != nil {
		return err
	}

	if res.Result != "ok" {
		return errors.New(fmt.Sprintf("获取contentId失败:%s,%s", res.Result, res.Desc))
	}

	taskid := res.TaskId
	//fmt.Println("content_id: " + taskid)

	//b. 获取到taskid后再调用群推接口推送
	pushListParam := &PushListParam{
		Cid:        cids,
		Taskid:     taskid,
		NeedDetail: true,
	}

	res2, err := PushList(g.Config.AppId, token, pushListParam)
	if err != nil {
		return err
	}

	fmt.Println(res2)

	return nil
}

//推送给所有人
func (g *GeTuiPush) SendAll(payload *Payload) error {

	//获取签名
	token, _ := GetGeTuiToken(g.Config.AppId, g.Config.AppKey, g.Config.MasterSecret)

	//消息体
	message := GetMessage()
	message.AppKey = g.Config.AppKey
	message.MsgType = MsgType.Transmission

	//推送模板
	template, pushInfo, err := IGtTransmissionTemplate(payload)
	if err != nil {
		return err
	}
	conditions := Condition{}
	conditions = append(conditions, AppCondition{
		Key:    PHONE_TYPE,
		Values: []string{"ANDROID", "IOS"},
	})
	pushAppParam := &PushAppParam{
		Message:      message,
		Transmission: template,
		PushInfo:     pushInfo,
		Condition:    &conditions,
		RequestId:    time.Now().Format("20160102150405"),
	}

	res, err := PushApp(g.Config.AppId, token, pushAppParam)
	if err != nil {
		return err
	}

	fmt.Println(res)

	return nil
}

//获取推送结果接口
func (g *GeTuiPush) GetPushResult(taskidlist []string) (*PushResult, error) {

	//获取签名
	token, _ := GetGeTuiToken(g.Config.AppId, g.Config.AppKey, g.Config.MasterSecret)

	pushParam := &PushResultParam{
		Taskidlist: taskidlist,
	}

	res, err := GetPushResult(g.Config.AppId, token, pushParam)
	if err != nil {
		return nil, err
	}

	//fmt.Println(res)

	return res, nil
}
