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
            "name": "Granite Bagas",
            "email": "granitebagas28@gmail.com"
        },
        "license": {
            "name": "MIT"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/api/v1/admin/categories": {
            "post": {
                "security": [
                    {
                        "BasicAuth": []
                    }
                ],
                "description": "Create fasting category",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Admin"
                ],
                "summary": "Create category",
                "parameters": [
                    {
                        "description": "JSON payload",
                        "name": "payload",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/requests.CategoryRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "desc",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/utils.JSONResponse"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/transformer.CategoryTransformer"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/utils.JSONResponse"
                        }
                    }
                }
            }
        },
        "/api/v1/admin/categories/:id": {
            "put": {
                "security": [
                    {
                        "BasicAuth": []
                    }
                ],
                "description": "Update fasting category",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Admin"
                ],
                "summary": "Update category",
                "parameters": [
                    {
                        "description": "JSON payload",
                        "name": "payload",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/requests.CategoryRequest"
                        }
                    },
                    {
                        "type": "integer",
                        "description": "Category ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "desc",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/utils.JSONResponse"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/transformer.CategoryTransformer"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/utils.JSONResponse"
                        }
                    }
                }
            },
            "delete": {
                "security": [
                    {
                        "BasicAuth": []
                    }
                ],
                "description": "Delete fasting category",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Admin"
                ],
                "summary": "Delete category",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Category ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "desc",
                        "schema": {
                            "$ref": "#/definitions/utils.JSONResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/utils.JSONResponse"
                        }
                    }
                }
            }
        },
        "/api/v1/admin/fastings": {
            "post": {
                "security": [
                    {
                        "BasicAuth": []
                    }
                ],
                "description": "Create fasting",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Admin"
                ],
                "summary": "Create fasting",
                "parameters": [
                    {
                        "description": "JSON payload",
                        "name": "payload",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/requests.FastingCreateUpdateRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "desc",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/utils.JSONResponse"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/transformer.FastingTransformer"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/utils.JSONResponse"
                        }
                    }
                }
            }
        },
        "/api/v1/admin/fastings/:id": {
            "put": {
                "security": [
                    {
                        "BasicAuth": []
                    }
                ],
                "description": "Update fasting",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Admin"
                ],
                "summary": "Update fasting",
                "parameters": [
                    {
                        "description": "JSON payload",
                        "name": "payload",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/requests.FastingCreateUpdateRequest"
                        }
                    },
                    {
                        "type": "integer",
                        "description": "Fasting ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "desc",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/utils.JSONResponse"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/transformer.FastingTransformer"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/utils.JSONResponse"
                        }
                    }
                }
            },
            "delete": {
                "security": [
                    {
                        "BasicAuth": []
                    }
                ],
                "description": "Delete fasting",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Admin"
                ],
                "summary": "Delete fasting",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Fasting ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "desc",
                        "schema": {
                            "$ref": "#/definitions/utils.JSONResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/utils.JSONResponse"
                        }
                    }
                }
            }
        },
        "/api/v1/admin/sources": {
            "post": {
                "security": [
                    {
                        "BasicAuth": []
                    }
                ],
                "description": "Create fasting source",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Admin"
                ],
                "summary": "Create source",
                "parameters": [
                    {
                        "description": "JSON payload",
                        "name": "payload",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/requests.SourceRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "desc",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/utils.JSONResponse"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/transformer.SourceTransformer"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/utils.JSONResponse"
                        }
                    }
                }
            }
        },
        "/api/v1/admin/sources/:id": {
            "put": {
                "security": [
                    {
                        "BasicAuth": []
                    }
                ],
                "description": "Update fasting source",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Admin"
                ],
                "summary": "Update source",
                "parameters": [
                    {
                        "description": "JSON payload",
                        "name": "payload",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/requests.SourceRequest"
                        }
                    },
                    {
                        "type": "integer",
                        "description": "Source ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "desc",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/utils.JSONResponse"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/transformer.SourceTransformer"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/utils.JSONResponse"
                        }
                    }
                }
            },
            "delete": {
                "security": [
                    {
                        "BasicAuth": []
                    }
                ],
                "description": "Delete fasting source",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Admin"
                ],
                "summary": "Delete source",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Source ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "desc",
                        "schema": {
                            "$ref": "#/definitions/utils.JSONResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/utils.JSONResponse"
                        }
                    }
                }
            }
        },
        "/api/v1/admin/types": {
            "post": {
                "security": [
                    {
                        "BasicAuth": []
                    }
                ],
                "description": "Create fasting type",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Admin"
                ],
                "summary": "Create type",
                "parameters": [
                    {
                        "description": "JSON payload",
                        "name": "payload",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/requests.TypeRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "desc",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/utils.JSONResponse"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/transformer.TypeTransformer"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/utils.JSONResponse"
                        }
                    }
                }
            }
        },
        "/api/v1/admin/types/:id": {
            "put": {
                "security": [
                    {
                        "BasicAuth": []
                    }
                ],
                "description": "Update fasting type",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Admin"
                ],
                "summary": "Update type",
                "parameters": [
                    {
                        "description": "JSON payload",
                        "name": "payload",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/requests.TypeRequest"
                        }
                    },
                    {
                        "type": "integer",
                        "description": "Type ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "desc",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/utils.JSONResponse"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/transformer.TypeTransformer"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/utils.JSONResponse"
                        }
                    }
                }
            },
            "delete": {
                "security": [
                    {
                        "BasicAuth": []
                    }
                ],
                "description": "Delete fasting type",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Admin"
                ],
                "summary": "Delete type",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Type ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "desc",
                        "schema": {
                            "$ref": "#/definitions/utils.JSONResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/utils.JSONResponse"
                        }
                    }
                }
            }
        },
        "/api/v1/categories": {
            "get": {
                "description": "Get list of categories",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Categories"
                ],
                "summary": "List Categories",
                "responses": {
                    "200": {
                        "description": "desc",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/utils.JSONResponse"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "type": "array",
                                            "items": {
                                                "$ref": "#/definitions/transformer.CategoryTransformer"
                                            }
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/utils.JSONResponse"
                        }
                    }
                }
            }
        },
        "/api/v1/fastings": {
            "get": {
                "description": "Get list of sunnah fasting",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Fastings"
                ],
                "summary": "List Sunnah Fastings",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Type ID",
                        "name": "type_id",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "Day in month",
                        "name": "day",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "Month",
                        "name": "month",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "Year",
                        "name": "Year",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "desc",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/utils.JSONResponse"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "type": "array",
                                            "items": {
                                                "$ref": "#/definitions/transformer.FastingTransformer"
                                            }
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/utils.JSONResponse"
                        }
                    }
                }
            }
        },
        "/api/v1/sources": {
            "get": {
                "description": "Get list of sources",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Sources"
                ],
                "summary": "List Sources",
                "responses": {
                    "200": {
                        "description": "desc",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/utils.JSONResponse"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "type": "array",
                                            "items": {
                                                "$ref": "#/definitions/transformer.SourceTransformer"
                                            }
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/utils.JSONResponse"
                        }
                    }
                }
            }
        },
        "/api/v1/types": {
            "get": {
                "description": "Get list of types",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Types"
                ],
                "summary": "List Types",
                "responses": {
                    "200": {
                        "description": "desc",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/utils.JSONResponse"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "type": "array",
                                            "items": {
                                                "$ref": "#/definitions/transformer.TypeTransformer"
                                            }
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/utils.JSONResponse"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "requests.CategoryRequest": {
            "type": "object",
            "required": [
                "name"
            ],
            "properties": {
                "name": {
                    "type": "string"
                }
            }
        },
        "requests.FastingCreateUpdateRequest": {
            "type": "object",
            "required": [
                "category_id",
                "date",
                "day",
                "month",
                "type_id",
                "year"
            ],
            "properties": {
                "category_id": {
                    "type": "integer"
                },
                "date": {
                    "type": "string"
                },
                "day": {
                    "type": "integer",
                    "maximum": 31,
                    "minimum": 1
                },
                "month": {
                    "type": "integer",
                    "maximum": 12,
                    "minimum": 1
                },
                "type_id": {
                    "type": "integer"
                },
                "year": {
                    "type": "integer"
                }
            }
        },
        "requests.SourceRequest": {
            "type": "object",
            "required": [
                "url"
            ],
            "properties": {
                "url": {
                    "type": "string"
                }
            }
        },
        "requests.TypeRequest": {
            "type": "object",
            "required": [
                "background_color",
                "name",
                "text_color"
            ],
            "properties": {
                "background_color": {
                    "type": "string"
                },
                "description": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "text_color": {
                    "type": "string"
                }
            }
        },
        "transformer.CategoryTransformer": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "integer",
                    "example": 1
                },
                "name": {
                    "type": "string",
                    "example": "lorem impsum"
                }
            }
        },
        "transformer.FastingTransformer": {
            "type": "object",
            "properties": {
                "category": {
                    "$ref": "#/definitions/transformer.CategoryTransformer"
                },
                "category_id": {
                    "type": "integer",
                    "example": 1
                },
                "date": {
                    "type": "string",
                    "example": "2020-01-01"
                },
                "day": {
                    "type": "integer",
                    "example": 1
                },
                "human_date": {
                    "type": "string"
                },
                "id": {
                    "type": "integer",
                    "example": 1
                },
                "month": {
                    "type": "integer",
                    "example": 1
                },
                "type": {
                    "$ref": "#/definitions/transformer.TypeTransformer"
                },
                "type_id": {
                    "type": "integer",
                    "example": 1
                },
                "year": {
                    "type": "integer",
                    "example": 2020
                }
            }
        },
        "transformer.SourceTransformer": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "integer",
                    "example": 1
                },
                "url": {
                    "type": "string",
                    "example": "https://granitebps.com"
                }
            }
        },
        "transformer.TypeTransformer": {
            "type": "object",
            "properties": {
                "background_color": {
                    "type": "string",
                    "example": "#FFFFFF"
                },
                "description": {
                    "type": "string",
                    "example": "lorem ipsum"
                },
                "id": {
                    "type": "integer",
                    "example": 1
                },
                "name": {
                    "type": "string",
                    "example": "lorem ipsum"
                },
                "text_color": {
                    "type": "string",
                    "example": "#FFFFFF"
                }
            }
        },
        "utils.JSONResponse": {
            "type": "object",
            "properties": {
                "data": {},
                "message": {
                    "type": "string"
                },
                "success": {
                    "type": "boolean"
                },
                "timestamp": {
                    "type": "integer"
                }
            }
        }
    },
    "securityDefinitions": {
        "BasicAuth": {
            "type": "basic"
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "",
	BasePath:         "/",
	Schemes:          []string{},
	Title:            "Puasa Sunnah API",
	Description:      "This is a Puasa Sunnah API Docs",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
