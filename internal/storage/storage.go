package storage

import (
	"context"
	"dietcalc/internal/dto"
	"dietcalc/internal/model"
	"errors"
)

var (
	ErrNotFound      = errors.New("requested resource not found")
	ErrExists        = errors.New("resourse already exists")
	ErrWrongPassword = errors.New("wrong password")
)

type UserStorager interface {
	Create(ctx context.Context, dto dto.User) error
	GetAll(ctx context.Context) ([]model.User, error)
	GetByLogin(ctx context.Context, login string) (*model.User, error)
	CheckByLogin(tx context.Context, login string) bool
}

type ProductStorager interface {
	Create(ctx context.Context, dto dto.Product) error
	GetById(ctx context.Context, dto dto.Product) (*model.Product, error)
}
