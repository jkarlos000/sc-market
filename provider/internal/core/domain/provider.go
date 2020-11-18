package domain

import (
	"database/sql"
	"time"
)

type Provider struct {
	ID        uint `gorm:"primary_key"`
	Email	string	`gorm:"unique;not null"`
	Salt	string
	Verify	string
	SessionKeyAuth	string
	Names	string
	BusinessName	string
	Nit	string
	LastNames	string
	BirthDate	string
	Gender	string
	Address	string
	Reference	string
	CI	string	`gorm:"unique;not null"`
	Telephone	string
	JoinDate	sql.NullTime
	LastIP	string
	Locked	*bool	`gorm:"default:false"`
	LastLogin	string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `sql:"index"`
}
