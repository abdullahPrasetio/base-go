package routers

/********************************************************************************
* Temancode Example Routers Package                                             *
*                                                                               *
* Version: 1.0.0                                                                *
* Date:    2023-01-05                                                           *
* Author:  Waluyo Ade Prasetio                                                  *
* Github:  https://github.com/abdullahPrasetio                                  *
********************************************************************************/

import (
	"database/sql"
	"errors"
	"net/http"

	"github.com/abdullahPrasetio/base-go/constants"
	"github.com/abdullahPrasetio/base-go/database"
	"github.com/abdullahPrasetio/base-go/library/newvalidator"
	"github.com/abdullahPrasetio/base-go/middleware"
	http2 "github.com/abdullahPrasetio/base-go/utils/http"
	log "github.com/abdullahPrasetio/base-go/utils/log"
	formatter "github.com/abdullahPrasetio/validation-formatter"
	formatterLang "github.com/abdullahPrasetio/validation-formatter/lang"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"

	//formattervalidator "github.com/abdullahPrasetio/validation-formatter"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type routes struct {
	router             *gin.Engine
	db                 *sql.DB
	validatorFormatter formatter.ValidateFormatter
	validate           *validator.Validate
}

func SetupRouter() *gin.Engine {

	NewLang := formatterLang.NewLanguage("ID")
	newvalidate := formatter.NewValidateFormatter(NewLang, "SnackCase")
	db, err := database.GetConnection()
	if err != nil {
		log.Logger.Info("Error connecting to database...")
	}
	binding.Validator = new(newvalidator.DefaultValidator)
	r := routes{
		router:             gin.Default(),
		db:                 db,
		validatorFormatter: newvalidate,
		validate:           formatter.Validate,
	}
	// Memasang middleware bawaan gin
	r.router.Use(gin.Logger())   // Logger
	r.router.Use(gin.Recovery()) // Jika error panic maka akan recover
	r.router.Use(cors.Default())
	api := r.router.Group(constants.ServerDefaultRoute, middleware.WriteRequestLog(), middleware.AddDefaultHeader())
	r.addExampleRoute(api)
	r.router.NoRoute(func(c *gin.Context) {
		c.JSON(404, http2.APIResponseError("Page Not Found", constants.ErrorNotFound, errors.New("Page Not Found").Error()))
	})
	api.GET("/healtz", checkHealtz)
	return r.router
}

func checkHealtz(c *gin.Context) {
	c.JSON(http.StatusOK, http2.APIResponseSuccess("success", "ok"))
}
