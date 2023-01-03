package example

import (
	"fmt"
	example2 "github.com/abdullahPrasetio/base-go/models/example"
	formattervalidator "github.com/abdullahPrasetio/validation-formatter"
	"github.com/gin-gonic/gin"
	"net/http"
)

type exampleController struct {
	service            example2.Service
	validatorFormatter formattervalidator.ValidateFormatter
}

func NewController(service example2.Service, validatorFormatter formattervalidator.ValidateFormatter) *exampleController {
	return &exampleController{
		service:            service,
		validatorFormatter: validatorFormatter,
	}
}

func (h *exampleController) CreateEmploye(c *gin.Context) {
	var input example2.RequestCreateEmployee
	err := c.ShouldBindJSON(&input)
	if err != nil {
		fmt.Println(err)
		response := h.validatorFormatter.GetErrorMsgValidation(err)
		fmt.Println(response)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}
}
