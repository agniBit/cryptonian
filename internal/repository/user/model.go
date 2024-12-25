package user

import "github.com/agniBit/cryptonian/internal/storage/postgres"

type User struct {
	ID              string `gorm:"primary_key"`
	Name            string `gorm:"unique;not null"`
	Email           string `gorm:"unique;not null"`
	PhoneNumber     string `gorm:"unique;not null"`
	IsActive        *bool  `gorm:"default:true;index"`
	IsEmailVerified *bool  `gorm:"default:false"`
	postgres.Common
}
