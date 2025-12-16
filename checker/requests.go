package checker

import (
	"github.com/valyala/fasthttp"
)

func make_request(ip string) bool {
	_, _, err := fasthttp.Get(nil, ip)
	return err == nil

}
