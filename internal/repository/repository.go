package repository

import (
	"context"

	"github.com/jmoiron/sqlx"
	"github.com/l-orlov/user-month-expenses/internal/config"
	"github.com/l-orlov/user-month-expenses/internal/models"
	"github.com/l-orlov/user-month-expenses/internal/repository/postgres"
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
		CreateUserExpense(ctx context.Context, expense models.UserExpense) error
		GetUserExpenseByID(ctx context.Context, id uint64) (*models.UserExpense, error)
		UpdateUserExpense(ctx context.Context, expense models.UserExpense) error
		GetAllUserExpenses(ctx context.Context) ([]models.UserExpense, error)
		GetUserExpensesWithParameters(ctx context.Context, params models.UserExpenseParams) ([]models.UserExpense, error)
		GetUserExpensesByCategories(
			ctx context.Context, userID *uint64, size uint16,
		) ([]models.UserExpenseByCategory, error)
		DeleteUserExpense(ctx context.Context, id uint64) error
	}
	Repository struct {
		User
		UserExpense
	}
)

func NewRepository(
	cfg *config.Config, db *sqlx.DB,
) *Repository {
	dbTimeout := cfg.PostgresDB.Timeout.Duration()

	return &Repository{
		User:        postgres.NewUserPostgres(db, dbTimeout),
		UserExpense: postgres.NewUserExpensePostgres(db, dbTimeout),
	}
}
