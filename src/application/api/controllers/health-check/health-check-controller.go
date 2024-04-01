package healthCheckController

import (
	"github.com/gin-gonic/gin"
	models "go-be-boilerplate/src/application/api/controllers/health-check/models"
	"go-be-boilerplate/src/application/api/presenters"
)

type HealthCheckController struct {
	JsonPresenter presenters.JsonPresenter
}

func (hc HealthCheckController) GetIndex(ctx *gin.Context) {
	payload := models.HealthCheck{Message: "Server running!"}

	response := hc.JsonPresenter.Envelope(&payload)

	ctx.JSON(200, response)
}

func (hc HealthCheckController) PostWebHook(ctx *gin.Context) {
	var req *models.WebHookModel
	req.Method = "POST"
	ctx.BindJSON(req.Content)

	response := hc.JsonPresenter.Envelope(req)

	ctx.JSON(200, response)
}

func Handler(
	ginApp *gin.Engine,
	jsonPresenter *presenters.JsonPresenter,
) {
	healthCheckController := HealthCheckController{
		JsonPresenter: *jsonPresenter,
	}

	ginApp.GET("/", healthCheckController.GetIndex)
	ginApp.POST("/", healthCheckController.PostWebHook)
}
