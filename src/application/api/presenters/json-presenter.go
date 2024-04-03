package presenters

import models "backend-boilerplate/src/application/api/presenters/models"

type JsonPresenter struct{}

func (jp JsonPresenter) Envelope(payload any) models.JsonModel {

	return models.JsonModel{
		Payload: payload,
	}

}
