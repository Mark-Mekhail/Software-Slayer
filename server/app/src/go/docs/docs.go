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
        "/login": {
            "post": {
                "description": "Login with an email and password",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Users"
                ],
                "summary": "Login",
                "parameters": [
                    {
                        "description": "Credentials object that needs to be added",
                        "name": "credentials",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/user.Credentials"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/user.LoginResponse"
                        }
                    }
                }
            }
        },
        "/skill": {
            "put": {
                "description": "Update a skill for a user",
                "consumes": [
                    "application/json"
                ],
                "tags": [
                    "Skills"
                ],
                "summary": "Update a skill",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Bearer token",
                        "name": "Authorization",
                        "in": "header"
                    },
                    {
                        "description": "Old and new skill topics",
                        "name": "skill",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/skills.UpdateSkillRequest"
                        }
                    }
                ],
                "responses": {
                    "204": {
                        "description": "No Content"
                    }
                }
            },
            "post": {
                "description": "Add a new skill for a user",
                "consumes": [
                    "application/json"
                ],
                "tags": [
                    "Skills"
                ],
                "summary": "Create a new skill",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Bearer token",
                        "name": "Authorization",
                        "in": "header"
                    },
                    {
                        "description": "Skill topic to add",
                        "name": "skill",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/skills.CreateSkillRequest"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created"
                    }
                }
            }
        },
        "/skill/:topic": {
            "delete": {
                "description": "Delete a skill for a user",
                "consumes": [
                    "application/json"
                ],
                "tags": [
                    "Skills"
                ],
                "summary": "Delete a skill",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Bearer token",
                        "name": "Authorization",
                        "in": "header"
                    },
                    {
                        "type": "string",
                        "description": "Skill topic to delete",
                        "name": "topic",
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
        },
        "/skill/{user_id}": {
            "get": {
                "description": "Get all skills for a user",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Skills"
                ],
                "summary": "Get skills by user id",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "User ID",
                        "name": "user_id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "type": "string"
                            }
                        }
                    }
                }
            }
        },
        "/user": {
            "get": {
                "description": "Get a filtered set of users",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Users"
                ],
                "summary": "Get users",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Bearer token",
                        "name": "Authorization",
                        "in": "header"
                    },
                    {
                        "type": "boolean",
                        "description": "Get only the current user",
                        "name": "current",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/user.GetUserResponse"
                            }
                        }
                    }
                }
            },
            "post": {
                "description": "Register a new user with an email, password, and name",
                "consumes": [
                    "application/json"
                ],
                "tags": [
                    "Users"
                ],
                "summary": "Create a new user",
                "parameters": [
                    {
                        "description": "User object that needs to be added",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/user.CreateUserRequest"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created"
                    }
                }
            }
        }
    },
    "definitions": {
        "skills.CreateSkillRequest": {
            "type": "object",
            "properties": {
                "topic": {
                    "type": "string"
                }
            }
        },
        "skills.UpdateSkillRequest": {
            "type": "object",
            "properties": {
                "oldTopic": {
                    "type": "string"
                },
                "updatedTopic": {
                    "type": "string"
                }
            }
        },
        "user.CreateUserRequest": {
            "type": "object",
            "properties": {
                "email": {
                    "type": "string"
                },
                "first_name": {
                    "type": "string"
                },
                "last_name": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                },
                "username": {
                    "type": "string"
                }
            }
        },
        "user.Credentials": {
            "type": "object",
            "properties": {
                "identifier": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                }
            }
        },
        "user.GetCurrentUserResponse": {
            "type": "object",
            "properties": {
                "email": {
                    "type": "string"
                },
                "first_name": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "last_name": {
                    "type": "string"
                },
                "username": {
                    "type": "string"
                }
            }
        },
        "user.GetUserResponse": {
            "type": "object",
            "properties": {
                "first_name": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "last_name": {
                    "type": "string"
                },
                "username": {
                    "type": "string"
                }
            }
        },
        "user.LoginResponse": {
            "type": "object",
            "properties": {
                "token": {
                    "type": "string"
                },
                "user_info": {
                    "$ref": "#/definitions/user.GetCurrentUserResponse"
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
