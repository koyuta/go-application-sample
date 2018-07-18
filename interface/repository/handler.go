package repository

import "context"

// MySQLHandler is the interface implemented by the object that
// acts the MySQL client.
type MySQLHandler interface {
	Execute(string, ...interface{}) (Result, error)
	ExecuteContext(context.Context, string, ...interface{}) (Result, error)
	Query(string, ...interface{}) (Rows, error)
	QueryContext(context.Context, string, ...interface{}) (Rows, error)
}

// Result is the interface implemented by the object that
// represents the result of MySQL.
type Result interface {
	LastInsertId() (int64, error)
	RowsAffected() (int64, error)
}

// Rows is the interface implemented by the object that represents
// the result rows of MySQL.
type Rows interface {
	Scan(...interface{}) error
	Next() bool
	Close() error
}
