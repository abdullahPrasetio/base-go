package exceptions

import (
	"net/http"

	"github.com/abdullahPrasetio/base-go/constants"
	modellog "github.com/abdullahPrasetio/base-go/models/log"
	httpRes "github.com/abdullahPrasetio/base-go/utils/http"
	"github.com/abdullahPrasetio/base-go/utils/log"
	"github.com/gin-gonic/gin"
)

type QueryError struct {
	Error string
}

func NewQueryError(error string) QueryError {
	return QueryError{Error: error}
}

func errorQuery(c *gin.Context, err any) bool {
	param := modellog.LogResponseParam{
		ThirdParty:     "",
		ResponseCode:   "99",
		ResponseBody:   "",
		ResponseHeader: c.Writer.Header(),
	}
	exception, ok := err.(NotFoundError)
	if ok {
		rc := constants.ErrorQueryDB
		response := httpRes.APIResponseError(constants.GetError(rc).Description, rc, exception.Error)
		param.ResponseBody = response
		param.ResponseCode = rc
		log.LogResponse(param)
		c.JSON(http.StatusInternalServerError, response)
		return true
	} else {
		return false
	}
}
