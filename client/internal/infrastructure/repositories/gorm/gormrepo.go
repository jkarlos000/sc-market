package gorm

import (
	"errors"
	"fmt"
	"github.com/jkarlos000/sc-market/client/internal/core/domain"

	"github.com/hashicorp/go-hclog"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type repository struct{
	DBConn *gorm.DB
}

func NewClientsRepository(ip, user, password, dbname, tmz string, port int) *repository {
	var repo repository
	var err error
	logger := hclog.Default()
	repo.DBConn, err = gorm.Open("postgres", fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable TimeZone=%s", ip, port, user, password, dbname, tmz))

	if err != nil {
		logger.Error("Error interno en la base de datos", "err", err)
		panic("failed to connect database")
	}
	logger.Info("Migrando modelo hacia la base de datos")
	repo.DBConn.AutoMigrate(domain.Client{})
	return &repo
}

func (r *repository) Get(id int) (*domain.Client, error) {
	var client *domain.Client
	cliente := r.DBConn.First(client, id )
	if cliente.RecordNotFound() {
		return nil, errors.New("Usuario no encontrado")
	}
	return client, nil
}

func (r *repository) GetByEmail(email string) (*domain.Client, error) {
	var client *domain.Client
	cliente := r.DBConn.First(client, &domain.Client{Email: email} )
	if cliente.RecordNotFound() {
		return nil, errors.New("Email no registrado.")
	}
	return client, nil
}

func (r *repository) List() ([]*domain.Client, error) {
	clients := []*domain.Client{}
	r.DBConn.Find(clients)
	return clients, nil
}

func (r *repository) Save(client *domain.Client) error {
	if r := r.DBConn.Create(client).Error; r != nil {
		return errors.New("No se ha logrado crear el cliente, revise los campos")
	}
	return nil
}

func (r *repository) Update(id int, client *domain.Client) error {
	cliente, err := r.Get(id)
	if err != nil {
		return err
	}
	r.DBConn.Model(cliente).Omit("id").Updates(client)
	return nil
}

func (r *repository) Delete(id int) error {
	client, err := r.Get(id)
	if err != nil {
		return err
	}
	r.DBConn.Delete(client)
	_, errVerify := r.Get(id)
	if errVerify != nil {
		return errors.New("Ha ocurrido un error interno al eliminar este cliente")
	}
	return nil
}
