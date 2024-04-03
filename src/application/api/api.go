package api

import (
	healthCheckController "backend-boilerplate/src/application/api/controllers/health-check"
	helloWorldController "backend-boilerplate/src/application/api/controllers/hello-world"
	"backend-boilerplate/src/application/api/presenters"
	"database/sql"
	"fmt"

	"github.com/gin-gonic/gin"
)

func Setup(ginApp *gin.Engine, db *sql.DB) {
	jsonPresenter := &presenters.JsonPresenter{}

	healthCheckController.Handler(ginApp, jsonPresenter)
	helloWorldController.Handler(ginApp, db, jsonPresenter)

	fmt.Println("API OK!")
}
