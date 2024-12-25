package exchange

import (
	"github.com/agniBit/cryptonian/internal/repository/user"
	"github.com/agniBit/cryptonian/internal/storage/postgres"
)

type Account struct {
	ID         string `gorm:"primary_key"`
	UserId     string `gorm:"not null;index"`
	User       *user.User
	ExchangeID string
	Exchange   *Exchange
	ApiKey     string `gorm:"not null"`
	SecretKey  string `gorm:"not null"`
	IsTestAcc  *bool  `gorm:"default:true;index"`
	postgres.Common
}

type Exchange struct {
	ID   string `gorm:"primary_key"`
	Name string `gorm:"unique;not null"`
}
