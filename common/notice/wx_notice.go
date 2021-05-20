package notice

import (
	"fmt"
	"github.com/parnurzeal/gorequest"
	"net/http"
)

const (
	NotifyWxTypeText     = "text"
	NotifyWxTypeMarkdown = "markdown"
	NotifyUrlWx          = "https://qyapi.weixin.qq.com/cgi-bin/webhook/send?key=%s"
)

var (
	DefaultKey = "" // 默认通知key
)

type NotifyDataWx struct {
	MsgType  string        `json:"msgtype"`
	Markdown NotifyContent `json:"markdown"`
	Text     NotifyContent `json:"text"`
}

type NotifyContent struct {
	Content string `json:"content"`
}

func SendNotifyWx(msgType, msg, key string) error {
	if key == "" {
		return nil
	}
	data := NotifyDataWx{
		MsgType:  msgType,
		Markdown: NotifyContent{},
		Text:     NotifyContent{},
	}
	content := NotifyContent{Content: msg}
	switch msgType {
	case NotifyWxTypeText:
		data.Text = content
	case NotifyWxTypeMarkdown:
		data.Markdown = content
	default:
		data.MsgType = NotifyWxTypeText
		data.Text = content
	}
	url := fmt.Sprintf(NotifyUrlWx, key)
	resp, _, errs := gorequest.New().Post(url).SendStruct(&data).End()
	if len(errs) == 0 {
		return fmt.Errorf("errs:%v", errs)
	} else if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("http code:%d", resp.StatusCode)
	}
	return nil
}

// 方法调用异常通知
func SendNotifyWxCallFuncErr(key, funcName, errInfo, keyInfo string) error {
	msg := `<font color="warning">方法调用异常</font>
> 方法名：%s
> 错误信息：%s
> 关键信息：%s`
	msg = fmt.Sprintf(msg, funcName, errInfo, keyInfo)
	return SendNotifyWx(NotifyWxTypeMarkdown, msg, key)
}