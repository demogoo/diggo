package services

import (
	"testing"

	"github.com/demogoo/diggo/internal/config"
	"github.com/demogoo/diggo/internal/data"
	"github.com/stretchr/testify/assert"
)

func TestMemberFindAll(t *testing.T) {
	conf := config.NewConfig()
	memberService := NewMemberService(conf)
	members := memberService.SelectAll()
	assert.Equal(t, len(data.Members), len(members))
	assert.Equal(t, data.Members[0], members[0])
}
