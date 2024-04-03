package presenters

import models "go-be-boilerplate/src/application/api/presenters/models"

type JsonPresenter struct{}

type ApiPresenter interface {
	Encode() []byte
}

func (jp JsonPresenter) Envelope(payload ApiPresenter) models.JsonModel {

	return models.JsonModel{
		Payload: payload.Encode(),
	}

}

func (jp JsonPresenter) EnvelopeList(payload []ApiPresenter) models.JsonModel {
	var encodedPayload [][]byte

	for _, entity := range payload {
		encodedPayload = append(encodedPayload, entity.Encode())
	}
	return models.JsonModel{
		Payload: encodedPayload,
	}

}
