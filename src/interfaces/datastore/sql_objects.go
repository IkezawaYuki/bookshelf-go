package datastore

type Result interface {
	LastInsertId() (int64, error)
	RowsAffected() (int64, error)
}

type Rows interface {
	Scan(...interface{}) error
	Next() bool
	Close() error
}

type Row interface {
	Scan(...interface{}) error
}

type Stmt interface {
	Exec(query string, args ...interface{}) (Result, error)
	Query(query string, args ...interface{}) (Row, error)
	QueryRow(query string, args ...interface{}) Row
	Close() error
}

type Tx interface {
	Commit() error
	Rollback() error
	Exec(query string, args ...interface{}) (Result, error)
	Query(query string, args ...interface{}) (Row, error)
	QueryRow(query string, args ...interface{}) Row
}
