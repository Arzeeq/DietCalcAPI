package postgres

import (
	"context"
	"dietcalc/internal/dto"
	"dietcalc/internal/model"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type ProductStorage struct {
	pool *pgxpool.Pool
}

func NewProductStorage(pool *pgxpool.Pool) *ProductStorage {
	return &ProductStorage{pool: pool}
}

func (s *ProductStorage) Create(ctx context.Context, dto dto.Product) error {
	query :=
		`INSERT INTO "products"
	("name", "calories", "proteins", "fats", "carbs", "user_login")
	VALUES (@name, @calories, @proteins, @fats, @carbs, @user_login)`

	args := pgx.NamedArgs{
		"name":       dto.Name,
		"calories":   dto.Calories,
		"proteins":   dto.Proteins,
		"fats":       dto.Fats,
		"carbs":      dto.Carbs,
		"user_login": dto.UserLogin,
	}

	_, err := s.pool.Exec(ctx, query, args)

	return err
}

func (s *ProductStorage) GetById(ctx context.Context, dto dto.Product) (*model.Product, error) {
	query := `SELECT * FROM "product" WHERE id = @id;`

	args := pgx.NamedArgs{
		"id": dto.Id,
	}

	rows, err := s.pool.Query(ctx, query, args)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	product, err := pgx.CollectExactlyOneRow(rows, pgx.RowToStructByName[model.Product])

	if err != nil {
		return nil, err
	}

	return &product, nil
}
