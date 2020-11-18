package gorm

import (
	"errors"
	"fmt"
	"github.com/hashicorp/go-hclog"
	"github.com/jinzhu/gorm"
	"github.com/jkarlos000/sc-market/provider/internal/core/domain"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"log"
)

type repository struct{
	DBConn *gorm.DB
}

func NewProvidersRepository(ip, user, password, dbname, tmz string, port int) *repository {
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
	repo.DBConn.AutoMigrate(domain.Provider{})
	return &repo
}

func (r *repository) Get(id int) (*domain.Provider, error) {
	var provider domain.Provider
	proveedor := r.DBConn.Find(&provider, id )
	if proveedor.RecordNotFound() {
		return nil, errors.New("Proveedor no encontrado")
	}
	return &provider, nil
}

func (r *repository) GetByEmail(email string) (*domain.Provider, error) {
	var provider domain.Provider
	proveedor := r.DBConn.Find(&provider, &domain.Provider{Email: email} )
	if proveedor.RecordNotFound() {
		return nil, errors.New("Email no registrado.")
	}
	return &provider, nil
}

func (r *repository) List() ([]*domain.Provider, error) {
	providers := []*domain.Provider{}
	r.DBConn.Find(providers)
	return providers, nil
}

func (r *repository) Save(provider *domain.Provider) error {
	if r := r.DBConn.Create(provider).Error; r != nil {
		return errors.New("No se ha logrado crear el proveedor, revise los campos")
	}
	return nil
}

func (r *repository) Update(id int, provider *domain.Provider) error {
	proveedor, err := r.Get(id)
	if err != nil {
		return err
	}
	r.DBConn.Model(proveedor).Omit("id").Updates(provider)
	return nil
}

func (r *repository) Delete(id int) error {
	provider, err := r.Get(id)
	if err != nil {
		return err
	}
	r.DBConn.Delete(provider)
	_, errVerify := r.Get(id)
	if errVerify != nil {
		return errors.New("Ha ocurrido un error interno al eliminar este proveedor")
	}
	return nil
}

