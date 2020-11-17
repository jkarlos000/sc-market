package ports

import "github.com/jkarlos000/sc-market/client/internal/core/domain"

type ClientsService interface {
	Get(id int) (*domain.Client, error)
	Create(Client *domain.Client, password string) error
	List() ([]*domain.Client, error)
	Delete(id int) error
	Deletes(id []int) error
	Update(id int, Client *domain.Client) error
	// Uses Case
	Login(email, password string) (bool, *domain.Client, error)
	Ban(id int) error
	Unban(id int) error
}
