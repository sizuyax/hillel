// Package docs Code generated by swaggo/swag. DO NOT EDIT
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/create-item": {
            "post": {
                "description": "Create item",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "models.Item"
                ],
                "summary": "Create item",
                "responses": {
                    "201": {
                        "description": "Created"
                    }
                }
            }
        },
        "/delete-item": {
            "delete": {
                "description": "Delete item",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "models.Item"
                ],
                "summary": "Delete item",
                "responses": {
                    "200": {
                        "description": "OK"
                    }
                }
            }
        },
        "/items": {
            "get": {
                "description": "Get all items",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "models.Item"
                ],
                "summary": "Get items",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/handlers.Item"
                            }
                        }
                    }
                }
            }
        },
        "/item{id}": {
            "get": {
                "description": "Get item by id",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "models.Item"
                ],
                "summary": "Get item",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/handlers.Item"
                        }
                    }
                }
            }
        },
        "/update-item": {
            "put": {
                "description": "Update item",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "models.Item"
                ],
                "summary": "Update item",
                "responses": {
                    "200": {
                        "description": "OK"
                    }
                }
            }
        }
    },
    "definitions": {
        "handlers.Item": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                },
                "price": {
                    "type": "integer"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "http://swagger.io/terms/",
	BasePath:         "/",
	Schemes:          []string{},
	Title:            "Project-Auction API",
	Description:      "Hillel Project",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}