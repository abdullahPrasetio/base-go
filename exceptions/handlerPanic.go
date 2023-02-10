package exceptions

import (
	"net/http"

	modellog "github.com/abdullahPrasetio/base-go/models/log"
	httpRes "github.com/abdullahPrasetio/base-go/utils/http"
	"github.com/abdullahPrasetio/base-go/utils/log"
	formatter "github.com/abdullahPrasetio/validation-formatter"
	formatterLang "github.com/abdullahPrasetio/validation-formatter/lang"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func ErrorHandlerRecovery(c *gin.Context, err any) {
	// goErr := errors.Wrap(err, 2)
	ok := errorValidation(c, err)
	if ok {
		return
	}

	if errorDataNotFound(c, err) {
		return
	}
	if errorQuery(c, err) {
		return
	}

	internalServerError(c, err)
}

func errorValidation(c *gin.Context, err any) bool {
	param := modellog.LogResponseParam{
		ThirdParty:     "",
		ResponseCode:   "99",
		ResponseBody:   "",
		ResponseHeader: c.Writer.Header(),
	}
	NewLang := formatterLang.NewLanguage("ID")
	newvalidate := formatter.NewValidateFormatter(NewLang, "SnackCase")
	exception, ok := err.(validator.ValidationErrors)
	if ok {
		res := newvalidate.GetErrorMsgValidation(exception)
		response := httpRes.APIResponseError("Error validate", "08", res)
		param.ResponseBody = response
		param.ResponseCode = "08"
		log.LogResponse(param)
		c.JSON(http.StatusUnprocessableEntity, response)
		return true
	} else {
		return false
	}
}

func internalServerError(c *gin.Context, err any) bool {
	param := modellog.LogResponseParam{
		ThirdParty:     "",
		ResponseCode:   "99",
		ResponseBody:   "",
		ResponseHeader: c.Writer.Header(),
	}
	exceptions, _ := err.(error)

	response := httpRes.APIResponseError("General Error", "99", exceptions.Error())
	param.ResponseBody = response
	log.LogResponse(param)

	c.JSON(http.StatusInternalServerError, response)
	return true
}
