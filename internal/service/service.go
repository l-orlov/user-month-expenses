package service

import (
	"context"

	"github.com/l-orlov/user-month-expenses/internal/models"
	"github.com/l-orlov/user-month-expenses/internal/repository"
)

type (
	User interface {
		CreateUser(ctx context.Context, user models.UserToCreate) (uint64, error)
		GetUserByID(ctx context.Context, id uint64) (*models.User, error)
		UpdateUser(ctx context.Context, user models.User) error
		GetAllUsers(ctx context.Context) ([]models.User, error)
		GetUsersWithParameters(ctx context.Context, params models.UserParams) ([]models.User, error)
		DeleteUser(ctx context.Context, id uint64) error
	}
	UserExpense interface {
		CreateUserExpense(ctx context.Context, user models.UserExpense) error
		GetUserExpenseByID(ctx context.Context, id uint64) (*models.UserExpense, error)
		UpdateUserExpense(ctx context.Context, user models.UserExpense) error
		GetAllUserExpenses(ctx context.Context) ([]models.UserExpense, error)
		GetUserExpensesWithParameters(ctx context.Context, params models.UserExpenseParams) ([]models.UserExpense, error)
		GetUserExpensesByCategories(ctx context.Context, size uint16) ([]models.UserExpenseByCategory, error)
		GetUserExpensesByUserIDAndCategories(
			ctx context.Context, userID uint64, size uint16,
		) ([]models.UserExpenseByCategory, error)
		DeleteUserExpense(ctx context.Context, id uint64) error
	}
	Service struct {
		User
		UserExpense
	}
)

func NewService(
	repo *repository.Repository,
) *Service {
	return &Service{
		User:        NewUserService(repo.User),
		UserExpense: NewUserExpenseService(repo.UserExpense),
	}
}
