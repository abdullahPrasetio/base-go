package example

import (
	"database/sql"
	"fmt"
	"testing"

	"github.com/abdullahPrasetio/base-go/configs"
	"github.com/abdullahPrasetio/base-go/database"
	"github.com/abdullahPrasetio/base-go/models/example"
	"github.com/abdullahPrasetio/base-go/utils/log"
	"github.com/gin-gonic/gin"
)

func init() {
	// dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println(dir)
	config, _ := configs.LoadConfig("../..")
	gin.SetMode(config.GIN_MODE)

	log.LoadLogger()
}

var DB *sql.DB
var REPOSITORY example.Repository

func TestConnectionDatabase(t *testing.T) {
	db, err := database.GetConnection()
	if err != nil {
		t.Error(err.Error())
	}
	DB = db
	fmt.Println("Test Connection Database Success")

}

func TestNewRepository(t *testing.T) {
	// repository := example.NewRepository(DB)
	// REPOSITORY = repository
	// fmt.Println("Test New Repository Success")
}

func TestRepositoryGetAll(t *testing.T) {
	repository := example.NewRepository(DB)
	result, err := repository.GetAll()
	if err != nil {
		t.Error(err.Error())
	}

	if len(result) < 1 {
		t.Error("Data not found")
	}
	fmt.Println(result)
	fmt.Println("Test Repository Get All Success")
}

func TestServiceGetALl(t *testing.T) {
	// repository := ""
}
