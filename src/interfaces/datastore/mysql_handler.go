package datastore

type DBHandler interface {
	Exec(query string, args ...interface{}) (Result, error)
	Query(query string, args ...interface{}) (Row, error)
	QueryRow(query string, args ...interface{}) Row
	Begin() (Tx, error)
}
