package api

import (
	_ "backend-boilerplate/docs"
	healthCheckController "backend-boilerplate/src/application/api/controllers/health-check"
	helloWorldController "backend-boilerplate/src/application/api/controllers/hello-world"
	"backend-boilerplate/src/application/api/presenters"
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title           Backend BoilerPlate
// @version         1.0
// @description     A complete boilerplate in golang using gin framework

// @contact.name   Augusto Timm
// @contact.url    https://github.com/augustotimm
// @contact.email  augusto.timm@gmail.com

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:1337
// @BasePath  /api/v1

func Setup(ginApp *gin.Engine, db *sql.DB) {
	jsonPresenter := &presenters.JsonPresenter{}
	baseApiAddress := "/api/v1"
	healthCheckController.Handler(ginApp, jsonPresenter, baseApiAddress)
	helloWorldController.Handler(ginApp, db, jsonPresenter)
	ginApp.GET("/api/v1/docs/swagger", ginSwagger.WrapHandler(swaggerFiles.Handler))

	ginApp.Run()
	fmt.Println("API OK!")
}
