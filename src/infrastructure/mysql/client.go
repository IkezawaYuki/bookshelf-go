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
	tx, err := m.db.Begin()
	t := Tx{
		tx: tx,
	}
	return t, err
}

type Tx struct {
	tx *sql.Tx
}

func (t Tx) Commit() error {
	return t.tx.Commit()
}

func (t Tx) Rollback() error {
	return t.tx.Rollback()
}

func (t Tx) Exec(query string, args ...interface{}) (datastore.Result, error) {
	return t.tx.Exec(query, args...)
}

func (t Tx) Query(query string, args ...interface{}) (datastore.Row, error) {
	return t.tx.Query(query, args...)
}

func (t Tx) QueryRow(query string, args ...interface{}) datastore.Row {
	return t.tx.QueryRow(query, args...)
}
