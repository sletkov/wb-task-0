package models

import (
	"database/sql/driver"
	"encoding/json"
)

type Delivery struct {
	Name    string `json:"name"`
	Phone   string `json:"phone"`
	Zip     string `json:"zip"`
	City    string `json:"city"`
	Address string `json:"address"`
	Region  string `json:"region"`
	Email   string `json:"email"`
}

func (d *Delivery) Value() driver.Value {
	value, _ := json.Marshal(d)
	return value
}
