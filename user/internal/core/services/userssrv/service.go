package userssrv

import (
	"errors"
	"fmt"
	"github.com/jkarlos000/sc-market/user/internal/core/domain"
	"github.com/jkarlos000/sc-market/user/internal/core/ports"

	"github.com/jkarlos000/srp6"
)

type service struct {
	repo ports.UsersRepository
}

func NewService(repo ports.UsersRepository) *service {
	return &service{
		repo: repo,
	}
}

func (s *service) Get(id int) (*domain.User, error) {
	return s.repo.Get(id)
}

func (s *service) Create(user *domain.User) error {
	return s.repo.Save(user)
}

func (s *service) List() ([]*domain.User, error) {
	return s.repo.List()
}

func (s *service) Delete(id int) error {
	return s.repo.Delete(id)
}

func (s *service) Deletes(id []int) error {
	for _, userId := range id{
		err := s.repo.Delete(userId)
		if err != nil {
			// Requiere LOG
			return fmt.Errorf("Ha ocurrido un error al eliminar los usuarios. Referencia: %v", err)
		}
	}
	return nil

}

func (s *service) Update(id int, user *domain.User) error {
	return s.repo.Update(id, user)
}

func (s *service) Ban(id int) error {
	user, err := s.repo.Get(id)
	if err != nil {
		return err
	}
	locked := true
	user.Locked = &locked
	return s.repo.Update(id, user)
}

func (s *service) UnBan(id int) error {
	user, err := s.repo.Get(id)
	if err != nil {
		return err
	}
	unlocked := false
	user.Locked = &unlocked
	return s.repo.Update(id, user)
}

func (s*service) Login(email, password string) (bool, *domain.User, error) {
	user, err := s.repo.GetByEmail(email)
	if err != nil {
		return false, nil, err
	}
	auth := srp.New()
	auth.SetSalt(user.Salt)
	identifier := srp.Hash(email,password)
	auth.ComputeVerifier(identifier)
	if !auth.ProofVerifier(user.Verify) {
		return false, nil, errors.New("La contraseña es incorrecta")
	}
	return true, user, nil
}
