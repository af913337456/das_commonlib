package wx_chain

import (
	"context"
	"crypto/x509"
	"encoding/json"
	"fmt"
	"github.com/wechatpay-apiv3/wechatpay-go/core"
	"github.com/wechatpay-apiv3/wechatpay-go/core/option"
	"github.com/wechatpay-apiv3/wechatpay-go/utils"
	"io/ioutil"
	"net/http"
	"time"
)

type Client struct {
	ctx                 context.Context
	client              *core.Client
	mchID               string // 直连商户号
	mchCertSerialNumber string // 商户证书序列号
	notifyUrl           string // 支付回调通知URL
	appID               string // 应用ID
	apiV3Key            string // 通知解密密钥
	dev                 bool   // 开发环境
}

type InitConfig struct {
	MchID                 string
	MchCertSerialNumber   string
	PrivateKeyPath        string
	WechatCertificatePath string
	AppID                 string
	ApiV3Key              string
	NotifyUrl             string
	Dev                   bool
}

func Initialize(p InitConfig) (*Client, error) {
	// 加载商户私钥
	privateKey, err := utils.LoadPrivateKeyWithPath(p.PrivateKeyPath)
	if err != nil {
		return nil, fmt.Errorf("LoadPrivateKeyWithPath err:%s", err.Error())
	}
	// 加载微信支付平台证书
	wechatPayCertificate, err := utils.LoadCertificateWithPath(p.WechatCertificatePath)
	if err != nil {
		return nil, fmt.Errorf("LoadCertificateWithPath err:%s", err.Error())
	}
	ctx := context.Background()
	opts := []option.ClientOption{
		option.WithMerchant(p.MchID, p.MchCertSerialNumber, privateKey), // 设置商户信息，用于生成签名信息
		option.WithWechatPay([]*x509.Certificate{wechatPayCertificate}), // 设置微信支付平台证书信息，对回包进行校验
		option.WithHTTPClient(&http.Client{}),                           // 可以不设置
		option.WithTimeout(30 * time.Second),                            // 自行进行超时时间配置
		option.WithHeader(&http.Header{}),                               // 可以自行设置Header
		//option.WithoutValidator(),
	}
	client, err := core.NewClient(ctx, opts...)
	if err != nil {
		return nil, fmt.Errorf("NewClient err:%s", err.Error())
	}
	return &Client{ctx: ctx, client: client, mchID: p.MchID, appID: p.AppID, apiV3Key: p.ApiV3Key, notifyUrl: p.NotifyUrl, dev: p.Dev}, nil
}

// 校验回包
func checkResponse(response *http.Response) ([]byte, error) {
	// 校验回包内容是否有逻辑错误
	if err := core.CheckResponse(response); err != nil {
		return nil, fmt.Errorf("check response err:%s", err.Error())
	}
	// 读取回包信息
	if body, err := ioutil.ReadAll(response.Body); err != nil {
		return nil, fmt.Errorf("read response body err:%s", err.Error())
	} else {
		return body, nil
	}
}

type RespCreateOrder struct {
	CodeUrl string `json:"code_url"`
}

// Native支付统一下单
func (c *Client) CreateOrder(orderId, description string, amount int64) (RespCreateOrder, error) {
	if c.dev {
		amount = 1
	}
	var resp RespCreateOrder
	//设置请求地址
	URL := "https://api.mch.weixin.qq.com/v3/pay/transactions/native"
	//设置请求信息,此处也可以使用结构体来进行请求
	mapInfo := map[string]interface{}{
		"mchid":        c.mchID,
		"out_trade_no": orderId,
		"appid":        c.appID,
		"description":  description,
		"notify_url":   c.notifyUrl,
		"time_expire":  time.Now().Add(time.Hour * 24).Format("2006-01-02T15:04:05+08:00"),
		"amount": map[string]interface{}{
			"total":    amount, //单位 分
			"currency": "CNY",
		},
	}
	// 发起请求
	response, err := c.client.Post(c.ctx, URL, mapInfo)
	if err != nil {
		return resp, fmt.Errorf("CreateOrder client post err:%s", err.Error())
	}
	// 校验回包
	body, err := checkResponse(response)
	if err != nil {
		return resp, fmt.Errorf("CreateOrder checkResponse err:%s", err.Error())
	}
	if err := json.Unmarshal(body, &resp); err != nil {
		return resp, fmt.Errorf("CreateOrder Unmarshal err:%s", err.Error())
	}
	return resp, nil
}

