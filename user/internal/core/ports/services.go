package ports

import "github.com/jkarlos000/sc-market/user/internal/core/domain"

type UsersService interface {
	Get(id int) (*domain.User, error)
	Create(user *domain.User, password string) error
	List() ([]*domain.User, error)
	Delete(id int) error
	Deletes(id []int) error
	Update(id int, user *domain.User) error
	// Uses Case
	Login(email, password string) (bool, *domain.User, error)
	Ban(id int) error
	Unban(id int) error
}
