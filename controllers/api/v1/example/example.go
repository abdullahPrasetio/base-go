package example

/********************************************************************************
* Temancode Example Controller Package                                          *
*                                                                               *
* Version: 1.0.0                                                                *
* Date:    2023-01-05                                                           *
* Author:  Waluyo Ade Prasetio                                                  *
* Github:  https://github.com/abdullahPrasetio                                  *
********************************************************************************/

import (
	"net/http"

	model "github.com/abdullahPrasetio/base-go/models/example"
	modellog "github.com/abdullahPrasetio/base-go/models/log"
	resp "github.com/abdullahPrasetio/base-go/utils/http"
	"github.com/abdullahPrasetio/base-go/utils/log"
	formattervalidator "github.com/abdullahPrasetio/validation-formatter"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

type exampleController struct {
	service            model.Service
	validatorFormatter formattervalidator.ValidateFormatter
}

func NewController(service model.Service, validatorFormatter formattervalidator.ValidateFormatter) *exampleController {
	return &exampleController{
		service:            service,
		validatorFormatter: validatorFormatter,
	}
}

func (h *exampleController) CreateEmploye(c *gin.Context) {
	var input model.RequestEmployee
	err := c.ShouldBindBodyWith(&input, binding.JSON)
	param := modellog.LogResponseParam{
		ThirdParty:     "",
		ResponseCode:   "99",
		ResponseBody:   "",
		ResponseHeader: c.Writer.Header(),
	}
	if err != nil {
		res := h.validatorFormatter.GetErrorMsgValidation(err)
		response := resp.APIResponseError("Error validate", "08", res)
		param.ResponseBody = response
		log.LogResponse(param)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}
	result, err := h.service.Create(input)
	if err != nil {
		response := resp.APIResponseError("Error to create Employee", "99", err)
		param.ResponseBody = response
		log.LogResponse(param)
		log.Logger.Error(err)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := resp.APIResponseSuccess("Successfully created Employee", result)
	param.ResponseBody = response
	log.Logger.Info("Successfully created employee")
	log.LogResponse(param)
	c.JSON(http.StatusOK, response)
	return
}

func (h *exampleController) GetAllEmployees(c *gin.Context) {
	param := modellog.LogResponseParam{
		ThirdParty:     "",
		ResponseCode:   "99",
		ResponseBody:   "",
		ResponseHeader: c.Writer.Header(),
	}
	result, err := h.service.GetAll()
	if err != nil {
		response := resp.APIResponseError("Error to get Employee", "99", err)
		param.ResponseBody = response
		log.LogResponse(param)
		log.Logger.Error(err)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	response := resp.APIResponseSuccess("Successfully get Employee", result)
	param.ResponseBody = response
	log.Logger.Info("Successfully get employee")
	log.LogResponse(param)
	c.JSON(http.StatusOK, response)
	return
}
