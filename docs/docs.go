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
        "/books": {
            "post": {
                "description": "Create a new book with the provided details",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "books"
                ],
                "summary": "Create a new book",
                "parameters": [
                    {
                        "description": "Book details",
                        "name": "book",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/book.CreateBookRequest"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/book.CreateBookResponse"
                        }
                    }
                }
            }
        },
        "/books/{id}": {
            "get": {
                "description": "Get a book by the provided ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "books"
                ],
                "summary": "Get a book by ID",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Book ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/book.GetBookResponse"
                        }
                    }
                }
            },
            "put": {
                "description": "Update a book with the provided details",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "books"
                ],
                "summary": "Update a book by ID",
                "parameters": [
                    {
                        "type": "string",
                        "format": "uuid",
                        "description": "Book ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Book details",
                        "name": "book",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/book.UpdateBookRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/book.CreateBookResponse"
                        }
                    }
                }
            },
            "delete": {
                "description": "Delete a book by the provided ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "books"
                ],
                "summary": "Delete a book by ID",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Book ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "book.CreateBookRequest": {
            "type": "object",
            "required": [
                "author",
                "isbn",
                "publishedYear",
                "title"
            ],
            "properties": {
                "author": {
                    "type": "string",
                    "example": "F. Scott Fitzgerald"
                },
                "isbn": {
                    "type": "string",
                    "example": "9780743273565"
                },
                "publishedYear": {
                    "type": "integer",
                    "example": 1925
                },
                "title": {
                    "type": "string",
                    "example": "The Great Gatsby"
                }
            }
        },
        "book.CreateBookResponse": {
            "type": "object",
            "properties": {
                "author": {
                    "type": "string",
                    "example": "F. Scott Fitzgerald"
                },
                "createdAt": {
                    "type": "string",
                    "example": "2024-01-01T00:00:00Z"
                },
                "id": {
                    "type": "string",
                    "format": "uuid",
                    "example": "123e4567-e89b-12d3-a456-426614174000"
                },
                "isbn": {
                    "type": "string",
                    "example": "9780743273565"
                },
                "publishedYear": {
                    "type": "integer",
                    "example": 1925
                },
                "title": {
                    "type": "string",
                    "example": "The Great Gatsby"
                }
            }
        },
        "book.GetBookResponse": {
            "type": "object",
            "properties": {
                "author": {
                    "type": "string",
                    "example": "F. Scott Fitzgerald"
                },
                "created_at": {
                    "type": "string",
                    "example": "2024-01-01T00:00:00Z"
                },
                "id": {
                    "type": "string",
                    "format": "uuid",
                    "example": "123e4567-e89b-12d3-a456-426614174000"
                },
                "isbn": {
                    "type": "string",
                    "example": "9780743273565"
                },
                "published_year": {
                    "type": "integer",
                    "example": 1925
                },
                "title": {
                    "type": "string",
                    "example": "The Great Gatsby"
                },
                "updated_at": {
                    "type": "string",
                    "example": "2024-01-01T00:00:00Z"
                }
            }
        },
        "book.UpdateBookRequest": {
            "type": "object",
            "required": [
                "author",
                "isbn",
                "published_year",
                "title"
            ],
            "properties": {
                "author": {
                    "type": "string",
                    "example": "F. Scott Fitzgerald"
                },
                "isbn": {
                    "type": "string",
                    "example": "9780743273565"
                },
                "published_year": {
                    "type": "integer",
                    "example": 1925
                },
                "title": {
                    "type": "string",
                    "example": "The Great Gatsby"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:8080",
	BasePath:         "/v1/api",
	Schemes:          []string{"http"},
	Title:            "Book Review API",
	Description:      "This is a sample server Book Review server.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
