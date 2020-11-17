package domain

import (
	"database/sql"
	"time"

	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	Email	string	`gorm:"unique;not null"`
	Salt	string
	Verify	string
	SessionKeyAuth	string
	Names	string
	LastNames	string
	BirdDate	string
	Gender	string
	Address	string
	Reference	string
	CI	string	`gorm:"unique;not null"`
	Telephone	string
	JoinDate	sql.NullTime
	LastIP	string
	Locked	*bool	`gorm:"default:false"`
	LastLogin	time.Time
}
