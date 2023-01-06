package log

/********************************************************************************
* Temancode Example Log Models Package                                          *
*                                                                               *
* Version: 1.0.0                                                                *
* Date:    2023-01-05                                                           *
* Author:  Waluyo Ade Prasetio                                                  *
* Github:  https://github.com/abdullahPrasetio                                  *
********************************************************************************/

import "net/http"

type LogResponseParam struct {
	ThirdParty     string
	ResponseHeader http.Header
	ResponseBody   interface{}
	ResponseCode   string
	Timestamp      string
}
