package routers

import (
	"github.com/abdullahPrasetio/base-go/controllers/api/v1/example"
	example2 "github.com/abdullahPrasetio/base-go/models/example"
	"github.com/gin-gonic/gin"
)

func (r *routes) addExampleRoute(rg *gin.RouterGroup) {
	repository := example2.NewRepository(r.db)
	service := example2.NewService(repository)
	controller := example.NewController(service, r.validatorFormatter)

	example := rg.Group("example")
	example.POST("/create", controller.CreateEmploye)
}
