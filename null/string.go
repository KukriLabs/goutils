package null

import (
	"database/sql"
	"encoding/json"
	"errors"
)

type NullString struct {
	sql.NullString
}

func (ns *NullString) MarshalJSON() ([]byte, error) {
	if !ns.Valid {
		return []byte("null"), nil
	}
	return json.Marshal(ns.String)
}

func (ns *NullString) UnmarshalJSON(data []byte) error {
	if ns == nil {
		return errors.New("Trying to UnmarshalJSON into nil struct")
	}
	err := json.Unmarshal(data, &ns.String)
	ns.Valid = (err == nil)
	return err
}
