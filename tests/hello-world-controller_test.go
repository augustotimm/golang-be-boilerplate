package tests

import (
	helloWorldController "backend-boilerplate/src/application/api/controllers/hello-world"
	helloWorldControllerModels "backend-boilerplate/src/application/api/controllers/hello-world/models"

	presenterJsonModels "backend-boilerplate/src/application/api/presenters/models"
	"backend-boilerplate/src/core/orm/models"
	"encoding/json"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/mitchellh/mapstructure"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/fatih/structs"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestHelloWorldController(t *testing.T) {
	_assert := assert.New(t)

	controller := &helloWorldController.HelloWorldController{
		DB: db,
	}

	t.Run("Should receive 200 with empty users from GET /hello-world", func(t *testing.T) {
		gin.SetMode(gin.TestMode)
		router := gin.New()
		router.GET("/hello-world", controller.List)

		expectedUsers := []interface{}([]interface{}{})

		w := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", "/hello-world", nil)
		router.ServeHTTP(w, req)
		var response presenterJsonModels.JsonModel
		json.NewDecoder(w.Body).Decode(&response)

		_assert.Equal(expectedUsers, response.Payload)
	})

	t.Run("Should receive 200 with one user from GET /hello-world", func(t *testing.T) {
		gin.SetMode(gin.TestMode)
		router := gin.New()
		router.GET("/hello-world", controller.List)

		sampleEntity := helloWorldControllerModels.HelloWorldPresenter{
			ID:   "7144fc10-6d13-4a69-859e-29d33e6d6a79",
			Name: "Example Name",
		}
		columns := getEntityColumns()
		queryRows := sqlmock.NewRows(columns).AddRow(sampleEntity.ID, sampleEntity.Name)

		mockDB.ExpectQuery("SELECT (.+)").WillReturnRows(queryRows)

		w := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", "/hello-world", nil)
		router.ServeHTTP(w, req)
		var response presenterJsonModels.JsonModel
		json.NewDecoder(w.Body).Decode(&response)

		var message []helloWorldControllerModels.HelloWorldPresenter
		mapstructure.Decode(response.Payload, &message)

		_assert.Equal(message[0], sampleEntity)
	})

}

func getEntityColumns() []string {
	columnsInterface := structs.Values(models.HelloWorldEntityColumns)

	columnsString := make([]string, len(columnsInterface))
	for i, v := range columnsInterface {
		columnsString[i] = v.(string)
	}
	return columnsString
}
