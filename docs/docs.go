// Package docs GENERATED BY SWAG; DO NOT EDIT
// This file was generated by swaggo/swag
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "termsOfService": "http://swagger.io/terms/",
        "contact": {
            "name": "API Support",
            "url": "http://www.swagger.io/support",
            "email": "support@swagger.io"
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
        "/ext/register": {
            "post": {
                "description": "get uid",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "extension"
                ],
                "parameters": [
                    {
                        "maxLength": 64,
                        "minLength": 1,
                        "type": "string",
                        "description": "client name",
                        "name": "name",
                        "in": "query",
                        "required": true
                    },
                    {
                        "maxLength": 256,
                        "minLength": 1,
                        "type": "string",
                        "description": "extension id",
                        "name": "extension",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/ExtensionResponse"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "allOf": [
                                                {
                                                    "type": "object"
                                                },
                                                {
                                                    "type": "object",
                                                    "properties": {
                                                        "uid": {
                                                            "type": "string"
                                                        }
                                                    }
                                                }
                                            ]
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        },
        "/ext/{uid}/group": {
            "get": {
                "description": "get client's all groups and tabs",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "pull"
                ],
                "parameters": [
                    {
                        "type": "string",
                        "description": "client uid",
                        "name": "uid",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/ExtensionResponse"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/pb.Client"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            },
            "post": {
                "description": "store one or more tab group",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "add"
                ],
                "parameters": [
                    {
                        "type": "string",
                        "description": "client uid",
                        "name": "uid",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "groups",
                        "name": "groups",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/pb.Group"
                            }
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/ExtensionResponse"
                        }
                    }
                }
            }
        },
        "/ext/{uid}/swap/tab": {
            "post": {
                "description": "swap 2 tab",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "modify"
                ],
                "parameters": [
                    {
                        "type": "string",
                        "description": "client uid",
                        "name": "uid",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "group(a) uid",
                        "name": "groupA",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "group(b) uid",
                        "name": "groupB",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "group(a) tab(a) uid",
                        "name": "tabA",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "group(b) tab(b) uid",
                        "name": "tabB",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/ExtensionResponse"
                        }
                    }
                }
            }
        },
        "/ext/{uid}/{group}/": {
            "delete": {
                "description": "remove tab group and all tabs",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "remove"
                ],
                "parameters": [
                    {
                        "type": "string",
                        "description": "client uid",
                        "name": "uid",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "group uid",
                        "name": "group",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/ExtensionResponse"
                        }
                    }
                }
            }
        },
        "/ext/{uid}/{group}/{tab}": {
            "delete": {
                "description": "remove tab",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "remove"
                ],
                "parameters": [
                    {
                        "type": "string",
                        "description": "client uid",
                        "name": "uid",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "group uid",
                        "name": "group",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "tab uid",
                        "name": "tab",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/ExtensionResponse"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "ExtensionResponse": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer"
                },
                "data": {
                    "type": "any"
                },
                "error": {
                    "type": "string"
                }
            }
        },
        "pb.Client": {
            "type": "object",
            "properties": {
                "ext_id": {
                    "type": "string"
                },
                "groups": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/pb.Group"
                    }
                },
                "last_access_time": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                },
                "uid": {
                    "type": "string"
                }
            }
        },
        "pb.Group": {
            "type": "object",
            "properties": {
                "index": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                },
                "option": {
                    "$ref": "#/definitions/pb.GroupOption"
                },
                "tabs": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/pb.Tab"
                    }
                },
                "uid": {
                    "type": "string"
                }
            }
        },
        "pb.GroupOption": {
            "type": "object",
            "properties": {
                "tags": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                }
            }
        },
        "pb.Tab": {
            "type": "object",
            "properties": {
                "favicon": {
                    "type": "string"
                },
                "index": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                },
                "uid": {
                    "type": "string"
                },
                "url": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:8080",
	BasePath:         "/v1",
	Schemes:          []string{},
	Title:            "Swagger Example API",
	Description:      "This is a sample server celler server.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
