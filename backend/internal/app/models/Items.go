package models

import (
	"database/sql/driver"
	"encoding/json"
)

type Items struct {
	Items []Item `json:"items"`
}

func (i *Items) Value() driver.Value {
	value, _ := json.Marshal(i)
	return value
}
