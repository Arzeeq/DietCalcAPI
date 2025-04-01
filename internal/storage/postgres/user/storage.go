package postgres

import (
	"context"
	"dietcalc/internal/dto"
	"dietcalc/internal/model"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type UserStorage struct {
	pool *pgxpool.Pool
}

func NewUserStorage(pool *pgxpool.Pool) *UserStorage {
	return &UserStorage{pool: pool}
}

func (s *UserStorage) Create(ctx context.Context, dto dto.User) error {
	query :=
		`INSERT INTO "users"
	("login", "password", "sex", "birthdate", "height", "purpose", "created_at")
	VALUES (@login, @password, @sex, @birthdate, @height, @purpose, @created_at)`

	args := pgx.NamedArgs{
		"login":      dto.Login,
		"password":   dto.Password,
		"sex":        dto.Sex,
		"birthdate":  dto.Birthdate,
		"height":     dto.Height,
		"purpose":    dto.Purpose,
		"created_at": dto.CreatedAt,
	}

	_, err := s.pool.Exec(ctx, query, args)

	return err
}

func (s *UserStorage) GetAll(ctx context.Context) ([]model.User, error) {
	query := `SELECT * FROM "users"`

	rows, err := s.pool.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	users, err := pgx.CollectRows(rows, pgx.RowToStructByName[model.User])
	if err != nil {
		return nil, err
	}

	return users, nil
}

func (s *UserStorage) GetByLogin(ctx context.Context, dto dto.User) (*model.User, error) {
	query := `SELECT * FROM "users" WHERE login = @login;`

	args := pgx.NamedArgs{
		"login": dto.Login,
	}

	rows, err := s.pool.Query(ctx, query, args)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	user, err := pgx.CollectExactlyOneRow(rows, pgx.RowToStructByName[model.User])

	if err != nil {
		return nil, err
	}

	return &user, nil
}
