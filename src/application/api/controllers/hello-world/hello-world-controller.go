package helloWorldController

import (
	helloWorldModel "backend-boilerplate/src/application/api/controllers/hello-world/models"
	"backend-boilerplate/src/application/api/presenters"
	"backend-boilerplate/src/core/orm/models"
	"context"
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type HelloWorldController struct {
	DB            *sql.DB
	JsonPresenter presenters.JsonPresenter
}

// List Hello World backend-boilerplate
// Lists data structure from database
// @Summary Simple list data example
// @Produce json
// @Success 200 {array} helloWorldModel.HelloWorldPresenter
// @Failure 404 {string} string
// @Router /hello-world [get]
func (hw HelloWorldController) List(ctx *gin.Context) {
	helloWorldExample, queryErr := models.HelloWorldEntities().All(context.Background(), hw.DB)

	if queryErr != nil {
		fmt.Println(queryErr)
		ctx.Status(http.StatusInternalServerError)
	}

	presenterList := helloWorldModel.InitListFromModelSlice(helloWorldExample)

	response := hw.JsonPresenter.Envelope(presenterList)

	ctx.JSON(200, response)
}

func Handler(
	ginApp *gin.Engine,
	db *sql.DB,
	jsonPresenter *presenters.JsonPresenter,
	baseApiAddress string,
) {
	helloWorldController := HelloWorldController{
		DB:            db,
		JsonPresenter: *jsonPresenter,
	}

	apiAddress := fmt.Sprintf("%s/health-check", baseApiAddress)
	ginApp.GET(apiAddress, helloWorldController.List)
}
