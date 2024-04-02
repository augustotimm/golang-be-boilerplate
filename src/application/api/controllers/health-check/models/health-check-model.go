package healthCheckController

import (
	"encoding/json"
	_ "go-be-boilerplate/src/application/api/presenters"
)

type HealthCheck struct {
	Message string `json:"message"`
}

func (entity *HealthCheck) Encode() []byte {
	encoded, err := json.Marshal(entity)
	if err != nil {
		return nil
	}
	return encoded
}
