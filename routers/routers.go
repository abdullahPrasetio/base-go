package routers

import (
	"database/sql"
	"github.com/abdullahPrasetio/base-go/constants"
	"github.com/abdullahPrasetio/base-go/database"
	"github.com/abdullahPrasetio/base-go/helpers"
	"github.com/abdullahPrasetio/base-go/library/newvalidator"
	log "github.com/abdullahPrasetio/base-go/utils/log"
	formatter "github.com/abdullahPrasetio/validation-formatter"
	formatterLang "github.com/abdullahPrasetio/validation-formatter/lang"
	"github.com/gin-gonic/gin/binding"
	"net/http"
	//formattervalidator "github.com/abdullahPrasetio/validation-formatter"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type routes struct {
	router             *gin.Engine
	db                 *sql.DB
	validatorFormatter formatter.ValidateFormatter
}

func SetupRouter() *gin.Engine {

	NewLang := formatterLang.NewLanguage("ID")
	newvalidate := formatter.NewValidateFormatter(NewLang, "SnackCase")
	db, err := database.GetConnection()
	if err != nil {
		log.Logger.Info("server shutdown gracefully")
	}
	binding.Validator = new(newvalidator.DefaultValidator)
	r := routes{
		router:             gin.Default(),
		db:                 db,
		validatorFormatter: newvalidate,
	}
	// Memasang middleware bawaan gin
	r.router.Use(gin.Logger())   // Logger
	r.router.Use(gin.Recovery()) // Jika error panic maka akan recover
	r.router.Use(cors.Default())
	api := r.router.Group(constants.ServerDefaultRoute)
	r.addExampleRoute(api)
	api.GET("/healtz", checkHealtz)
	return r.router
}

func checkHealtz(c *gin.Context) {
	c.JSON(http.StatusOK, helpers.APIResponseSuccess("success", "ok"))
}
