package clientssrv
import (
	"fmt"
	"errors"

	"github.com/jkarlos000/sc-market/client/internal/core/domain"
	"github.com/jkarlos000/sc-market/client/internal/core/ports"
	srp "github.com/jkarlos000/srp6"
	)

type service struct {
	repo ports.ClientsRepository
}

func NewService(repo ports.ClientsRepository) *service {
	return &service{
		repo: repo,
	}
}

func (s *service) Get(id int) (*domain.Client, error) {
	return s.repo.Get(id)
}

func (s *service) Create(Client *domain.Client, password string) error {
	auth := srp.New()
	auth.GenerateSalt()
	hash := srp.Hash(Client.Email, password)
	auth.ComputeVerifier(hash)
	Client.Salt = auth.GetSalt()
	Client.Verify = auth.GetVerifier()
	return s.repo.Save(Client)
}

func (s *service) List() ([]*domain.Client, error) {
	return s.repo.List()
}

func (s *service) Delete(id int) error {
	return s.repo.Delete(id)
}

func (s *service) Deletes(id []int) error {
	for _, ClientId := range id{
		err := s.repo.Delete(ClientId)
		if err != nil {
			// Requiere LOG
			return fmt.Errorf("Ha ocurrido un error al eliminar los usuarios. Referencia: %v", err)
		}
	}
	return nil
}

func (s *service) Update(id int, Client *domain.Client) error {
	return s.repo.Update(id, Client)
}

func (s *service) Login(email, password string) (bool, *domain.Client, error) {
	Client, err := s.repo.GetByEmail(email)
	if err != nil {
		return false, nil, err
	}
	auth := srp.New()
	auth.SetSalt(Client.Salt)
	identifier := srp.Hash(email,password)
	auth.ComputeVerifier(identifier)
	if !auth.ProofVerifier(Client.Verify) {
		return false, nil, errors.New("La contraseña es incorrecta")
	}
	if *Client.Locked == true {
		return false, nil, errors.New("La cuenta se encuentra bloqueada, contacte con el administrador")
	}
	return true, Client, nil
}

func (s *service) Ban(id int) error {
	Client, err := s.repo.Get(id)
	if err != nil {
		return err
	}
	locked := true
	Client.Locked = &locked
	return s.repo.Update(id, Client)
}

func (s *service) Unban(id int) error {
	Client, err := s.repo.Get(id)
	if err != nil {
		return err
	}
	unlocked := false
	Client.Locked = &unlocked
	return s.repo.Update(id, Client)
}
