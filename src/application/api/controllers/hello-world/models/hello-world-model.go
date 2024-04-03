package models

import (
	"backend-boilerplate/src/core/orm/models"
	"encoding/json"
)

type HelloWorldPresenter struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

func (presenter *HelloWorldPresenter) JsonMarshal() []byte {
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
