CREATE OR REPLACE FUNCTION update_updated_at()
RETURNS TRIGGER AS $$
BEGIN
   NEW.updated_at = now();
   RETURN NEW;
END;
$$ language 'plpgsql';

CREATE TABLE es_aggregates (
  aggregate_id        UUID NOT NULL,
  aggregate_type      TEXT NOT NULL,
  aggregate_version   INT NOT NULL,
  PRIMARY KEY(aggregate_id,aggregate_type)
);

CREATE TABLE es_snapshots (
  aggregate_id        UUID NOT NULL,
  aggregate_type      TEXT NOT NULL,
  aggregate_version   BIGINT NOT NULL,
  data                JSONB DEFAULT '{}',
  created_at          TIMESTAMP DEFAULT NOW(),
  updated_at          TIMESTAMP DEFAULT NOW(),
  PRIMARY KEY(aggregate_id, aggregate_type),
  FOREIGN KEY(aggregate_id, aggregate_type) 
    REFERENCES es_aggregates(aggregate_id, aggregate_type) INITIALLY DEFERRED
);

CREATE TRIGGER es_snapshots_updated_at BEFORE UPDATE
  ON es_snapshots FOR EACH ROW EXECUTE PROCEDURE update_updated_at();

CREATE TABLE es_events (
  event_id            UUID NOT NULL,
  event_type          TEXT NOT NULL,
  correlation_id      UUID DEFAULT NULL,
  aggregate_id        UUID NOT NULL,
  aggregate_type      TEXT NOT NULL,
  aggregate_version   INT NOT NULL,
  data                JSONB DEFAULT '{}',
  is_dispatched       BOOLEAN NOT NULL DEFAULT 'f',
  created_at          TIMESTAMP DEFAULT NOW(),
  PRIMARY KEY(aggregate_id, aggregate_type, aggregate_version),
  UNIQUE(event_id),
  FOREIGN KEY(aggregate_id, aggregate_type) 
    REFERENCES es_aggregates(aggregate_id, aggregate_type) INITIALLY DEFERRED
);

CREATE TABLE es_aggregate_slug_map (
  aggregate_id                  UUID NOT NULL,
  aggregate_type                TEXT NOT NULL,
  slug                 TEXT NOT NULL,
  created_at          TIMESTAMP DEFAULT NOW(),
  PRIMARY KEY(aggregate_id, aggregate_type),
  FOREIGN KEY(aggregate_id, aggregate_type) 
    REFERENCES es_aggregates(aggregate_id, aggregate_type) INITIALLY DEFERRED,
  UNIQUE(slug,aggregate_type)
);
