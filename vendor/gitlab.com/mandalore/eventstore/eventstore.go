package eventstore

import (
	"database/sql"

	uuid "github.com/satori/go.uuid"
)

// IEventStore ...
type IEventStore interface {
	GetBySlug(string, string) (*Aggregate, error)
	GetByID(string, uuid.UUID) (*Aggregate, error)
	Store(*Aggregate) error
	SetDBContext(IDBContext)
}

// IDBContext ...
type IDBContext interface {
	Exec(query string, args ...interface{}) (sql.Result, error)
	Prepare(query string) (*sql.Stmt, error)
	Query(query string, args ...interface{}) (*sql.Rows, error)
	QueryRow(query string, args ...interface{}) *sql.Row
}

// Aggregate ...
type Aggregate struct {
	ID      uuid.UUID
	Type    string // aggregate type
	Version int
	Slug    string // foreign id used only for external mappings
	Payload []byte
	Changes []*Event
}

// Event ...
type Event struct {
	ID      uuid.UUID
	Type    string
	Payload []byte
}
