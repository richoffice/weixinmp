package weixinmp

import (
	"encoding/json"
	"net/url"

	"github.com/fastwego/offiaccount"
)

type WeixinMp struct {
	CTX     *offiaccount.OffiAccount
	Account *Account
	User    *User
}

type WeixinMpFunc func(ctx *offiaccount.OffiAccount, payload []byte) (resp []byte, err error)

type WeixinMpUrlFunc func(ctx *offiaccount.OffiAccount, params url.Values) (resp []byte, err error)

func Call(f WeixinMpFunc, in interface{}, ctx *offiaccount.OffiAccount) interface{} {
	data, e := json.Marshal(in)
	if e != nil {
		return e
	}
	resp, err := f(ctx, data)
	if err != nil {
		return err
	}

	out := make(map[string]interface{})
	ee := json.Unmarshal(resp, &out)
	if ee != nil {
		return ee
	}

	return out
}

func Query(f WeixinMpUrlFunc, params url.Values, ctx *offiaccount.OffiAccount) interface{} {
	resp, err := f(ctx, params)
	if err != nil {
		return err
	}

	out := make(map[string]interface{})
	ee := json.Unmarshal(resp, &out)
	if ee != nil {
		return ee
	}
	return out
}

func NewWeixinMp(appid string, appkey string, encodingKey string) *WeixinMp {
	app := offiaccount.New(offiaccount.Config{
		Appid:  appid,
		Secret: appkey,
	})

	if encodingKey != "" {
		app = offiaccount.New(offiaccount.Config{
			Appid:          appid,
			Secret:         appkey,
			EncodingAESKey: encodingKey,
		})
	}

	return &WeixinMp{
		CTX: app,
		Account: &Account{
			CTX: app,
		},
		User: &User{
			CTX: app,
		},
	}
}
