package exchange

import (
	"gorm.io/gorm"

	"github.com/agniBit/cryptonian/internal/logger"
	"github.com/agniBit/cryptonian/internal/utils"
)

const entryptedTextPrefix = "enc:"

func (a *Account) BeforeSave(tx *gorm.DB) (err error) {
	if a.ApiKey != "" {
		aks, err := utils.Encrypt(a.ApiKey)
		if err == nil {
			a.ApiKey = entryptedTextPrefix + aks
		} else {
			logger.Error(tx.Statement.Context, "error in encrypting text", err, nil)
		}
	}

	if a.SecretKey != "" {
		ske, err := utils.Encrypt(a.SecretKey)
		if err == nil {
			a.SecretKey = entryptedTextPrefix + ske
		} else {
			logger.Error(tx.Statement.Context, "error in encrypting text", err, nil)
		}
	}

	return
}

func (a *Account) AfterFind(tx *gorm.DB) (err error) {
	if len(a.ApiKey) > len(entryptedTextPrefix) && a.ApiKey[:len(entryptedTextPrefix)] == entryptedTextPrefix {
		akd, err := utils.Decrypt(a.ApiKey[len(entryptedTextPrefix):])
		if err == nil {
			a.ApiKey = akd
		} else {
			logger.Error(tx.Statement.Context, "error in decrypting text", err, nil)
			return err
		}
	}

	if len(a.SecretKey) > len(entryptedTextPrefix) && a.SecretKey[:len(entryptedTextPrefix)] == entryptedTextPrefix {
		skd, err := utils.Decrypt(a.SecretKey[len(entryptedTextPrefix):])
		if err == nil {
			a.SecretKey = skd
		} else {
			logger.Error(tx.Statement.Context, "error in decrypting text", err, nil)
			return err
		}
	}

	return
}
