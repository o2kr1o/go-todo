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
        "/todos": {
            "get": {
                "description": "すべてのTODOを取得します",
                "produces": [
                    "application/json"
                ],
                "summary": "TODO一覧取得",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.Todo"
                            }
                        }
                    }
                }
            },
            "post": {
                "description": "新しいTODOを作成します",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "TODOを作成",
                "parameters": [
                    {
                        "description": "TODO情報",
                        "name": "todo",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.Todo"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/models.Todo"
                        }
                    }
                }
            }
        },
        "/todos/{id}": {
            "put": {
                "description": "指定したIDのTODOを更新します",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "TODOを更新",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "TODO ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "更新内容",
                        "name": "todo",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.Todo"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Todo"
                        }
                    }
                }
            },
            "delete": {
                "description": "指定したIDのTODOを削除します",
                "summary": "TODOを削除",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "TODO ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "204": {
                        "description": "No Content"
                    }
                }
            }
        }
    },
    "definitions": {
        "models.Todo": {
            "type": "object",
            "properties": {
                "completed": {
                    "type": "boolean",
                    "example": false
                },
                "id": {
                    "type": "integer",
                    "example": 1
                },
                "title": {
                    "type": "string",
                    "example": "買い物に行く"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "go-todo-hh8n.onrender.com",
	BasePath:         "/",
	Schemes:          []string{"https"},
	Title:            "TODO API",
	Description:      "シンプルなTODO管理API",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
