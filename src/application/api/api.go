package api

import (
	"backend-boilerplate/docs"
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

func Setup(ginApp *gin.Engine, db *sql.DB) {
	jsonPresenter := &presenters.JsonPresenter{}
	baseApiAddress := "/api/v1"

	docs.SwaggerInfo.BasePath = baseApiAddress
	healthCheckController.Handler(ginApp, jsonPresenter, baseApiAddress)
	helloWorldController.Handler(ginApp, db, jsonPresenter, baseApiAddress)
	ginApp.GET("/api/v1/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	fmt.Println("API OK!")
}
