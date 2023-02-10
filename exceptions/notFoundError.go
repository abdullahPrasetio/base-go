package exceptions

import (
	"net/http"

	"github.com/abdullahPrasetio/base-go/constants"
	modellog "github.com/abdullahPrasetio/base-go/models/log"
	httpRes "github.com/abdullahPrasetio/base-go/utils/http"
	"github.com/abdullahPrasetio/base-go/utils/log"
	"github.com/gin-gonic/gin"
)

type NotFoundError struct {
	Error string
}

func NewDataNotFoundError(error string) NotFoundError {
	return NotFoundError{Error: error}
}

func errorDataNotFound(c *gin.Context, err any) bool {
	param := modellog.LogResponseParam{
		ThirdParty:     "",
		ResponseCode:   "99",
		ResponseBody:   "",
		ResponseHeader: c.Writer.Header(),
	}
	exception, ok := err.(NotFoundError)
	if ok {
		rc := constants.ErrorNotFound
		response := httpRes.APIResponseError(constants.GetError(rc).Description, rc, exception.Error)
		param.ResponseBody = response
		param.ResponseCode = rc
		log.LogResponse(param)
		c.JSON(http.StatusNotFound, response)
		return true
	} else {
		return false
	}
}
