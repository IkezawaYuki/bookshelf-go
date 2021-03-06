package mysql

import (
	"database/sql"
	"fmt"
	"github.com/IkezawaYuki/bookshelf-go/src/interfaces/datastore"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

type Handler struct {
	db *sql.DB
}

func NewMySQLHandler(db *sql.DB) datastore.DBHandler {
	return &Handler{
		db: db,
	}
}

func GetMySQLConnection() *sql.DB {
	conn, err := sql.Open("mysql",
		fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=true",
			os.Getenv("DB_USER"),
			os.Getenv("DB_PASS"),
			os.Getenv("DB_HOST"),
			os.Getenv("DB_NAME"),
		))
	if err != nil {
		log.Fatalln(err)
	}
	return conn
}

func (m *Handler) Exec(query string, args ...interface{}) (datastore.Result, error) {
	return m.db.Exec(query, args...)
}

func (m *Handler) Query(query string, args ...interface{}) (datastore.Rows, error) {
	return m.db.Query(query, args...)
}

func (m *Handler) QueryRow(query string, args ...interface{}) datastore.Row {
	return m.db.QueryRow(query, args...)
}

func (m *Handler) Begin() (datastore.Tx, error) {
	tx, err := m.db.Begin()
	t := Tx{
		tx: tx,
	}
	return t, err
}

func (m *Handler) Close() error {
	return m.db.Close()
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
