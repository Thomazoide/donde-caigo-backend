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
            "name": "Thomas Tellerias",
            "email": "ttellerias01@outlook.com"
        },
        "license": {
            "name": "Donde Caigo"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/auth/login": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Autenticación"
                ],
                "summary": "Inicio de sesión",
                "parameters": [
                    {
                        "description": "Credenciales de usuario",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/structs.LoginPayload"
                        }
                    }
                ],
                "responses": {
                    "202": {
                        "description": "Accepted",
                        "schema": {
                            "$ref": "#/definitions/structs.LoginResponse"
                        }
                    }
                }
            }
        },
        "/cuenta": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Usuarios"
                ],
                "summary": "Entrega todas las cuentas de usuario",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/structs.ApiResponse"
                        }
                    }
                }
            },
            "put": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Usuarios"
                ],
                "summary": "Actualiza una cuenta de usuario",
                "parameters": [
                    {
                        "description": "Usuario a actualizar",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.UserSchema"
                        }
                    }
                ],
                "responses": {
                    "202": {
                        "description": "Accepted",
                        "schema": {
                            "$ref": "#/definitions/structs.ApiResponse"
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
                "tags": [
                    "Usuarios"
                ],
                "summary": "Crea una cuenta de usuario",
                "parameters": [
                    {
                        "description": "Usuario a crear",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.UserSchema"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/structs.ApiResponse"
                        }
                    }
                }
            },
            "delete": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Usuarios"
                ],
                "summary": "Elimina una cuenta de usuario",
                "parameters": [
                    {
                        "description": "Usuario a eliminar",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.UserSchema"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/structs.ApiResponse"
                        }
                    }
                }
            }
        },
        "/cuenta/{id}": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Usuarios"
                ],
                "summary": "Entrega una cuenta de usuario por su id",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "ID del usuario",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/structs.ApiResponse"
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
                "tags": [
                    "Usuarios"
                ],
                "summary": "Actualiza la contraseña de un usuario",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "ID del usuario",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Payload para cambiar la contraseña",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/structs.ChangePasswordPayload"
                        }
                    }
                ],
                "responses": {
                    "202": {
                        "description": "Accepted",
                        "schema": {
                            "$ref": "#/definitions/structs.ApiResponse"
                        }
                    }
                }
            }
        },
        "/publicaciones": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Publicaciones"
                ],
                "summary": "Entrega todos los posts",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/structs.ApiResponse"
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
                "tags": [
                    "Publicaciones"
                ],
                "summary": "Crear un post",
                "parameters": [
                    {
                        "description": "Nuevo post",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/structs.NewPostPayload"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/structs.ApiResponse"
                        }
                    }
                }
            },
            "delete": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Publicaciones"
                ],
                "summary": "Eliminar un post",
                "parameters": [
                    {
                        "description": "Post a eliminar",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.PostSchema"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/structs.ApiResponse"
                        }
                    }
                }
            },
            "patch": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Publicaciones"
                ],
                "summary": "Agregar un like a un post",
                "parameters": [
                    {
                        "description": "Like a agregar",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/structs.LikePayload"
                        }
                    }
                ],
                "responses": {
                    "202": {
                        "description": "Accepted",
                        "schema": {
                            "$ref": "#/definitions/structs.ApiResponse"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "models.PostSchema": {
            "type": "object",
            "properties": {
                "author_id": {
                    "type": "integer"
                },
                "description": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "pics": {
                    "type": "string"
                },
                "stars": {
                    "type": "string"
                },
                "title": {
                    "type": "string"
                }
            }
        },
        "models.UserSchema": {
            "type": "object",
            "properties": {
                "age": {
                    "type": "integer"
                },
                "email": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "nombre": {
                    "type": "string"
                },
                "profileDescription": {
                    "type": "string"
                },
                "profilePicture": {
                    "type": "string"
                },
                "rut": {
                    "type": "string"
                }
            }
        },
        "structs.ApiResponse": {
            "type": "object",
            "properties": {
                "data": {},
                "error": {
                    "type": "string"
                },
                "message": {
                    "type": "string"
                },
                "statusCode": {
                    "type": "integer"
                }
            }
        },
        "structs.ChangePasswordPayload": {
            "type": "object",
            "properties": {
                "actualPassword": {
                    "type": "string"
                },
                "newPassword": {
                    "type": "string"
                }
            }
        },
        "structs.LikePayload": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "integer"
                },
                "post": {
                    "$ref": "#/definitions/models.PostSchema"
                }
            }
        },
        "structs.LoginPayload": {
            "type": "object",
            "properties": {
                "email": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                }
            }
        },
        "structs.LoginResponse": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string"
                },
                "statusCode": {
                    "type": "integer"
                },
                "token": {
                    "type": "string"
                },
                "userData": {
                    "$ref": "#/definitions/models.UserSchema"
                }
            }
        },
        "structs.NewPostPayload": {
            "type": "object",
            "properties": {
                "pics": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "post": {
                    "$ref": "#/definitions/models.PostSchema"
                }
            }
        }
    },
    "externalDocs": {
        "description": "Basado en OpenAPI 3.0",
        "url": "https://swagger.io/resources/open-api/"
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "0.1",
	Host:             "localhost:8080",
	BasePath:         "/",
	Schemes:          []string{"http"},
	Title:            "Donde Caigo APP",
	Description:      "API para la aplicacion de Donde Caigo",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
