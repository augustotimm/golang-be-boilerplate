package tests

import (
	healthCheckController "backend-boilerplate/src/application/api/controllers/health-check"
	healthCheckControllerModels "backend-boilerplate/src/application/api/controllers/health-check/models"
	"github.com/mitchellh/mapstructure"

	presenterJsonModels "backend-boilerplate/src/application/api/presenters/models"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestHealthCheckController(t *testing.T) {
	_assert := assert.New(t)

	controller := &healthCheckController.HealthCheckController{}

	t.Run("Should receive 200 from GET /", func(t *testing.T) {
		gin.SetMode(gin.TestMode)
		router := gin.New()
		router.GET("/", controller.GetIndex)

		w := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", "/", nil)
		router.ServeHTTP(w, req)
		var response presenterJsonModels.JsonModel
		json.NewDecoder(w.Body).Decode(&response)
		expectedMessage := healthCheckControllerModels.HealthCheck{Message: "Server running!"}
		var message healthCheckControllerModels.HealthCheck
		mapstructure.Decode(response.Payload, &message)

		_assert.Equal(expectedMessage, message)
	})

}
