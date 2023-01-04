package log

import "net/http"

type LogResponseParam struct {
	ThirdParty     string
	ResponseHeader http.Header
	ResponseBody   interface{}
	ResponseCode   string
	Timestamp      string
}
