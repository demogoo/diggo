package services

import (
	"testing"

	"github.com/demogoo/diggo/internal/config"
	"github.com/demogoo/diggo/internal/data"
	"github.com/demogoo/diggo/internal/storage"
	"github.com/stretchr/testify/assert"
)

func TestItemFindAll(t *testing.T) {
	conf := config.NewConfig()
	mongoDB := storage.NewMongoDBConn(conf)
	itemService := NewItemService(mongoDB)
	items := itemService.FindAll()
	assert.Equal(t, len(data.Items), len(items))
	assert.Equal(t, data.Items[0], items[0])
}
