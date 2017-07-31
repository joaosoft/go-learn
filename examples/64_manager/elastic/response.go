package elastic

import "encoding/json"

// Response ...
type Response struct {
	Took   int                `json:"took,omitempty"`
	Hits   ResponseHits       `json:"hits,omitempty"`
	Errors bool               `json:"errors,omitempty"`
	Items  []*json.RawMessage `json:"items,omitempty"`
}

// ResponseHits ...
type ResponseHits struct {
	Total int           `json:"total,omitempty"`
	Hits  []ResponseHit `json:"hits,omitempty"`
}

// ResponseHit ...
type ResponseHit struct {
	Index  string          `json:"_index,omitempty"`
	Type   string          `json:"_type,omitempty"`
	ID     string          `json:"_id,omitempty"`
	Source json.RawMessage `json:"_source,omitempty"`
}