type RespQueryOrder struct {
	OutTradeNo     string `json:"out_trade_no"`
	TransactionId  string `json:"transaction_id"`
	TradeType      string `json:"trade_type"`
	TradeStateDesc string `json:"trade_state_desc"`
	TradeState     string `json:"trade_state"`
	SuccessTime    string `json:"success_time"`
	Payer          struct {
		Openid string `json:"openid"`
	} `json:"payer"`
	MchId    string `json:"mchid"`
	BankType string `json:"bank_type"`
	AppId    string `json:"appid"`
	Attach   string `json:"attach"`
	Amount   struct {
		Currency      string `json:"currency"`
		PayerCurrency string `json:"payer_currency"`
		PayerTotal    int64  `json:"payer_total"`
		Total         int64  `json:"total"`
	} `json:"amount"`
}

// 查询订单
func (c *Client) QueryOrder(orderId string) (RespQueryOrder, error) {
	var resp RespQueryOrder
	//设置请求地址
	URL := "https://api.mch.weixin.qq.com/v3/pay/transactions/out-trade-no/%s?mchid=%s"
	URL = fmt.Sprintf(URL, orderId, c.mchID)
	// 发起请求
	response, err := c.client.Get(c.ctx, URL)
	if err != nil {
		return resp, fmt.Errorf("QueryOrder client get err:%s", err.Error())
	}
	// 校验回包
	body, err := checkResponse(response)
	if err != nil {
		return resp, fmt.Errorf("QueryOrder checkResponse err:%s", err.Error())
	}
	if err := json.Unmarshal(body, &resp); err != nil {
		return resp, fmt.Errorf("QueryOrder Unmarshal err:%s", err.Error())
	}
	return resp, nil
}

// 关闭订单
func (c *Client) CloseOrder(txID string) error {
	//设置请求地址
	URL := "https://api.mch.weixin.qq.com/v3/pay/transactions/out-trade-no/%s/close"
	URL = fmt.Sprintf(URL, txID)
	//设置请求信息,此处也可以使用结构体来进行请求
	mapInfo := map[string]interface{}{
		"mchid": c.mchID,
	}
	// 发起请求
	response, err := c.client.Post(c.ctx, URL, mapInfo)
	if err != nil {
		return fmt.Errorf("CloseOrder client post err:%s", err.Error())
	}
	// 校验回包
	body, err := checkResponse(response)
	if err != nil {
		return fmt.Errorf("CloseOrder checkResponse err:%s", err.Error())
	}
	fmt.Println(string(body))
	return nil
}

type PayNotify struct {
	ID           string `json:"id"`
	CreateTime   string `json:"create_time"`
	ResourceType string `json:"resource_type"`
	EventType    string `json:"event_type"`
	Resource     struct {
		Algorithm      string `json:"algorithm"`
		Ciphertext     string `json:"ciphertext"`
		Nonce          string `json:"nonce"`
		OriginalType   string `json:"original_type"`
		AssociatedData string `json:"associated_data"`
	} `json:"resource"`
	Summary string `json:"summary"`
}

type OrderNotifyData struct {
	AppID         string `json:"appid"`
	MchID         string `json:"mchid"`
	OutTradeNo    string `json:"out_trade_no"`
	TransactionId string `json:"transaction_id"`
	TradeType     string `json:"trade_type"`
	TradeState    string `json:"trade_state"`
	BankType      string `json:"bank_type"`
	SuccessTime   string `json:"success_time"`
	Payer         struct {
		Openid string `json:"openid"`
	} `json:"payer"`
	Amount struct {
		Total         int64  `json:"total"`
		PayerTotal    int64  `json:"payer_total"`
		Currency      string `json:"currency"`
		PayerCurrency string `json:"payer_currency"`
	} `json:"amount"`
}

// 支付通知解析
func (c *Client) OrderNotifyResolve(notify PayNotify) (OrderNotifyData, error) {
	var data OrderNotifyData
	// 解密数据
	certificate, err := utils.DecryptToString(c.apiV3Key, notify.Resource.AssociatedData, notify.Resource.Nonce, notify.Resource.Ciphertext)
	if err != nil {
		return data, fmt.Errorf("DecryptToString err:%s", err.Error())
	}
	fmt.Println(certificate)
	// 解析
	if err := json.Unmarshal([]byte(certificate), &data); err != nil {
		return data, fmt.Errorf("json.Unmarshal err:%s", err.Error())
	}
	return data, nil
}

