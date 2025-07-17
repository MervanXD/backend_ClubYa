package utils

import (
	"database/sql"
	"encoding/json"
)

type NullString struct {
	sql.NullString
}

func (ns *NullString) Scan(value interface{}) error {
	return ns.NullString.Scan(value)
}

func (ns NullString) MarshalJSON() ([]byte, error) {
	if !ns.Valid {
		// return []byte("null"), nil
		return json.Marshal("-")
	}
	return json.Marshal(ns.String)
}

func (ns *NullString) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		ns.Valid = false
		return nil
	}
	ns.Valid = true
	return json.Unmarshal(data, &ns.String)
}
