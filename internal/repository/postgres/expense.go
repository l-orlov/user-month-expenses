package postgres

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/l-orlov/user-month-expenses/internal/models"
	"github.com/pkg/errors"
)

type UserExpensePostgres struct {
	db        *sqlx.DB
	dbTimeout time.Duration
}

func NewUserExpensePostgres(db *sqlx.DB, dbTimeout time.Duration) *UserExpensePostgres {
	return &UserExpensePostgres{
		db:        db,
		dbTimeout: dbTimeout,
	}
}

func (r *UserExpensePostgres) CreateUserExpense(ctx context.Context, expense models.UserExpense) error {
	query := fmt.Sprintf(`INSERT INTO %s (user_id, category, amount) VALUES ($1, $2, $3)`, userMonthExpenseTable)

	dbCtx, cancel := context.WithTimeout(ctx, r.dbTimeout)
	defer cancel()

	if _, err := r.db.ExecContext(dbCtx, query, &expense.UserID, &expense.Category, &expense.Amount); err != nil {
		return err
	}

	return nil
}

func (r *UserExpensePostgres) GetUserExpenseByID(ctx context.Context, userID uint64) (*models.UserExpense, error) {
	query := fmt.Sprintf(`SELECT user_id, category, amount FROM %s WHERE user_id=$1`, userMonthExpenseTable)
	var expense models.UserExpense
	var err error

	dbCtx, cancel := context.WithTimeout(ctx, r.dbTimeout)
	defer cancel()

	if err = r.db.GetContext(dbCtx, &expense, query, &userID); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}

		return nil, err
	}

	return &expense, nil
}

func (r *UserExpensePostgres) UpdateUserExpense(ctx context.Context, expense models.UserExpense) error {
	query := fmt.Sprintf(
		`UPDATE %s SET category = :category, amount = :amount WHERE user_id = :user_id`,
		userMonthExpenseTable,
	)

	dbCtx, cancel := context.WithTimeout(ctx, r.dbTimeout)
	defer cancel()

	_, err := r.db.NamedExecContext(dbCtx, query, &expense)
	if err != nil {
		return getDBError(err)
	}

	return nil
}

func (r *UserExpensePostgres) GetAllUserExpenses(ctx context.Context) ([]models.UserExpense, error) {
	query := fmt.Sprintf(`SELECT user_id, category, amount FROM %s ORDER BY user_id ASC`, userMonthExpenseTable)
	var expenses []models.UserExpense

	dbCtx, cancel := context.WithTimeout(ctx, r.dbTimeout)
	defer cancel()

	if err := r.db.SelectContext(dbCtx, &expenses, query); err != nil {
		return nil, err
	}

	return expenses, nil
}

func (r *UserExpensePostgres) GetUserExpensesWithParameters(
	ctx context.Context, params models.UserExpenseParams,
) ([]models.UserExpense, error) {
	query := fmt.Sprintf(`
SELECT user_id, category, amount FROM %s
WHERE (user_id = $1 OR $1 is null) AND (category = $2 OR $2 is null) AND (amount = $3 OR $3 is null)
ORDER BY user_id ASC`, userMonthExpenseTable)

	var expenses []models.UserExpense

	dbCtx, cancel := context.WithTimeout(ctx, r.dbTimeout)
	defer cancel()

	err := r.db.SelectContext(dbCtx, &expenses, query, &params.UserID, &params.Category, &params.Amount)

	return expenses, err
}

func (r *UserExpensePostgres) GetUserExpensesByCategories(
	ctx context.Context, size uint16,
) ([]models.UserExpenseByCategory, error) {
	query := fmt.Sprintf(`
SELECT category, sum(amount) as amount FROM %s
GROUP BY category
ORDER BY amount DESC
LIMIT $1`, userMonthExpenseTable)

	var expenses []models.UserExpenseByCategory

	dbCtx, cancel := context.WithTimeout(ctx, r.dbTimeout)
	defer cancel()

	err := r.db.SelectContext(dbCtx, &expenses, query, &size)

	return expenses, err
}

func (r *UserExpensePostgres) GetUserExpensesByUserIDAndCategories(
	ctx context.Context, userID uint64, size uint16,
) ([]models.UserExpenseByCategory, error) {
	query := fmt.Sprintf(`
SELECT category, sum(amount) as amount FROM %s
WHERE user_id = $1
GROUP BY category
ORDER BY amount DESC
LIMIT $2`, userMonthExpenseTable)

	var expenses []models.UserExpenseByCategory

	dbCtx, cancel := context.WithTimeout(ctx, r.dbTimeout)
	defer cancel()

	err := r.db.SelectContext(dbCtx, &expenses, query, &userID, &size)

	return expenses, err
}

func (r *UserExpensePostgres) DeleteUserExpense(ctx context.Context, userID uint64) error {
	query := fmt.Sprintf(`DELETE FROM %s WHERE user_id = $1`, userMonthExpenseTable)

	dbCtx, cancel := context.WithTimeout(ctx, r.dbTimeout)
	defer cancel()

	if _, err := r.db.ExecContext(dbCtx, query, &userID); err != nil {
		return err
	}

	return nil
}
