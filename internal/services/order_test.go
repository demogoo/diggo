package services

import (
	"testing"

	"github.com/demogoo/diggo/internal/config"
	"github.com/demogoo/diggo/internal/data"
	"github.com/stretchr/testify/assert"
)

func TestOrderGet(t *testing.T) {
	conf := config.NewConfig()
	orderService := NewOrderService(conf)
	orders := orderService.GetAll()
	assert.Equal(t, len(data.Orders), len(orders))
	assert.Equal(t, data.Orders[0], orders[0])
}
