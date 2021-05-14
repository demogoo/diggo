package services

import (
	"github.com/demogoo/diggo/internal/config"
	"github.com/demogoo/diggo/internal/data"
	"github.com/demogoo/diggo/internal/models"
)

type ItemService struct {
	config *config.Config
}

func NewItemService(config *config.Config) *ItemService {
	return &ItemService{config: config}
}

func (i *ItemService) FindAll() []*models.Item {
	return data.Items
}
