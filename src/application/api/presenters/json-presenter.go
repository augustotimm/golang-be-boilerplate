package presenters

import models "go-be-boilerplate/src/application/api/presenters/models"

type JsonPresenter struct{}

type ApiReturnModel interface {
	Encode() []byte
}

func (jp JsonPresenter) Envelope(payload ApiReturnModel) models.JsonModel {

	return models.JsonModel{
		Payload: payload.Encode(),
	}

}

func (jp JsonPresenter) EnvelopeList(payload []ApiReturnModel) models.JsonModel {
	var encodedPayload [][]byte

	for _, entity := range payload {
		encodedPayload = append(encodedPayload, entity.Encode())
	}
	return models.JsonModel{
		Payload: encodedPayload,
	}

}
