package healthCheckController

import (
	models "backend-boilerplate/src/application/api/controllers/health-check/models"
	"backend-boilerplate/src/application/api/presenters"
	"fmt"
	"github.com/gin-gonic/gin"
)

type HealthCheckController struct {
	JsonPresenter presenters.JsonPresenter
}

func (hc HealthCheckController) GetIndex(ctx *gin.Context) {
	payload := models.HealthCheck{Message: "Server running!"}

	response := hc.JsonPresenter.Envelope(payload)

	ctx.JSON(200, response)
}

func Handler(
	ginApp *gin.Engine,
	jsonPresenter *presenters.JsonPresenter,
	baseApiAddress string,
) {
	healthCheckController := HealthCheckController{
		JsonPresenter: *jsonPresenter,
	}
	apiAddress := fmt.Sprintf("%s/health-check", baseApiAddress)
	ginApp.GET(apiAddress, healthCheckController.GetIndex)
}
