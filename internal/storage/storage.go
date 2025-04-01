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
	GetByLogin(ctx context.Context, dto dto.User) (*model.User, error)
}

type ProductStorager interface {
}
