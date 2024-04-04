package healthCheckController

import (
	_ "backend-boilerplate/src/application/api/presenters"
)

type HealthCheck struct {
	Message string `json:"message" example:"Server running!"`
}
