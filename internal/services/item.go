package services

import (
	"github.com/demogoo/diggo/internal/data"
	"github.com/demogoo/diggo/internal/models"
	"github.com/demogoo/diggo/internal/storage"
)

type ItemService struct {
	MongoDB *storage.MongoDBConn
}

func NewItemService(mongoDB *storage.MongoDBConn) *ItemService {
	return &ItemService{MongoDB: mongoDB}
}

func (i *ItemService) FindAll() []*models.Item {
	return data.Items
}
