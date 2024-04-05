package tests

import (
	helloWorldControllerModels "backend-boilerplate/src/application/api/controllers/hello-world/models"
	"backend-boilerplate/src/application/api/middlewares"
	"backend-boilerplate/src/core/orm/models"
	"bytes"
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/fatih/structs"
	"github.com/gin-gonic/gin"
)

func TestValidatorMiddleware(t *testing.T) {
	_assert := assert.New(t)

	t.Run("should accept valid name", func(t *testing.T) {
		gin.SetMode(gin.TestMode)
		w := httptest.NewRecorder()
		input := helloWorldControllerModels.HelloWorldBodyInput{Name: "just a name"}

		jsonInput, jsonErr := json.Marshal(input)
		if jsonErr != nil {
			t.Fail()
		}

		gin.SetMode(gin.TestMode)
		router := gin.New()
		router.POST("/TestMiddleWare", middlewares.ValidateJsonBody[helloWorldControllerModels.HelloWorldBodyInput]())

		req, _ := http.NewRequest("POST", "/TestMiddleWare", bytes.NewReader(jsonInput))
		router.ServeHTTP(w, req)

		_assert.Equal(http.StatusOK, w.Code)
	})

	t.Run("should return 400 when HelloWorldBodyInput is invalid", func(t *testing.T) {
		gin.SetMode(gin.TestMode)
		w := httptest.NewRecorder()
		input := helloWorldControllerModels.HelloWorldBodyInput{Name: ""}

		jsonInput, jsonErr := json.Marshal(input)
		if jsonErr != nil {
			t.Fail()
		}

		gin.SetMode(gin.TestMode)
		router := gin.New()
		router.POST("/TestMiddleWare", middlewares.ValidateJsonBody[helloWorldControllerModels.HelloWorldBodyInput]())

		req, _ := http.NewRequest("POST", "/TestMiddleWare", bytes.NewReader(jsonInput))
		router.ServeHTTP(w, req)

		var response map[string]interface{}
		json.NewDecoder(w.Body).Decode(&response)

		_assert.Equal(http.StatusBadRequest, w.Code)
		_assert.GreaterOrEqual(1, len(response))
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
