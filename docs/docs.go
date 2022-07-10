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
                                            "$ref": "#/definitions/ent.ExtensionClient"
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
                                "$ref": "#/definitions/entity.GroupInfo"
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
        "ent.ExtensionClient": {
            "type": "object",
            "properties": {
                "client_uid": {
                    "description": "ClientUID holds the value of the \"client_uid\" field.",
                    "type": "string"
                },
                "edges": {
                    "description": "Edges holds the relations/edges for other nodes in the graph.\nThe values are being populated by the ExtensionClientQuery when eager-loading is set.",
                    "$ref": "#/definitions/ent.ExtensionClientEdges"
                },
                "extension_id": {
                    "description": "ExtensionID holds the value of the \"extension_id\" field.",
                    "type": "string"
                },
                "id": {
                    "description": "ID of the ent.",
                    "type": "integer"
                },
                "last_access_time": {
                    "description": "LastAccessTime holds the value of the \"last_access_time\" field.",
                    "type": "string"
                },
                "name": {
                    "description": "Name holds the value of the \"name\" field.",
                    "type": "string"
                }
            }
        },
        "ent.ExtensionClientEdges": {
            "type": "object",
            "properties": {
                "groups": {
                    "description": "Groups holds the value of the groups edge.",
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/ent.Group"
                    }
                }
            }
        },
        "ent.Group": {
            "type": "object",
            "properties": {
                "created_at": {
                    "description": "CreatedAt holds the value of the \"created_at\" field.",
                    "type": "string"
                },
                "edges": {
                    "description": "Edges holds the relations/edges for other nodes in the graph.\nThe values are being populated by the GroupQuery when eager-loading is set.",
                    "$ref": "#/definitions/ent.GroupEdges"
                },
                "id": {
                    "description": "ID of the ent.",
                    "type": "integer"
                },
                "name": {
                    "description": "Name holds the value of the \"name\" field.",
                    "type": "string"
                },
                "option": {
                    "description": "Option holds the value of the \"option\" field.",
                    "$ref": "#/definitions/entity.GroupOption"
                },
                "share_url": {
                    "description": "ShareURL holds the value of the \"share_url\" field.",
                    "type": "string"
                },
                "uid": {
                    "description": "UID holds the value of the \"uid\" field.",
                    "type": "string"
                }
            }
        },
        "ent.GroupEdges": {
            "type": "object",
            "properties": {
                "tabs": {
                    "description": "Tabs holds the value of the tabs edge.",
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/ent.Tab"
                    }
                }
            }
        },
        "ent.Tab": {
            "type": "object",
            "properties": {
                "favicon": {
                    "description": "Favicon holds the value of the \"favicon\" field.",
                    "type": "string"
                },
                "id": {
                    "description": "ID of the ent.",
                    "type": "integer"
                },
                "name": {
                    "description": "Name holds the value of the \"name\" field.",
                    "type": "string"
                },
                "seq": {
                    "description": "Seq holds the value of the \"seq\" field.\nsequence",
                    "type": "integer"
                },
                "url": {
                    "description": "URL holds the value of the \"url\" field.",
                    "type": "string"
                }
            }
        },
        "entity.GroupInfo": {
            "type": "object",
            "properties": {
                "name": {
                    "type": "string"
                },
                "option": {
                    "$ref": "#/definitions/entity.GroupOption"
                },
                "tabs": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/entity.Tab"
                    }
                }
            }
        },
        "entity.GroupOption": {
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
        "entity.Tab": {
            "type": "object",
            "properties": {
                "favicon": {
                    "type": "string"
                },
                "name": {
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
