package postgres

import (
	"database/sql"
	"sletkov/backend/wb-task-0/internal/app/models"
	"sletkov/backend/wb-task-0/internal/app/store"
)

type OrderRepository struct {
	store *Store
}

func (r *OrderRepository) Create(order *models.Order) error {
	if err := r.store.db.QueryRow("INSERT INTO orders VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14) RETURNING order_uid",
		order.Id,
		order.TrackNumber,
		order.Entry,
		order.Delivery.Value(),
		order.Payment.Value(),
		order.Items.Value(),
		order.Locale,
		order.InternalSignature,
		order.CustomerId,
		order.DeliveryService,
		order.ShardKey,
		order.SmId,
		order.DateCreated,
		order.OofShard,
	).Scan(&order.Id); err != nil {
		return err
	}
	return nil
}

func (r *OrderRepository) GetAll() ([]*models.Order, error) {
	var orders []*models.Order
	query := "SELECT * FROM orders"

	if err := r.store.db.QueryRow(query).Scan(); err != nil {
		return nil, err
	}

	return orders, nil
}

func (r *OrderRepository) GetOrderById(id int) (*models.Order, error) {
	order := &models.Order{}

	if err := r.store.db.QueryRow(
		"SELECT * FROM orders WHERE id = $1",
		id,
	).Scan(); err != nil {
		if err == sql.ErrNoRows {
			return nil, store.ErrRecordNotFound
		}
		return nil, err
	}

	return order, nil
}
