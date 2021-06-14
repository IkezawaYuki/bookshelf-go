# bookshelf-go

## Summary
読書メーターを模倣した本管理API

## Description
|  機能  |  エンドポイント  |
| ---- | ---- |
| バージョン出力| /v1/version [get] |
|  ログイン  |  /v1/auth/login [get] |
|  ログアウト  |  /v1/auth/logout [get] |
|  本の取得  |  /v1/book/{id} [get] |
|  本の全取得|  /v1/books [get] |
|  本の登録  |  /v1/book [post]  |
|  本の更新  |  /v1/book [patch]  |
|  本の削除  |  /v1/book/{id} [delete] |
|  本の詳細取得 |  /v1/book/detail/{id} [get]  |
|  ユーザーの詳細取得 |  /v1/user/detail/{id} [get]  |
|  ユーザーの前取得  |  /v1/users [get]  |
|  ユーザー情報のスプレッドシート出力 | /v1/users/report [get]|

### Layout

```tree
├── bin
│
├── docs
│   
├── src
│   ├──domain
│   │   ├── entity
│   │   ├── model
│   │   └── repository
│   │ 
│   ├──infrastructure
│   │   ├── auth
│   │   ├── http_handler
│   │   ├── mysql
│   │   ├── redis
│   │   ├── slack
│   │   ├── spreadsheet
│   │   └── sql
│   │ 
│   ├──interfaces
│   │   ├── adapter
│   │   ├── controller
│   │   └── datastore
│   │ 
│   ├──logger
│   │ 
│   ├──registry
│   │ 
│   └──usecase
│       ├── inputport
│       ├── interactor
│       └── outputport
│
├ .env
├ Docker
├ go.mod
├ main.go
└ README.md
