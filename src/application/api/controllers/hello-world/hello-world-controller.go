package helloWorldController

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	modelsPresenter "go-be-boilerplate/src/application/api/controllers/hello-world/models"
	"go-be-boilerplate/src/application/api/presenters"
	"go-be-boilerplate/src/core/orm/models"
	"net/http"
)

type HelloWorldController struct {
	DB            *sql.DB
	JsonPresenter presenters.JsonPresenter
}

func (hw HelloWorldController) List(ctx *gin.Context) {
	var helloWorldExample []*models.HelloWorldEntity
	helloWorldExample, queryErr := models.HelloWorldEntities().All(context.Background(), hw.DB)

	if queryErr != nil {
		fmt.Println(queryErr)
		ctx.Status(http.StatusInternalServerError)
	}

	presenterList := modelsPresenter.CastModelListToPresenter(helloWorldExample)

	response := hw.JsonPresenter.EnvelopeList(presenterList)

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
