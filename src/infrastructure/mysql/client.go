package mysql

import (
	"database/sql"
	"github.com/IkezawaYuki/bookshelf-go/src/interfaces/datastore"

	_ "github.com/go-sql-driver/mysql"
)

type mysqlHandler struct {
	db *sql.DB
}

func NewMySQLHandler(db *sql.DB) datastore.MySQLHandler {
	return &mysqlHandler{
		db: db,
	}
}

func (m *mysqlHandler) Exec(query string, args ...interface{}) (datastore.Result, error) {
	return m.db.Exec(query, args...)
}

func (m *mysqlHandler) Query(query string, args ...interface{}) (datastore.Row, error) {
	return m.db.Query(query, args...)
}

func (m *mysqlHandler) QueryRow(query string, args ...interface{}) datastore.Row {
	return m.db.QueryRow(query, args...)
}

func (m *mysqlHandler) Begin() (datastore.Tx, error) {
	panic("implement me")
}
