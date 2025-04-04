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
        "/": {
            "get": {
                "security": [
                    {
                        "Bearer Authenticate": []
                    }
                ],
                "description": "get all todos with specified info and Account ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "todos"
                ],
                "summary": "get all todo",
                "responses": {
                    "200": {
                        "description": "always returns status 200 but body contains error",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/todo.BaseResult"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "result": {
                                            "type": "array",
                                            "items": {
                                                "$ref": "#/definitions/models.TodoEntity"
                                            }
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            },
            "post": {
                "security": [
                    {
                        "Bearer Authenticate": []
                    }
                ],
                "description": "add new todo task with specified info and Account ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "todos"
                ],
                "summary": "add new todo task",
                "parameters": [
                    {
                        "description": "todo input",
                        "name": "input",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.TodoInput"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "always returns status 200 but body contains error",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/todo.BaseResult"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "result": {
                                            "$ref": "#/definitions/models.TodoEntity"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        },
        "/change_pass": {
            "patch": {
                "security": [
                    {
                        "Bearer Authenticate": []
                    }
                ],
                "description": "change password with specified id",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "accounts"
                ],
                "summary": "change password",
                "parameters": [
                    {
                        "description": "change password input",
                        "name": "input",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.ChangePasswordInput"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "always returns status 200 but body contains error",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/todo.BaseResult"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "result": {
                                            "type": "string"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        },
        "/signin": {
            "post": {
                "description": "sign in to existing account with specified mobile and expected info",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "accounts"
                ],
                "summary": "sign in to existing account",
                "parameters": [
                    {
                        "description": "sign in input",
                        "name": "input",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.LoginInput"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "always returns status 200 but body contains error",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/todo.BaseResult"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "result": {
                                            "$ref": "#/definitions/models.TokenBundleOutput"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        },
        "/signup": {
            "post": {
                "description": "create new account with specified mobile and expected info",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "accounts"
                ],
                "summary": "signing up a new account",
                "parameters": [
                    {
                        "description": "sign up input",
                        "name": "input",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.SignUpInput"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "always returns status 200 but body contains error",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/todo.BaseResult"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "result": {
                                            "$ref": "#/definitions/models.TokenBundleOutput"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        },
        "/status/{id}/{status}": {
            "patch": {
                "security": [
                    {
                        "Bearer Authenticate": []
                    }
                ],
                "description": "update status todo task with specified info and Account ID AND ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "todos"
                ],
                "summary": "update status todo task",
                "parameters": [
                    {
                        "type": "string",
                        "description": "todo ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "todo status",
                        "name": "status",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "always returns status 200 but body contains error",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/todo.BaseResult"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "result": {
                                            "$ref": "#/definitions/models.TodoEntity"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        },
        "/{id}": {
            "get": {
                "security": [
                    {
                        "Bearer Authenticate": []
                    }
                ],
                "description": "get one todo with specified info and Account ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "todos"
                ],
                "summary": "get one todo",
                "parameters": [
                    {
                        "type": "string",
                        "description": "todo ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "always returns status 200 but body contains error",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/todo.BaseResult"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "result": {
                                            "$ref": "#/definitions/models.TodoEntity"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            },
            "put": {
                "security": [
                    {
                        "Bearer Authenticate": []
                    }
                ],
                "description": "update todo task with specified info and Account ID AND ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "todos"
                ],
                "summary": "update todo task",
                "parameters": [
                    {
                        "description": "todo input",
                        "name": "input",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.TodoInput"
                        }
                    },
                    {
                        "type": "string",
                        "description": "todo ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "always returns status 200 but body contains error",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/todo.BaseResult"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "result": {
                                            "$ref": "#/definitions/models.TodoEntity"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            },
            "delete": {
                "security": [
                    {
                        "Bearer Authenticate": []
                    }
                ],
                "description": "delete todo with specified info and Account ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "todos"
                ],
                "summary": "delete todo",
                "parameters": [
                    {
                        "type": "string",
                        "description": "todo ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "always returns status 200 but body contains error",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/todo.BaseResult"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "result": {
                                            "type": "string"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "models.AccountEntity": {
            "type": "object",
            "properties": {
                "first_name": {
                    "description": "FirstName",
                    "type": "string"
                },
                "last_name": {
                    "description": "LastName",
                    "type": "string"
                },
                "todos": {
                    "description": "@Todos",
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/models.TodoEntity"
                    }
                },
                "user_name": {
                    "description": "Username",
                    "type": "string"
                }
            }
        },
        "models.ChangePasswordInput": {
            "type": "object",
            "properties": {
                "new": {
                    "type": "string"
                },
                "old": {
                    "type": "string"
                },
                "subject": {
                    "type": "integer"
                }
            }
        },
        "models.LoginInput": {
            "type": "object",
            "properties": {
                "password": {
                    "type": "string"
                },
                "user_name": {
                    "type": "string"
                }
            }
        },
        "models.SignUpInput": {
            "type": "object",
            "properties": {
                "first_name": {
                    "type": "string"
                },
                "last_name": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                },
                "user_name": {
                    "type": "string"
                }
            }
        },
        "models.TodoEntity": {
            "type": "object",
            "properties": {
                "account_id": {
                    "description": "AccountID",
                    "type": "integer"
                },
                "description": {
                    "description": "Description",
                    "type": "string"
                },
                "priority": {
                    "description": "Priority\t[1 = not important, 2 = important, 3 = very important]",
                    "type": "integer"
                },
                "status": {
                    "description": "Status [in_progress, complete, ]",
                    "type": "string"
                },
                "title": {
                    "description": "Title",
                    "type": "string"
                }
            }
        },
        "models.TodoInput": {
            "type": "object",
            "properties": {
                "description": {
                    "type": "string"
                },
                "priority": {
                    "type": "integer"
                },
                "status": {
                    "type": "string"
                },
                "title": {
                    "type": "string"
                }
            }
        },
        "models.TokenBundleOutput": {
            "type": "object",
            "properties": {
                "expire": {
                    "description": "Expire time of Token and CentrifugeToken",
                    "type": "string"
                },
                "refresh": {
                    "description": "Refresh token string used for refreshing authentication and give fresh token",
                    "type": "string"
                },
                "token": {
                    "description": "Token is JWT/PASETO token staring for storing in client side as access token",
                    "type": "string"
                }
            }
        },
        "todo.BaseResult": {
            "type": "object",
            "properties": {
                "errors": {
                    "description": "Errors provides list off error that occurred in processing request",
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "result": {
                    "description": "Result single/array of any type (object/number/string/boolean) that returns as response"
                },
                "result_count": {
                    "description": "ResultCount specified number of records that returned in result_count field expected result been array.",
                    "type": "integer"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "",
	Host:             "",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "",
	Description:      "",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
