// GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag

package docs

import (
	"bytes"
	"encoding/json"
	"strings"

	"github.com/alecthomas/template"
	"github.com/swaggo/swag"
)

var doc = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{.Description}}",
        "title": "{{.Title}}",
        "contact": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/book": {
            "post": {
                "description": "本の更新",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "本の更新",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/entity.Book"
                        }
                    }
                }
            }
        },
        "/book/detail/{id}": {
            "get": {
                "description": "本の詳細情報の取得",
                "consumes": [
                    "application/json"
                ],
                "summary": "本の詳細情報の取得",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/outputport.BookDetail"
                        }
                    }
                }
            }
        },
        "/book/{id}": {
            "get": {
                "description": "idによる本の取得",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "idによる本の取得",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "本のID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/entity.Book"
                        }
                    }
                }
            },
            "delete": {
                "description": "本の削除",
                "consumes": [
                    "application/json"
                ],
                "summary": "本の削除",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "202": {
                        "description": ""
                    }
                }
            }
        },
        "/comment": {
            "post": {
                "description": "コメントの登録",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "コメントの登録",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/outputport.Comment"
                        }
                    }
                }
            },
            "patch": {
                "description": "コメントの更新",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "コメントの更新",
                "responses": {
                    "200": {
                        "description": ""
                    }
                }
            }
        },
        "/comment/{id}": {
            "get": {
                "description": "idによるコメントの取得",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "idによるコメントの取得",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "コメントのID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/outputport.Comment"
                        }
                    }
                }
            },
            "delete": {
                "description": "コメントの削除",
                "consumes": [
                    "application/json"
                ],
                "summary": "コメントの削除",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "202": {
                        "description": ""
                    }
                }
            }
        },
        "/review": {
            "post": {
                "description": "レビューの登録",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "レビューの登録",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/outputport.Review"
                        }
                    }
                }
            },
            "patch": {
                "description": "レビューの更新",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "レビューの更新",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/entity.Review"
                        }
                    }
                }
            }
        },
        "/review/detail/{id}": {
            "get": {
                "description": "レビュー詳細情報の取得",
                "consumes": [
                    "application/json"
                ],
                "summary": "レビュー詳細情報の取得",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/outputport.ReviewDetail"
                        }
                    }
                }
            }
        },
        "/review/{id}": {
            "get": {
                "description": "idによるレビューの取得",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "idによるレビューの取得",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "レビューのID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/outputport.Review"
                        }
                    }
                }
            },
            "delete": {
                "description": "レビューの削除",
                "consumes": [
                    "application/json"
                ],
                "summary": "レビューの削除",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "202": {
                        "description": ""
                    }
                }
            }
        },
        "/user/detail/{id}": {
            "get": {
                "description": "ユーザー情報の取得",
                "consumes": [
                    "application/json"
                ],
                "summary": "ユーザー情報の取得",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/outputport.User"
                        }
                    }
                }
            }
        },
        "/users": {
            "get": {
                "description": "ユーザー情報の全員取得",
                "consumes": [
                    "application/json"
                ],
                "summary": "ユーザー情報の全員取得",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/outputport.User"
                            }
                        }
                    }
                }
            }
        },
        "/users/report": {
            "get": {
                "description": "ユーザー情報全員のスプレッドシート出力",
                "consumes": [
                    "application/json"
                ],
                "summary": "ユーザー情報全員のスプレッドシート出力",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/version": {
            "get": {
                "summary": "バージョン情報を文字列で返す",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "entity.Book": {
            "type": "object",
            "properties": {
                "author": {
                    "type": "string"
                },
                "createDate": {
                    "type": "string"
                },
                "createUserID": {
                    "type": "integer"
                },
                "dateOfIssue": {
                    "type": "string"
                },
                "deleteDate": {
                    "type": "string"
                },
                "deleteFlag": {
                    "type": "integer"
                },
                "deleteUserID": {
                    "type": "integer"
                },
                "id": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                },
                "price": {
                    "type": "number"
                },
                "publisher": {
                    "type": "string"
                },
                "updateDate": {
                    "type": "string"
                },
                "updateUserID": {
                    "type": "integer"
                }
            }
        },
        "entity.Review": {
            "type": "object",
            "properties": {
                "bookID": {
                    "type": "integer"
                },
                "content": {
                    "type": "string"
                },
                "createDate": {
                    "type": "string"
                },
                "createUserID": {
                    "type": "integer"
                },
                "deleteDate": {
                    "type": "string"
                },
                "deleteFlag": {
                    "type": "integer"
                },
                "deleteUserID": {
                    "type": "integer"
                },
                "id": {
                    "type": "integer"
                },
                "readingDate": {
                    "type": "string"
                },
                "title": {
                    "type": "string"
                },
                "updateDate": {
                    "type": "string"
                },
                "updateUserID": {
                    "type": "integer"
                },
                "userID": {
                    "type": "integer"
                },
                "userName": {
                    "type": "string"
                }
            }
        },
        "outputport.BookDetail": {
            "type": "object",
            "properties": {
                "author": {
                    "type": "string"
                },
                "date_of_issue": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                },
                "price": {
                    "type": "number"
                },
                "publisher": {
                    "type": "string"
                },
                "reviews": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/outputport.Review"
                    }
                }
            }
        },
        "outputport.Comment": {
            "type": "object",
            "properties": {
                "content": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "review_id": {
                    "type": "integer"
                },
                "user_id": {
                    "type": "integer"
                }
            }
        },
        "outputport.Review": {
            "type": "object",
            "properties": {
                "content": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "reading_date": {
                    "type": "string"
                },
                "title": {
                    "type": "string"
                },
                "user": {
                    "type": "string"
                }
            }
        },
        "outputport.ReviewDetail": {
            "type": "object",
            "properties": {
                "comments": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/outputport.Comment"
                    }
                },
                "content": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "reading_date": {
                    "type": "string"
                },
                "title": {
                    "type": "string"
                },
                "user": {
                    "type": "string"
                }
            }
        },
        "outputport.User": {
            "type": "object",
            "properties": {
                "address": {
                    "type": "string"
                },
                "birth_date": {
                    "type": "string"
                },
                "email": {
                    "type": "string"
                },
                "gender": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                },
                "occupation": {
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
	Version:     "0.0.0",
	Host:        "",
	BasePath:    "/v1",
	Schemes:     []string{},
	Title:       "bookshelf-api",
	Description: "読書●ーターを参考にした下位互換API",
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
	swag.Register(swag.Name, &s{})
}
