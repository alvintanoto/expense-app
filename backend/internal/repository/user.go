package repository

import (
	"database/sql"
	"errors"
	"expense_app/internal/entity"
	"expense_app/internal/util/logger"

	"github.com/lib/pq"
)

type (
	UserRepository interface {
		CreateUser(*entity.User) error
		GetActiveUserByUsername(username string) (*entity.User, error)
	}

	implUser struct {
		db     *sql.DB
		logger logger.Logger
	}
)

func NewUserRepository(logger logger.Logger, db *sql.DB) (UserRepository, error) {
	return &implUser{logger: logger, db: db}, nil
}

func (i *implUser) CreateUser(user *entity.User) (err error) {
	query := `INSERT INTO 
		public.users (id, username, email, password, is_active, created_by)
		VALUES ($1, $2, $3, $4, $5, $6)`

	args := []interface{}{
		user.ID,
		user.Username,
		user.Email,
		user.Password,
		user.IsActive,
		user.CreatedBy,
	}

	row := i.db.QueryRow(query, args...)
	if row.Err() != nil {
		i.logger.Errorf("error inserting user: %+v", row.Err().Error())

		switch e := row.Err().(type) {
		case *pq.Error:
			switch e.Code {
			case "23505":
				return ErrorConstraintViolation
			}
		}

		return row.Err()
	}

	return nil
}

func (i *implUser) GetActiveUserByUsername(username string) (user *entity.User, err error) {
	user = new(entity.User)
	query := `SELECT id, username, email, password, is_active, 
				created_at, created_by, updated_at, updated_by, is_deleted 
				FROM public.users
				WHERE username = $1 AND is_deleted = false AND is_active = true`

	err = i.db.QueryRow(query, username).Scan(
		&user.ID,
		&user.Username,
		&user.Email,
		&user.Password,
		&user.IsActive,
		&user.CreatedAt,
		&user.CreatedBy,
		&user.UpdatedAt,
		&user.UpdatedBy,
		&user.IsDeleted,
	)
	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			i.logger.Error("no record found")
			return nil, ErrorRecordNotFound
		default:
			i.logger.Error("scan row into struct error")
			return nil, err
		}
	}

	return user, nil
}
