package middlewares

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func ValidateJsonBody[BodyType any]() func(*gin.Context) {
	return func(c *gin.Context) {
		var body BodyType

		err := c.ShouldBindJSON(&body)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.Set("jsonBody", body)

		c.Next()
	}
}

func GetJsonBody[BodyType any](c *gin.Context) BodyType {
	return c.MustGet("jsonBody").(BodyType)
}
