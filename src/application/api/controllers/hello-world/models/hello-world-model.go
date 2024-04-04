package models

import (
	"backend-boilerplate/src/core/orm/models"
)

type HelloWorldPresenter struct {
	ID   string `json:"id" example:"ea68c913-2e60-487d-b108-26836271e500"`
	Name string `json:"name" example:"jon doew"`
}

func (presenter *HelloWorldPresenter) InitFromModel(entity models.HelloWorldEntity) {
	presenter.ID = entity.ID
	presenter.Name = entity.Name.String
}

func InitListFromModelSlice(modelSlice models.HelloWorldEntitySlice) []HelloWorldPresenter {
	presenterSlice := make([]HelloWorldPresenter, len(modelSlice))
	for _, entity := range modelSlice {
		var presenter HelloWorldPresenter
		presenter.InitFromModel(*entity)
		presenterSlice = append(presenterSlice, presenter)
	}
	return presenterSlice
}
