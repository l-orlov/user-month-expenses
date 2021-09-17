package service

import (
	"context"

	"github.com/l-orlov/user-month-expenses/internal/models"
	"github.com/l-orlov/user-month-expenses/internal/repository"
)

type (
	UserExpenseService struct {
		repo repository.UserExpense
	}
)

func NewUserExpenseService(repo repository.UserExpense) *UserExpenseService {
	return &UserExpenseService{
		repo: repo,
	}
}

func (s *UserExpenseService) CreateUserExpense(ctx context.Context, expense models.UserExpense) error {
	return s.repo.CreateUserExpense(ctx, expense)
}

func (s *UserExpenseService) GetUserExpenseByID(ctx context.Context, id uint64) (*models.UserExpense, error) {
	return s.repo.GetUserExpenseByID(ctx, id)
}

func (s *UserExpenseService) UpdateUserExpense(ctx context.Context, expense models.UserExpense) error {
	return s.repo.UpdateUserExpense(ctx, expense)
}

func (s *UserExpenseService) GetAllUserExpenses(ctx context.Context) ([]models.UserExpense, error) {
	return s.repo.GetAllUserExpenses(ctx)
}

func (s *UserExpenseService) GetUserExpensesWithParameters(ctx context.Context, params models.UserExpenseParams) ([]models.UserExpense, error) {
	return s.repo.GetUserExpensesWithParameters(ctx, params)
}

func (s *UserExpenseService) GetUserExpensesByCategories(
	ctx context.Context, userID *uint64, size uint16,
) ([]models.UserExpenseByCategory, error) {
	return s.repo.GetUserExpensesByCategories(ctx, userID, size)
}

func (s *UserExpenseService) DeleteUserExpense(ctx context.Context, id uint64) error {
	return s.repo.DeleteUserExpense(ctx, id)
}
