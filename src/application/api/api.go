package api

import (
	"database/sql"
	"fmt"
	healthCheckController "go-be-boilerplate/src/application/api/controllers/health-check"
	helloWorldController "go-be-boilerplate/src/application/api/controllers/hello-world"
	"go-be-boilerplate/src/application/api/presenters"

	"github.com/gin-gonic/gin"
)

func Setup(ginApp *gin.Engine, db *sql.DB) {
	jsonPresenter := &presenters.JsonPresenter{}

	healthCheckController.Handler(ginApp, jsonPresenter)
	helloWorldController.Handler(ginApp, db, jsonPresenter)

	fmt.Println("API OK!")
}
