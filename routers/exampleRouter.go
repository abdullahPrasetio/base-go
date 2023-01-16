package routers

/********************************************************************************
* Temancode Example Router Package                                              *
*                                                                               *
* Version: 1.0.0                                                                *
* Date:    2023-01-05                                                           *
* Author:  Waluyo Ade Prasetio                                                  *
* Github:  https://github.com/abdullahPrasetio                                  *
********************************************************************************/

import (
	"github.com/abdullahPrasetio/base-go/controllers/api/v1/example"
	example2 "github.com/abdullahPrasetio/base-go/models/example"
	"github.com/gin-gonic/gin"
)

func (r *routes) addExampleRoute(rg *gin.RouterGroup) {
	repository := example2.NewRepository(r.db)
	service := example2.NewService(repository)
	controller := example.NewController(service, r.validatorFormatter, r.validate)

	example := rg.Group("example")
	example.POST("/create", controller.CreateEmploye)
	example.GET("/", controller.GetAllEmployees)
	example.GET("/from-api", controller.GetFromApi)
}
