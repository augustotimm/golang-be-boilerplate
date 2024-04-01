package helloWorldController

import (
	"go-be-boilerplate/src/application/api/presenters"
	"go-be-boilerplate/src/core/entities"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type HelloWorldController struct {
	DB            *gorm.DB
	JsonPresenter presenters.JsonPresenter
}

func (hw HelloWorldController) List(ctx *gin.Context) {
	var users []*entities.HelloWorldEntity
	hw.DB.Find(&users)

	payload := make([]presenters.ApiReturnModel, len(users))

	for _, entity := range users {
		payload = append(payload, entity.CastToApiReturnModel())
	}

	response := hw.JsonPresenter.EnvelopeList(payload)

	ctx.JSON(200, response)
}

func Handler(
	ginApp *gin.Engine,
	db *gorm.DB,
	jsonPresenter *presenters.JsonPresenter,
) {
	helloWorldController := HelloWorldController{
		DB:            db,
		JsonPresenter: *jsonPresenter,
	}

	ginApp.GET("/hello-world", helloWorldController.List)
}
