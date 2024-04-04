package validators

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	_ "github.com/go-playground/validator/v10"
	"net/http"
)

type HelloWorldInput struct {
	Name string `json:"name" validate:"required,min=1,max=40"`
}

func ValidateInput(ctx *gin.Context) {
	var input HelloWorldInput
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	validate := validator.New()

	// Validate the User struct
	validationErr := validate.Struct(input)
	if validationErr != nil {
		// Validation failed, handle the error
		errors := validationErr.(validator.ValidationErrors)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": errors})

		return
	}

	ctx.Next()
}
