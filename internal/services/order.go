package services

import (
	"github.com/demogoo/diggo/internal/data"
	"github.com/demogoo/diggo/internal/models"
	"github.com/demogoo/diggo/internal/storage"
)

type OrderService struct {
	Cache *storage.CacheConn
}

func NewOrderService(cache *storage.CacheConn) *OrderService {
	return &OrderService{Cache: cache}
}

func (m *OrderService) GetAll() []*models.Order {
	return data.Orders
}
