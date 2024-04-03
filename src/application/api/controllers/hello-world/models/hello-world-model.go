package models

import (
	"backend-boilerplate/src/core/orm/models"
)

type HelloWorldPresenter struct {
	ID   string `json:"id"`
	Name string `json:"name"`
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
