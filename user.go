package weixinmp

import (
	"net/url"

	"github.com/fastwego/offiaccount"
	"github.com/fastwego/offiaccount/apis/user"
)

type User struct {
	CTX *offiaccount.OffiAccount
}

func (a *Account) GetList(params url.Values) interface{} {
	return Query(user.Get, params, a.CTX)
}

func (a *Account) GetUserInfo(params url.Values) interface{} {
	return Query(user.GetUserInfo, params, a.CTX)
}
