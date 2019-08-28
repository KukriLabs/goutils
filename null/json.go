package null

import (
	"encoding/json"
	"errors"

	"github.com/jmoiron/sqlx/types"
)

type NullJSONText struct {
	types.NullJSONText
}

func (nt *NullJSONText) MarshalJSON() ([]byte, error) {
	if !nt.Valid {
		return []byte("null"), nil
	}
	return nt.JSONText.MarshalJSON()
}

func (nt *NullJSONText) UnmarshalJSON(data []byte) error {
	if nt == nil {
		return errors.New("Trying to UnmarshalJSON into nil struct")
	}
	err := json.Unmarshal(data, &nt.JSONText)
	nt.Valid = (err == nil)
	return err
}
