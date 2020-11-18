package ports

import "github.com/jkarlos000/sc-market/provider/internal/core/domain"

type ProvidersService interface {
	Get(id int) (*domain.Provider, error)
	Create(provider *domain.Provider, password string) error
	List() ([]*domain.Provider, error)
	Delete(id int) error
	Deletes(id []int) error
	Update(id int, provider *domain.Provider) error
	// Uses Case
	Login(email, password string) (bool, *domain.Provider, error)
	Ban(id int) error
	Unban(id int) error
}
