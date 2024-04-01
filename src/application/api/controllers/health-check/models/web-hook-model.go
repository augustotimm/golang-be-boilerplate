package healthCheckController

import (
	"encoding/json"
	_ "go-be-boilerplate/src/application/api/presenters"
)

type WebHookModel struct {
	Method  string `json:"method"`
	Content any    `json:"content"`
}

func (entity *WebHookModel) Encode() []byte {
	encoded, err := json.Marshal(entity)
	if err != nil {
		return nil
	}
	return encoded
}
