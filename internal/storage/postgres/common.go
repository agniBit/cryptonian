package postgres

import (
	"database/sql/driver"
	"encoding/json"
	"time"

	"gorm.io/gorm"
)

type Common struct {
	CreatedAt time.Time      `gorm:"autoCreateTime;not null;default:now()"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime;not null;default:now()"`
	DeletedAt gorm.DeletedAt `gorm:"index"` // Soft delete
}

type JSONB map[string]interface{}

func (j JSONB) Value() (driver.Value, error) {
	valueString, err := json.Marshal(j)
	return string(valueString), err
}

func (j *JSONB) Scan(value interface{}) error {
	if err := json.Unmarshal(value.([]uint8), &j); err != nil {
		return err
	}
	return nil
}
