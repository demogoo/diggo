package services

import (
	"github.com/demogoo/diggo/internal/config"
	"github.com/demogoo/diggo/internal/data"
	"github.com/demogoo/diggo/internal/models"
)

type OrderService struct {
	config *config.Config
}

func NewOrderService(config *config.Config) *OrderService {
	return &OrderService{config: config}
}

func (m *OrderService) GetAll() []*models.Order {
	return data.Orders
}
