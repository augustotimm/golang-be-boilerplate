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

// HealthCheck backend-boilerplate
// @Summary Shows information if api is running
// @Produce json
// @Success 200 {object} models.HealthCheck
// @Failure 404 {string} string
// @Router /health-check [get]
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
	apiAddress := fmt.Sprintf("%s/hello-world", baseApiAddress)
	ginApp.GET(apiAddress, healthCheckController.GetIndex)
}
