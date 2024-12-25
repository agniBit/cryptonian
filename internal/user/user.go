package user

import (
	"context"

	"github.com/agniBit/cryptonian/internal/logger"
	userRepo "github.com/agniBit/cryptonian/internal/repository/user"
	"github.com/agniBit/cryptonian/model/cfg"
	"github.com/agniBit/cryptonian/model/user"
	"github.com/jinzhu/copier"
)

type ServiceInterface interface {
	CreateUser(ctx context.Context, user *user.User) (*user.User, error)
	GetUser(ctx context.Context, id string) (*user.User, error)
	GetUserByPhoneNumber(ctx context.Context, phoneNumber string) (*user.User, error)
}

type service struct {
	cfg      *cfg.Config
	userRepo userRepo.RepositoryInterface
}

func NewUserService(cfg *cfg.Config, userRepo userRepo.RepositoryInterface) ServiceInterface {
	return &service{
		cfg:      cfg,
		userRepo: userRepo,
	}
}

func (s *service) CreateUser(ctx context.Context, user *user.User) (*user.User, error) {
	uRepo := &userRepo.User{}
	err := copier.Copy(uRepo, user)
	if err != nil {
		return nil, err
	}
	err = s.userRepo.CreateUser(ctx, uRepo)
	if err != nil {
		logger.Error(ctx, "error in creating user", err, nil)
		return nil, err
	}

	user.ID = uRepo.ID
	return user, nil
}

func (s *service) GetUser(ctx context.Context, id string) (*user.User, error) {
	uRepo, err := s.userRepo.GetUser(ctx, id)
	if err != nil {
		logger.Error(ctx, "error in getting user", err, nil)
		return nil, err
	}

	u := &user.User{}
	err = copier.Copy(u, uRepo)
	if err != nil {
		logger.Error(ctx, "error in copying user", err, nil)
		return nil, err
	}

	return u, nil
}

func (s *service) GetUserByPhoneNumber(ctx context.Context, phoneNumber string) (*user.User, error) {
	uRepo, err := s.userRepo.GetUserByPhoneNumber(ctx, phoneNumber)
	if err != nil {
		logger.Error(ctx, "error in getting user by phone number", err, nil)
		return nil, err
	}

	u := &user.User{}
	err = copier.Copy(u, uRepo)
	if err != nil {
		logger.Error(ctx, "error in copying user", err, nil)
		return nil, err
	}

	return u, nil
}
