package gorm

import (
	"errors"
	"fmt"
	"github.com/hashicorp/go-hclog"
	"github.com/jinzhu/gorm"
	"github.com/jkarlos000/sc-market/user/internal/core/domain"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"log"
)

type repository struct{
	DBConn *gorm.DB
}

func NewUsersRepository(ip, user, password, dbname, tmz string, port int) *repository {
	var repo repository
	var err error
	logger := hclog.Default()
	repo.DBConn, err = gorm.Open("postgres", fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable TimeZone=%s", ip, port, user, password, dbname, tmz))

	if err != nil {
		logger.Error("Error interno en la base de datos", "err", err)
		log.Printf("Error with: %v\n", err)
		panic("failed to connect database")
	}
	logger.Info("Migrando modelo hacia la base de datos")
	repo.DBConn.AutoMigrate(domain.User{})
	return &repo
}

func (r *repository) Get(id int) (*domain.User, error) {
	var user domain.User
	usuario := r.DBConn.Find(&user, id )
	if usuario.RecordNotFound() {
		return nil, errors.New("Usuario no encontrado")
	}
	return &user, nil
}

func (r *repository) GetByEmail(email string) (*domain.User, error) {
	var user domain.User
	usuario := r.DBConn.Find(&user, &domain.User{Email: email} )
	if usuario.RecordNotFound() {
		return nil, errors.New("Email no registrado.")
	}
	return &user, nil
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

