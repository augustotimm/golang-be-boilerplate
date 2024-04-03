package models

import (
	"encoding/json"
	"go-be-boilerplate/src/application/api/presenters"
)

type HelloWorldPresenter struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

func (entity *HelloWorldPresenter) Encode() []byte {
	encoded, err := json.Marshal(entity)
	if err != nil {
		return nil
	}
	return encoded
}

func (entity *HelloWorldPresenter) CastToApiReturnModel() presenters.ApiPresenter {

	return entity
}
