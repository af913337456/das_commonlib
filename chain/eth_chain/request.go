package eth_chain

import (
	"fmt"
	"github.com/parnurzeal/gorequest"
	"net/http"
	"sync"
	"time"
)

type Error struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type BaseResp struct {
	JsonRpc string      `json:"jsonrpc"`
	Id      int32       `json:"id"`
	Result  interface{} `json:"result"`
	Error   Error       `json:"error"`
}

var lock sync.Mutex

func (b *BaseResp) Request(url, method string) error {
	lock.Lock()
	defer lock.Unlock()
	ret, body, errs := gorequest.New().Timeout(time.Second*30).Post(url).Set("Content-Type", "application/json").Send(method).EndStruct(b)
	if errs != nil || ret.StatusCode != http.StatusOK {
		if len(body) > 100 {
			body = body[:100]
		}
		return fmt.Errorf("request err:%s %v %s", method, errs, string(body))
	}
	return nil
}
