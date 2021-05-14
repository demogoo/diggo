package services

import (
	"github.com/demogoo/diggo/internal/config"
	"github.com/demogoo/diggo/internal/data"
	"github.com/demogoo/diggo/internal/models"
)

type MemberService struct {
	config *config.Config
}

func NewMemberService(config *config.Config) *MemberService {
	return &MemberService{config: config}
}

func (m *MemberService) SelectAll() []*models.Member {
	return data.Members
}
