package null

import (
	"encoding/json"
	"errors"

	"github.com/lib/pq"
)

type NullPQTime struct {
	pq.NullTime
}

func (nt *NullPQTime) MarshalJSON() ([]byte, error) {
	if !nt.Valid {
		return []byte("null"), nil
	}
	return nt.Time.MarshalJSON()
}

func (nt *NullPQTime) UnmarshalJSON(data []byte) error {
	if nt == nil {
		return errors.New("Trying to UnmarshalJSON into nil struct")
	}
	err := json.Unmarshal(data, &nt.Time)
	nt.Valid = (err == nil)
	return err
}
