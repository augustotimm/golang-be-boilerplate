package helloWorldController

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"go-be-boilerplate/src/application/api/controllers/hello-world/models"
	"go-be-boilerplate/src/application/api/presenters"
)

type HelloWorldController struct {
	DB            *sql.DB
	JsonPresenter presenters.JsonPresenter
}

func (hw HelloWorldController) List(ctx *gin.Context) {
	var users []*models.HelloWorldPresenter
	// hw.DB.Find(&users)update to sqlboiler

	payload := make([]presenters.ApiPresenter, len(users))

	for _, entity := range users {
		payload = append(payload, entity.CastToApiReturnModel())
	}

	response := hw.JsonPresenter.EnvelopeList(payload)

	ctx.JSON(200, response)
}

func Handler(
	ginApp *gin.Engine,
	db *sql.DB,
	jsonPresenter *presenters.JsonPresenter,
) {
	helloWorldController := HelloWorldController{
		DB:            db,
		JsonPresenter: *jsonPresenter,
	}

	ginApp.GET("/hello-world", helloWorldController.List)
}
