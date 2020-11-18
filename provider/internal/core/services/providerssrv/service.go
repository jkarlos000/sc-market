package providerssrv

import (
	"errors"
	"fmt"
	"github.com/jkarlos000/sc-market/provider/internal/core/domain"
	"github.com/jkarlos000/sc-market/provider/internal/core/ports"
	srp "github.com/jkarlos000/srp6"
)

type service struct {
	repo ports.ProvidersRepository
}

func NewService(repo ports.ProvidersRepository) *service {
	return &service{
		repo: repo,
	}
}

func (s *service) Get(id int) (*domain.Provider, error) {
	return s.repo.Get(id)
}

func (s *service) Create(provider *domain.Provider, password string) error {
	auth := srp.New()
	auth.GenerateSalt()
	hash := srp.Hash(provider.Email, password)
	auth.ComputeVerifier(hash)
	provider.Salt = auth.GetSalt()
	provider.Verify = auth.GetVerifier()
	return s.repo.Save(provider)
}

func (s *service) List() ([]*domain.Provider, error) {
	return s.repo.List()
}

func (s *service) Delete(id int) error {
	return s.repo.Delete(id)
}

func (s *service) Deletes(id []int) error {
	for _, providerId := range id{
		err := s.repo.Delete(providerId)
		if err != nil {
			// Requiere LOG
			return fmt.Errorf("Ha ocurrido un error al eliminar los proveedores. Referencia: %v", err)
		}
	}
	return nil
}

func (s *service) Update(id int, provider *domain.Provider) error {
	return s.repo.Update(id, provider)
}

func (s *service) Login(email, password string) (bool, *domain.Provider, error) {
	provider, err := s.repo.GetByEmail(email)
	if err != nil {
		return false, nil, err
	}
	auth := srp.New()
	auth.SetSalt(provider.Salt)
	identifier := srp.Hash(email,password)
	auth.ComputeVerifier(identifier)
	if !auth.ProofVerifier(provider.Verify) {
		return false, nil, errors.New("La contraseña es incorrecta")
	}
	if *provider.Locked == true {
		return false, nil, errors.New("La cuenta se encuentra bloqueada, contacte con el administrador")
	}
	return true, provider, nil
}

func (s *service) Ban(id int) error {
	provider, err := s.repo.Get(id)
	if err != nil {
		return err
	}
	locked := true
	provider.Locked = &locked
	return s.repo.Update(id, provider)
}

func (s *service) Unban(id int) error {
	provider, err := s.repo.Get(id)
	if err != nil {
		return err
	}
	unlocked := false
	provider.Locked = &unlocked
	return s.repo.Update(id, provider)
}