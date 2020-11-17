package ports

import "github.com/jkarlos000/sc-market/client/internal/core/domain"

type ClientsRepository interface {
	Get(id int) (*domain.Client, error)
	GetByEmail(email string) (*domain.Client, error)
	List() ([]*domain.Client, error)
	Save(Client *domain.Client) error
	Update(id int, Client *domain.Client) error
	Delete(id int) error
}
