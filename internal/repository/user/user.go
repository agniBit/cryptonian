package user

import (
	"context"

	"github.com/agniBit/cryptonian/internal/logger"
	"github.com/agniBit/cryptonian/internal/storage/postgres"
	"github.com/agniBit/cryptonian/model/cfg"
	"gorm.io/gorm"
)

type RepositoryInterface interface {
	CreateUser(ctx context.Context, user *User) error
	GetUser(ctx context.Context, id string) (*User, error)
	GetUserByPhoneNumber(ctx context.Context, phoneNumber string) (*User, error)
}

type service struct {
	cfg *cfg.Config
	gdb *gorm.DB
}

func NewUserService(cfg *cfg.Config, gdb *gorm.DB) *service {
	return &service{
		cfg: cfg,
		gdb: gdb,
	}
}

func (s *service) Migrate() error {
	err := s.gdb.AutoMigrate(&User{})
	if err != nil {
		return err
	}

	// create sequence for User table
	err = s.gdb.Exec("DO $$ BEGIN CREATE SEQUENCE IF NOT EXISTS user_id_seq; EXCEPTION WHEN duplicate_table THEN END $$;").Error
	if err != nil {
		logger.Fatal(context.Background(), "Error while creating sequence", err, map[string]interface{}{"sequence": "user_id_seq"})
	}

	err = s.gdb.Exec("ALTER TABLE users ADD COLUMN IF NOT EXISTS id VARCHAR(10) GENERATED ALWAYS AS ('USR-' || nextval('user_id_seq')::text) STORED").Error
	if err != nil {
		logger.Fatal(context.Background(), "Error while adding column to table", err, map[string]interface{}{"table": "User"})
	}

	return nil
}

func (s *service) CreateUser(ctx context.Context, user *User) error {
	gdb := postgres.GetRepositoryFromContext(ctx, s.gdb)

	return gdb.Create(user).Error
}

func (s *service) GetUser(ctx context.Context, id string) (*User, error) {
	gdb := postgres.GetRepositoryFromContext(ctx, s.gdb)

	var user User
	err := gdb.Where("id = ?", id).First(&user).Error
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (s *service) GetUserByPhoneNumber(ctx context.Context, phoneNumber string) (*User, error) {
	gdb := postgres.GetRepositoryFromContext(ctx, s.gdb)

	var user User
	err := gdb.Where("phone_number = ?", phoneNumber).First(&user).Error
	if err != nil {
		return nil, err
	}

	return &user, nil
}
