package models

import (
	validation "github.com/go-ozzo/ozzo-validation"
)

type Order struct {
	Id          string   `json:"order_uid"`
	TrackNumber string   `json:"track_number"`
	Entry       string   `json:"entry"`
	Delivery    Delivery `json:"delivery"`
	Payment     Payment  `json:"payment"`
	Items
	Locale            string `json:"locale"`
	InternalSignature string `json:"internal_signature"`
	CustomerId        string `json:"customer_id"`
	DeliveryService   string `json:"delivery_service"`
	ShardKey          string `json:"shardkey"`
	SmId              int    `json:"sm_id"`
	DateCreated       string `json:"date_created"`
	OofShard          string `json:"oof_shard"`
}

func (o *Order) Validate() error {
	return validation.ValidateStruct(
		o,
		validation.Field(&o.Id, validation.Required),
		validation.Field(&o.TrackNumber, validation.Required),
		validation.Field(&o.Entry, validation.Required),
		validation.Field(&o.Delivery, validation.Required),
		validation.Field(&o.Payment, validation.Required),
		validation.Field(&o.Items, validation.Required),
		validation.Field(&o.Locale, validation.Required),
		validation.Field(&o.InternalSignature),
		validation.Field(&o.CustomerId, validation.Required),
		validation.Field(&o.DeliveryService, validation.Required),
		validation.Field(&o.ShardKey, validation.Required),
		validation.Field(&o.SmId, validation.Required),
		validation.Field(&o.DateCreated, validation.Required),
		validation.Field(&o.OofShard, validation.Required),
	)
}
