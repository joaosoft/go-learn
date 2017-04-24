package eventstore

import (
	"database/sql"
	"errors"
	"fmt"

	uuid "github.com/satori/go.uuid"
)

// EventStoreSQL represents the implementation of the IProductOffersAggregateRepository for a postgres db.
type EventStoreSQL struct {
	dbcontext IDBContext
}

// NewEventStoreSQL creates a new EventStoreSQL
func NewEventStoreSQL(dbcontext IDBContext) *EventStoreSQL {
	eventStore := new(EventStoreSQL)

	eventStore.dbcontext = dbcontext

	return eventStore
}

// GetDBContext ...
func (eventStore *EventStoreSQL) GetDBContext() IDBContext {
	return eventStore.dbcontext
}

// SetDBContext ...
func (eventStore *EventStoreSQL) SetDBContext(dbcontext IDBContext) {
	eventStore.dbcontext = dbcontext
}

// GetByID ...
func (eventStore *EventStoreSQL) GetByID(aggregateType string, aggregateID uuid.UUID) (*Aggregate, error) {
	aggregate, err := eventStore.fetchByID(aggregateType, aggregateID)

	if err != nil {
		return nil, err
	} else if aggregate == nil {
		return nil, nil
	}

	return aggregate, nil
}

// GetBySlug ...
func (eventStore *EventStoreSQL) GetBySlug(aggregateType string, aggregateSlug string) (*Aggregate, error) {
	aggregateID, err := eventStore.getIDForSlug(aggregateType, aggregateSlug)
	if err != nil {
		return nil, err
	}

	if aggregateID == uuid.Nil {
		return nil, nil
	}

	return eventStore.GetByID(aggregateType, aggregateID)
}

// Store ... Stores and Aggregate
func (eventStore *EventStoreSQL) Store(aggregate *Aggregate) (err error) {
	var tx *sql.Tx
	var db *sql.DB

	if con, found := eventStore.dbcontext.(*sql.DB); found {
		tx, err = con.Begin()

		if err != nil {
			return
		}
		db = con

		eventStore.SetDBContext(tx)
	}

	defer func() {
		if tx != nil {
			eventStore.SetDBContext(db)
			if err != nil {
				tx.Rollback()
			} else {
				err = tx.Commit()
			}
		}
	}()

	err = eventStore.storeEvents(aggregate)
	if err != nil {
		return
	}

	err = eventStore.storeSnapshot(aggregate)
	if err != nil {
		return
	}

	if aggregate.Slug != "" {
		err = eventStore.storeSlugMap(aggregate)
		if err != nil {
			return
		}
	}

	return
}

func (eventStore *EventStoreSQL) getIDForSlug(aggregateType string, aggregateSlug string) (uuid.UUID, error) {
	var row *sql.Row
	var aggregateID uuid.UUID
	sqlQuery := `
		SELECT aggregate_id
		FROM es_aggregate_slug_map
		WHERE slug = $1 AND aggregate_type = $2
	`

	row = eventStore.dbcontext.QueryRow(sqlQuery, aggregateSlug, aggregateType)

	err := row.Scan(&aggregateID)
	if err != nil {
		// special error case for when there were no returned rows
		if err == sql.ErrNoRows {
			return uuid.Nil, nil
		}
		return uuid.Nil, err
	}

	return aggregateID, err
}

// FetchByIDPg fetches an aggregate from a PostgreSQL database.
func (eventStore *EventStoreSQL) fetchByID(aggregateType string, aggregateID uuid.UUID) (*Aggregate, error) {
	var row *sql.Row

	aggregateWrapper := &Aggregate{}
	sqlQuery := `
		SELECT 
		  aggregate_id, 
		  aggregate_type, 
		  aggregate_version, 
		  data
		FROM es_snapshots
		WHERE aggregate_id = $1 AND aggregate_type = $2
	`

	row = eventStore.dbcontext.QueryRow(sqlQuery, aggregateID, aggregateType)

	if err := row.Scan(
		&aggregateWrapper.ID,
		&aggregateWrapper.Type,
		&aggregateWrapper.Version,
		&aggregateWrapper.Payload,
	); err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}

		return nil, err
	}

	return aggregateWrapper, nil
}

func (eventStore *EventStoreSQL) storeSnapshot(aggregate *Aggregate) error {
	_, err := eventStore.dbcontext.Exec(`
		INSERT INTO es_snapshots
		(aggregate_id, aggregate_type, aggregate_version, data)
		VALUES($1, $2, $3, $4)
        ON CONFLICT (aggregate_id, aggregate_type) DO UPDATE SET 
		  aggregate_version = EXCLUDED.aggregate_version,
		  data = EXCLUDED.data
	`,
		aggregate.ID,
		aggregate.Type,
		aggregate.Version,
		aggregate.Payload,
	)

	if err != nil {
		return err
	}

	return nil
}

func (eventStore *EventStoreSQL) storeEvents(aggregate *Aggregate) error {
	currentVersion := 0

	row := eventStore.dbcontext.QueryRow(`
		SELECT aggregate_version
		FROM es_aggregates
		WHERE aggregate_id = $1 AND aggregate_type = $2
		FOR UPDATE NOWAIT
	`, aggregate.ID, aggregate.Type)

	if err := row.Scan(&currentVersion); err != nil && err != sql.ErrNoRows {
		return err
	}

	if currentVersion != aggregate.Version {
		return errors.New("Concurrency Exception")
	}

	stmt, err := eventStore.dbcontext.Prepare(`
		INSERT INTO es_events
		(event_id, event_type, aggregate_id, aggregate_type, aggregate_version, data)
		VALUES ($1,$2,$3,$4,$5,$6)
	`)

	if err != nil {
		return err
	}

	defer stmt.Close()

	for _, eventWrapper := range aggregate.Changes {
		aggregate.Version++

		if _, err = stmt.Exec(
			eventWrapper.ID,
			eventWrapper.Type,
			aggregate.ID,
			aggregate.Type,
			aggregate.Version,
			eventWrapper.Payload,
		); err != nil {
			return err
		}
	}

	if currentVersion == 0 {
		_, err = eventStore.dbcontext.Exec(`
		INSERT INTO es_aggregates
		(aggregate_id,  aggregate_type, aggregate_version)
		VALUES($1,$2,$3)
	`, aggregate.ID, aggregate.Type, aggregate.Version)
	} else {
		_, err = eventStore.dbcontext.Exec(`
		UPDATE es_aggregates
		SET aggregate_version = $1
		WHERE aggregate_id = $2 AND aggregate_type = $3
	`, aggregate.Version, aggregate.ID, aggregate.Type)
	}

	if err != nil {
		return err
	}

	return nil
}

func (eventStore *EventStoreSQL) storeSlugMap(aggregate *Aggregate) error {
	// store product_external_id to product_id mapping
	stmt, err := eventStore.dbcontext.Prepare(`
	  INSERT INTO es_aggregate_slug_map (aggregate_id, aggregate_type, slug) 
	  VALUES ($1, $2, $3)
      ON CONFLICT (aggregate_id, aggregate_type) DO UPDATE
        SET slug = EXCLUDED.slug
	`)
	if err != nil {
		return err
	}

	res, err := stmt.Exec(aggregate.ID, aggregate.Type, aggregate.Slug)
	if err != nil {
		return err
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return err
	}

	if rows > 1 || rows < 0 {
		return fmt.Errorf("Error storing external id map [%d]", rows)
	}
	stmt.Close()

	return nil
}
