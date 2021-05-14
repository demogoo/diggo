package services

import (
	"github.com/demogoo/diggo/internal/data"
	"github.com/demogoo/diggo/internal/models"
	"github.com/demogoo/diggo/internal/storage"
)

type MemberService struct {
	DB *storage.DBConn
}

func NewMemberService(db *storage.DBConn) *MemberService {
	return &MemberService{DB: db}
}

func (m *MemberService) SelectAll() []*models.Member {
	return data.Members
}
