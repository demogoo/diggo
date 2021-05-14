package services

import (
	"testing"

	"github.com/demogoo/diggo/internal/config"
	"github.com/demogoo/diggo/internal/data"
	"github.com/demogoo/diggo/internal/storage"
	"github.com/stretchr/testify/assert"
)

func TestMemberFindAll(t *testing.T) {
	conf := config.NewConfig()
	db := storage.NewDBConn(conf)
	memberService := NewMemberService(db)
	members := memberService.SelectAll()
	assert.Equal(t, len(data.Members), len(members))
	assert.Equal(t, data.Members[0], members[0])
}
