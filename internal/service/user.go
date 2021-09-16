package service

import (
	"context"

	"github.com/l-orlov/user-month-expenses/internal/models"
	"github.com/l-orlov/user-month-expenses/internal/repository"
)

type (
	UserService struct {
		repo repository.User
	}
)

func NewUserService(repo repository.User) *UserService {
	return &UserService{
		repo: repo,
	}
}

func (s *UserService) CreateUser(ctx context.Context, user models.UserToCreate) (uint64, error) {
	return s.repo.CreateUser(ctx, user)
}

func (s *UserService) GetUserByID(ctx context.Context, id uint64) (*models.User, error) {
	return s.repo.GetUserByID(ctx, id)
}

func (s *UserService) UpdateUser(ctx context.Context, user models.User) error {
	return s.repo.UpdateUser(ctx, user)
}

func (s *UserService) GetAllUsers(ctx context.Context) ([]models.User, error) {
	return s.repo.GetAllUsers(ctx)
}

func (s *UserService) GetUsersWithParameters(ctx context.Context, params models.UserParams) ([]models.User, error) {
	return s.repo.GetUsersWithParameters(ctx, params)
}

func (s *UserService) DeleteUser(ctx context.Context, id uint64) error {
	return s.repo.DeleteUser(ctx, id)
}
