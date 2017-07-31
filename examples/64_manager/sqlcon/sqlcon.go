package sqlcon

import (
	"database/sql"
	"fmt"
)

// SQLConController ... creates a new sql pool controller
type SQLConController struct {
	Connection *sql.DB
	Config     *Config
}

// NewSQLPool ... create a new sql pool
func NewSQLConnection(config *Config) (*SQLConController, error) {
	conn, _ := newConnection(config)

	instance := &SQLConController{
		Connection: conn,
		Config:     config,
	}

	return instance, nil
}

// GetConnection ... get connection
func (instance *SQLConController) GetConnection() (*sql.DB, error) {
	conn := instance.Connection

	if conn == nil {
		return nil, fmt.Errorf("could not get a connection")
	}

	return conn, nil
}

// AddConnection ... set connection
func (instance *SQLConController) SetConnection(config *Config) (*sql.DB, error) {
	var err error

	instance.Connection, err = newConnection(config)

	return instance.Connection, err
}

func newConnection(config *Config) (*sql.DB, error) {
	var conn *sql.DB
	var err error

	if conn, err = sql.Open(config.Driver, config.Endpoint); err != nil {
		return nil, err
	}

	conn.SetMaxIdleConns(config.MaxIdleConnections)
	conn.SetMaxOpenConns(config.MaxOpenConnections)

	if err = conn.Ping(); err != nil {
		return nil, err
	}

	return conn, nil
}
