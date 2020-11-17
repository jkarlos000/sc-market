package gorm

import (
	"errors"
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/jkarlos000/sc-market/user/internal/core/domain"
)

type repository struct{
	DBConn *gorm.DB
}

func NewUsersRepository() *repository {
	var repo *repository
	var err error
	repo.DBConn, err = gorm.Open("postgres", fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable TimeZone=%s", "localhost", 5432, "usersvc", "m7TDiQqO7kb3aEY2", "erp_user", "America/La_Paz"))

	if err != nil {
		panic("failed to connect database")
	}
	repo.DBConn.AutoMigrate(domain.User{})
	return repo
}

func (r *repository) Get(id int) (*domain.User, error) {
	var user *domain.User
	usuario := r.DBConn.First(user, id )
	if usuario.RecordNotFound() {
		return nil, errors.New("Usuario no encontrado")
	}
	return user, nil
}

func (r *repository) GetByEmail(email string) (*domain.User, error) {
	var user *domain.User
	usuario := r.DBConn.First(user, &domain.User{Email: email} )
	if usuario.RecordNotFound() {
		return nil, errors.New("Email no registrado.")
	}
	return user, nil
}

func (r *repository) List() ([]*domain.User, error) {
	users := []*domain.User{}
	r.DBConn.Find(users)
	return users, nil
}

func (r *repository) Save(user *domain.User) error {
	if r := r.DBConn.Create(user).Error; r != nil {
		return errors.New("No se ha logrado crear el usuario, revise los campos")
	}
	return nil
}

func (r *repository) Update(id int, user *domain.User) error {
	usuario, err := r.Get(id)
	if err != nil {
		return err
	}
	r.DBConn.Model(usuario).Omit("id").Updates(user)
	return nil
}

func (r *repository) Delete(id int) error {
	user, err := r.Get(id)
	if err != nil {
		return err
	}
	r.DBConn.Delete(user)
	_, errVerify := r.Get(id)
	if errVerify != nil {
		return errors.New("Ha ocurrido un error interno al eliminar este usuario")
	}
	return nil
}

