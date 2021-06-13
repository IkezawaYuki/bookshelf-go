# bookshelf-go

## Summary
読書メーターを模倣した本管理API



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