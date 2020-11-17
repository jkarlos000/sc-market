package userssrv

import (
	"github.com/jkarlos000/sc-market/user/internal/core/domain"
	"github.com/jkarlos000/sc-market/user/internal/core/ports"
)

type service struct {
	repo ports.UsersRepository
}

func NewService(repo ports.UsersRepository) *service {
	return &service{repo: repo}
}

func (s *service) Get(id int) (domain.User, error) {
	return s.repo.Get(id)
}