type RespApplyRefund struct {
	Amount struct {
		Currency         string `json:"currency"`
		DiscountRefund   int64  `json:"discount_refund"`
		PayerRefund      int64  `json:"payer_refund"`
		PayerTotal       int64  `json:"payer_total"`
		Refund           int64  `json:"refund"`
		SettlementRefund int64  `json:"settlement_refund"`
		SettlementTotal  int64  `json:"settlement_total"`
		Total            int64  `json:"total"`
	} `json:"amount"`
	Channel             string `json:"channel"`
	CreateTime          string `json:"create_time"`
	FundsAccount        string `json:"funds_account"`
	OutRefundNo         string `json:"out_refund_no"`
	OutTradeNo          string `json:"out_trade_no"`
	RefundId            string `json:"refund_id"`
	Status              string `json:"status"`
	TransactionId       string `json:"transaction_id"`
	UserReceivedAccount string `json:"user_received_account"`
}

// 申请退款
func (c *Client) ApplyRefund(txID, refundID, reason string, amount int64) (RespApplyRefund, error) {
	if c.dev {
		amount = 1
	}
	var resp RespApplyRefund
	//设置请求地址
	URL := "https://api.mch.weixin.qq.com/v3/refund/domestic/refunds"
	//设置请求信息,此处也可以使用结构体来进行请求
	mapInfo := map[string]interface{}{
		"transaction_id": txID,
		"out_refund_no":  refundID,
		"reason":         reason,
		"notify_url":     c.notifyUrl,
		"funds_account":  "AVAILABLE",
		"amount": map[string]interface{}{
			"refund":   amount,
			"total":    amount,
			"currency": "CNY",
		},
	}
	// 发起请求
	response, err := c.client.Post(c.ctx, URL, mapInfo)
	if err != nil {
		return resp, fmt.Errorf("ApplyRefund client post err:%s", err.Error())
	}
	// 校验回包
	body, err := checkResponse(response)
	if err != nil {
		return resp, fmt.Errorf("ApplyRefund checkResponse err:%s", err.Error())
	}
	if err := json.Unmarshal(body, &resp); err != nil {
		return resp, fmt.Errorf("ApplyRefund Unmarshal err:%s", err.Error())
	}
	return resp, nil
}

// 查询退款
func (c *Client) QueryRefund(refundId string) (RespApplyRefund, error) {
	var resp RespApplyRefund
	//设置请求地址
	URL := "https://api.mch.weixin.qq.com/v3/refund/domestic/refunds/%s?mchid=%s"
	URL = fmt.Sprintf(URL, refundId, c.mchID)
	// 发起请求
	response, err := c.client.Get(c.ctx, URL)
	if err != nil {
		return resp, fmt.Errorf("QueryRefund client get err:%s", err.Error())
	}
	// 校验回包
	body, err := checkResponse(response)
	if err != nil {
		return resp, fmt.Errorf("QueryRefund checkResponse err:%s", err.Error())
	}
	if err := json.Unmarshal(body, &resp); err != nil {
		return resp, fmt.Errorf("QueryRefund Unmarshal err:%s", err.Error())
	}
	return resp, nil
}

type RefundNotifyData struct {
	MchID               string `json:"mchid"`
	TransactionId       string `json:"transaction_id"`
	OutTradeNo          string `json:"out_trade_no"`
	RefundID            string `json:"refund_id"`
	OutRefundNo         string `json:"out_refund_no"`
	RefundStatus        string `json:"refund_status"`
	SuccessTime         string `json:"success_time"`
	UserReceivedAccount string `json:"user_received_account"`
	Amount              struct {
		Total       int64 `json:"total"`
		PayerTotal  int64 `json:"payer_total"`
		Refund      int64 `json:"refund"`
		PayerRefund int64 `json:"payer_refund"`
	} `json:"amount"`
}

// 退款通知
func (c *Client) RefundNotifyResolve(notify PayNotify) (RefundNotifyData, error) {
	var data RefundNotifyData
	// 解密数据
	certificate, err := utils.DecryptToString(c.apiV3Key, notify.Resource.AssociatedData, notify.Resource.Nonce, notify.Resource.Ciphertext)
	if err != nil {
		return data, fmt.Errorf("DecryptToString err:%s", err.Error())
	}
	fmt.Println(certificate)
	// 解析
	if err := json.Unmarshal([]byte(certificate), &data); err != nil {
		return data, fmt.Errorf("json.Unmarshal err:%s", err.Error())
	}
	return data, nil
}
