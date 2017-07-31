package pm

import (
	"database/sql"
	"fmt"
	"github.com/labstack/gommon/log"
	"golang-learn/examples/64_manager/sqlcon"
)

// -------------- SQL POOLS --------------
// NewSQLPool ... creates a new sql connection pool
func (instance *manager) NewSQLConnection(config *sqlcon.Config) (*sqlcon.SQLConController, error) {
	log.Infof(fmt.Sprintf("sqlcon, connection created"))
	return sqlcon.NewSQLConnection(config)
}

// -------------- METHODS --------------
// GetConnection ... get a sql connection with key
func (instance *manager) GetConnection(key string) (*sql.DB, error) {
	connection, err := instance.SqlConController[key].GetConnection()
	return connection, err
}

// AddConnection ... add a connection with key
func (instance *manager) AddConnection(key string, SqlConController *sqlcon.SQLConController) error {
	instance.SqlConController[key] = SqlConController
	log.Infof(fmt.Sprintf("sqlcon, connection '%s' added", key))

	return nil
}

// RemConnection ... remove the connection by bey
func (instance *manager) RemConnection(key string) (*sql.DB, error) {
	// get connection
	controller := instance.SqlConController[key]

	// delete connection
	delete(instance.SqlConController, key)
	log.Infof(fmt.Sprintf("sqlcon, connection '%s' removed", key))

	return controller.Connection, nil
}
