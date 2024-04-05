package helloWorldController

import (
	helloWorldModel "backend-boilerplate/src/application/api/controllers/hello-world/models"
	middlewares "backend-boilerplate/src/application/api/middlewares"
	"backend-boilerplate/src/application/api/presenters"
	"backend-boilerplate/src/core/orm/models"
	"context"
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/prometheus/common/log"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"net/http"
)

type HelloWorldController struct {
	DB            *sql.DB
	JsonPresenter presenters.JsonPresenter
}

// List Hello-World
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

	presenterList := helloWorldModel.InitPresenterListFromModelSlice(helloWorldExample)

	response := hw.JsonPresenter.Envelope(presenterList)

	ctx.JSON(200, response)
}

// CreateHelloWorld Hello-World
// Create new entity in the database
// @Summary Simple Create entity example
// @Produce json
// @Accept  json
// @Param hw body helloWorldModel.HelloWorldBodyInput true "Input data to create new entity"
// @Success 201 "Created"
// @Failure 400 {string} string
// @Router /hello-world [post]
func (hw HelloWorldController) CreateHelloWorld(c *gin.Context) {
	body := middlewares.GetJsonBody[helloWorldModel.HelloWorldBodyInput](c)
	newHelloWorld := models.HelloWorldEntity{
		Name: null.String{String: body.Name, Valid: true},
		ID:   uuid.New().String(),
	}
	insertErr := newHelloWorld.Insert(context.Background(), hw.DB, boil.Infer())
	if insertErr != nil {
		log.Error(insertErr)
		c.Status(http.StatusInternalServerError)
	}
	c.Status(http.StatusCreated)
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

	apiAddress := fmt.Sprintf("%s/hello-world", baseApiAddress)
	ginApp.GET(apiAddress, helloWorldController.List)
	ginApp.POST(apiAddress, middlewares.ValidateJsonBody[helloWorldModel.HelloWorldBodyInput](), helloWorldController.CreateHelloWorld)

}
