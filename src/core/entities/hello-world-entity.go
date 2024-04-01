package entities

import (
	"encoding/json"
	"go-be-boilerplate/src/application/api/presenters"
)

type HelloWorldEntity struct {
	BaseEntity
	Name string `json:"name"`
}

func (entity *HelloWorldEntity) Encode() []byte {
	encoded, err := json.Marshal(entity)
	if err != nil {
		return nil
	}
	return encoded
}

func (entity *HelloWorldEntity) CastToApiReturnModel() presenters.ApiReturnModel {

	return entity
}
