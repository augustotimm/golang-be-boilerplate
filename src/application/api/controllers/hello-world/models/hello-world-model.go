package models

import (
	"encoding/json"
	"go-be-boilerplate/src/application/api/presenters"
	"go-be-boilerplate/src/core/orm/models"
)

type HelloWorldPresenter struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

func (presenter *HelloWorldPresenter) Encode() []byte {
	encoded, err := json.Marshal(presenter)
	if err != nil {
		return nil
	}
	return encoded
}

func (presenter *HelloWorldPresenter) InitFromModel(entity models.HelloWorldEntity) {
	presenter.ID = entity.ID
	presenter.Name = entity.Name.String
}

func (presenter *HelloWorldPresenter) CastToApiReturnModel() presenters.ApiPresenter {

	return presenter
}

func CastModelListToPresenter(entityList []*models.HelloWorldEntity) []presenters.ApiPresenter {
	presenterList := make([]presenters.ApiPresenter, len(entityList))

	for _, entity := range entityList {
		newPresenter := new(HelloWorldPresenter)

		newPresenter.InitFromModel(*entity)
		presenterList = append(presenterList, newPresenter)
	}

	return presenterList
}
