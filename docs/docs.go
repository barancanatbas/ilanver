// Package docs GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag
package docs

import (
	"bytes"
	"encoding/json"
	"strings"
	"text/template"

	"github.com/swaggo/swag"
)

var doc = `{
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
        "/Login": {
            "post": {
                "description": "Üyelerin giriş yapmasını sağlar",
                "tags": [
                    "user"
                ],
                "summary": "Login",
                "parameters": [
                    {
                        "description": " ",
                        "name": "body",
                        "in": "body",
                        "schema": {
                            "$ref": "#/definitions/request.UserLogin"
                        }
                    }
                ]
            }
        },
        "/category": {
            "put": {
                "description": "var olan kategoriyi günceller",
                "tags": [
                    "category"
                ],
                "summary": "update Category",
                "parameters": [
                    {
                        "description": " ",
                        "name": "body",
                        "in": "body",
                        "schema": {
                            "$ref": "#/definitions/request.CategoryUpdate"
                        }
                    }
                ]
            },
            "post": {
                "description": "yeni kategory ekler",
                "tags": [
                    "category"
                ],
                "summary": "Insert Category",
                "parameters": [
                    {
                        "description": " ",
                        "name": "body",
                        "in": "body",
                        "schema": {
                            "$ref": "#/definitions/request.CategoryInsert"
                        }
                    }
                ]
            },
            "delete": {
                "description": "Var olan kategori bilgilerini siler.",
                "tags": [
                    "category"
                ],
                "summary": "Delete Category",
                "parameters": [
                    {
                        "description": " ",
                        "name": "body",
                        "in": "body",
                        "schema": {
                            "$ref": "#/definitions/request.CategoryDelete"
                        }
                    }
                ]
            }
        },
        "/category/main": {
            "get": {
                "description": "root kategorileri getirir.",
                "tags": [
                    "category"
                ],
                "summary": "Main Categories"
            }
        },
        "/register": {
            "post": {
                "description": "Üyelerin kayıt yapmasını sağlar adres bilgisini kayıt eder, user detay bilgilerini kayıt eder.",
                "tags": [
                    "user"
                ],
                "summary": "Login",
                "parameters": [
                    {
                        "description": " ",
                        "name": "body",
                        "in": "body",
                        "schema": {
                            "$ref": "#/definitions/request.UserRegister"
                        }
                    }
                ]
            }
        }
    },
    "definitions": {
        "request.CategoryDelete": {
            "type": "object",
            "required": [
                "id"
            ],
            "properties": {
                "id": {
                    "type": "integer"
                }
            }
        },
        "request.CategoryInsert": {
            "type": "object",
            "required": [
                "category_name"
            ],
            "properties": {
                "category_name": {
                    "type": "string"
                },
                "main_category_id": {
                    "type": "integer"
                }
            }
        },
        "request.CategoryUpdate": {
            "type": "object",
            "required": [
                "category_name",
                "id"
            ],
            "properties": {
                "category_name": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "main_category_id": {
                    "type": "integer"
                }
            }
        },
        "request.UserLogin": {
            "type": "object",
            "required": [
                "password",
                "phone"
            ],
            "properties": {
                "password": {
                    "type": "string"
                },
                "phone": {
                    "type": "string"
                }
            }
        },
        "request.UserRegister": {
            "type": "object",
            "required": [
                "birthday",
                "description",
                "districtfk",
                "email",
                "name",
                "password",
                "phone",
                "surname"
            ],
            "properties": {
                "birthday": {
                    "type": "string"
                },
                "description": {
                    "type": "string"
                },
                "districtfk": {
                    "type": "integer"
                },
                "email": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                },
                "phone": {
                    "type": "string"
                },
                "surname": {
                    "type": "string"
                }
            }
        }
    }
}`

type swaggerInfo struct {
	Version     string
	Host        string
	BasePath    string
	Schemes     []string
	Title       string
	Description string
}

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = swaggerInfo{
	Version:     "1.0",
	Host:        "petstore.sw",
	BasePath:    "",
	Schemes:     []string{},
	Title:       "Swagger Example API",
	Description: "This is a sample server Petstore server.",
}

type s struct{}

func (s *s) ReadDoc() string {
	sInfo := SwaggerInfo
	sInfo.Description = strings.Replace(sInfo.Description, "\n", "\\n", -1)

	t, err := template.New("swagger_info").Funcs(template.FuncMap{
		"marshal": func(v interface{}) string {
			a, _ := json.Marshal(v)
			return string(a)
		},
		"escape": func(v interface{}) string {
			// escape tabs
			str := strings.Replace(v.(string), "\t", "\\t", -1)
			// replace " with \", and if that results in \\", replace that with \\\"
			str = strings.Replace(str, "\"", "\\\"", -1)
			return strings.Replace(str, "\\\\\"", "\\\\\\\"", -1)
		},
	}).Parse(doc)
	if err != nil {
		return doc
	}

	var tpl bytes.Buffer
	if err := t.Execute(&tpl, sInfo); err != nil {
		return doc
	}

	return tpl.String()
}

func init() {
	swag.Register("swagger", &s{})
}
