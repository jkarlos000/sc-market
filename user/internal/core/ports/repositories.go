package ports

import "github.com/jkarlos000/sc-market/user/internal/core/domain"

type UsersRepository interface {
	Get(id int) (*domain.User, error)
	GetByEmail(email string) (*domain.User, error)
	List() ([]*domain.User, error)
	Save(user *domain.User) error
	Update(id int, user *domain.User) error
	Delete(id int) error
}
