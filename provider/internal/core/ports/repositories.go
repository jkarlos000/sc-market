package ports

import "github.com/jkarlos000/sc-market/provider/internal/core/domain"

type ProvidersRepository interface {
	Get(id int) (*domain.Provider, error)
	GetByEmail(email string) (*domain.Provider, error)
	List() ([]*domain.Provider, error)
	Save(provider *domain.Provider) error
	Update(id int, provider *domain.Provider) error
	Delete(id int) error
}
