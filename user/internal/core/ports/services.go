package ports

import "github.com/jkarlos000/sc-market/user/internal/core/domain"

type UsersService interface {
	Get(id int) (domain.User, error)
	Create(user domain.User) error
	List() ([]domain.User, error)
	Delete(id int) error
	Deletes(id []int) error
	Update(id int, user domain.User) error
	// Uses Case
	Login(email string, password string) error
	Ban(id int) error
}
