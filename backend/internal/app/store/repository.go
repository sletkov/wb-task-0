package store

import "sletkov/backend/wb-task-0/internal/app/models"

type OrderRepository interface {
	Create(*models.Order) error
	GetAll() ([]*models.Order, error)
	GetOrderById(int) (*models.Order, error)
}
