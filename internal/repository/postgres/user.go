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

type UserPostgres struct {
	db        *sqlx.DB
	dbTimeout time.Duration
}

func NewUserPostgres(db *sqlx.DB, dbTimeout time.Duration) *UserPostgres {
	return &UserPostgres{
		db:        db,
		dbTimeout: dbTimeout,
	}
}

func (r *UserPostgres) CreateUser(ctx context.Context, user models.UserToCreate) (uint64, error) {
	query := fmt.Sprintf(`INSERT INTO %s (gender, age) VALUES ($1, $2) RETURNING id`, userTable)
	var err error

	dbCtx, cancel := context.WithTimeout(ctx, r.dbTimeout)
	defer cancel()

	row := r.db.QueryRowContext(dbCtx, query, &user.Gender, &user.Age)
	if err = row.Err(); err != nil {
		return 0, getDBError(err)
	}

	var id uint64
	if err = row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}

func (r *UserPostgres) GetUserByID(ctx context.Context, id uint64) (*models.User, error) {
	query := fmt.Sprintf(`SELECT id, gender, age FROM %s WHERE id=$1`, userTable)
	var user models.User
	var err error

	dbCtx, cancel := context.WithTimeout(ctx, r.dbTimeout)
	defer cancel()

	if err = r.db.GetContext(dbCtx, &user, query, &id); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}

		return nil, err
	}

	return &user, nil
}

func (r *UserPostgres) UpdateUser(ctx context.Context, user models.User) error {
	query := fmt.Sprintf(`UPDATE %s SET gender = :gender, age = :age WHERE id = :id`, userTable)

	dbCtx, cancel := context.WithTimeout(ctx, r.dbTimeout)
	defer cancel()

	_, err := r.db.NamedExecContext(dbCtx, query, &user)
	if err != nil {
		return getDBError(err)
	}

	return nil
}

func (r *UserPostgres) GetAllUsers(ctx context.Context) ([]models.User, error) {
	query := fmt.Sprintf(`SELECT id, gender, age FROM %s ORDER BY id ASC`, userTable)
	var users []models.User

	dbCtx, cancel := context.WithTimeout(ctx, r.dbTimeout)
	defer cancel()

	if err := r.db.SelectContext(dbCtx, &users, query); err != nil {
		return nil, err
	}

	return users, nil
}

func (r *UserPostgres) GetUsersWithParameters(ctx context.Context, params models.UserParams) ([]models.User, error) {
	query := fmt.Sprintf(`
SELECT id, gender, age FROM %s
WHERE (id = $1 OR $1 is null) AND (gender = $2 OR $2 is null) AND (age = $3 OR $3 is null)
ORDER BY id ASC`, userTable)

	var users []models.User

	dbCtx, cancel := context.WithTimeout(ctx, r.dbTimeout)
	defer cancel()

	err := r.db.SelectContext(dbCtx, &users, query, &params.ID, &params.Gender, &params.Age)

	return users, err
}

func (r *UserPostgres) DeleteUser(ctx context.Context, id uint64) error {
	query := fmt.Sprintf(`DELETE FROM %s WHERE id = $1`, userTable)

	dbCtx, cancel := context.WithTimeout(ctx, r.dbTimeout)
	defer cancel()

	if _, err := r.db.ExecContext(dbCtx, query, &id); err != nil {
		return err
	}

	return nil
}
