// Package docs Code generated by swaggo/swag. DO NOT EDIT
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {
            "name": "Augusto Timm",
            "url": "https://github.com/augustotimm",
            "email": "augusto.timm@gmail.com"
        },
        "license": {
            "name": "Apache 2.0",
            "url": "http://www.apache.org/licenses/LICENSE-2.0.html"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/health-check": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "summary": "Shows information if api is running",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/healthCheckController.HealthCheck"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/hello-world": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "summary": "Simple list data example",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.HelloWorldPresenter"
                            }
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            },
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Simple Create entity example",
                "parameters": [
                    {
                        "description": "Input data to create new entity",
                        "name": "hw",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.HelloWorldBodyInput"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created"
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "Object"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "healthCheckController.HealthCheck": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string",
                    "example": "Server running!"
                }
            }
        },
        "models.HelloWorldBodyInput": {
            "description": "Basic information to create a new entity",
            "type": "object",
            "required": [
                "name"
            ],
            "properties": {
                "name": {
                    "type": "string",
                    "maxLength": 40,
                    "minLength": 1,
                    "example": "Jonathan Doe"
                }
            }
        },
        "models.HelloWorldPresenter": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "string",
                    "example": "ea68c913-2e60-487d-b108-26836271e500"
                },
                "name": {
                    "type": "string",
                    "example": "jon doe"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:1337",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "Backend BoilerPlate",
	Description:      "A complete boilerplate in golang using gin framework",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
