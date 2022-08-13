package weixinmp

import (
	"github.com/fastwego/offiaccount"
	"github.com/fastwego/offiaccount/apis/account"
)

type Account struct {
	CTX *offiaccount.OffiAccount
}

func (a *Account) CreateQRCode(in interface{}) interface{} {
	return Call(account.CreateQRCode, in, a.CTX)
}
