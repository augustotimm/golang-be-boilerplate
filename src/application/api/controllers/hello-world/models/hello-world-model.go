package models

import (
	"backend-boilerplate/src/core/orm/models"
)

type HelloWorldPresenter struct {
	ID   string `json:"id" example:"ea68c913-2e60-487d-b108-26836271e500"`
	Name string `json:"name" example:"jon doe"`
}

// HelloWorldBodyInput model info
// @Description Basic information to create a new entity
type HelloWorldBodyInput struct {
	Name string `json:"name,omitempty" validate:"required,min=1,max=40,name" example:"Jonathan Doe" binding:"required,min=1,max=40"`
}

func (presenter *HelloWorldPresenter) InitFromModel(entity models.HelloWorldEntity) {
	presenter.ID = entity.ID
	presenter.Name = entity.Name.String
}

func InitPresenterListFromModelSlice(modelSlice models.HelloWorldEntitySlice) []HelloWorldPresenter {
	presenterSlice := make([]HelloWorldPresenter, len(modelSlice))
	for _, entity := range modelSlice {
		var presenter HelloWorldPresenter
		presenter.InitFromModel(*entity)
		presenterSlice = append(presenterSlice, presenter)
	}
	return presenterSlice
}
