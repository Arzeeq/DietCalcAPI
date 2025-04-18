package model

import (
	"time"

	"github.com/jackc/pgx/v5/pgtype"
)

type User struct {
	Login     string
	Password  string
	Sex       pgtype.Text
	Birthdate pgtype.Date
	Height    pgtype.Numeric
	Purpose   pgtype.Text
	CreatedAt time.Time
}

type Product struct {
	Id        uint32
	Name      pgtype.Text
	Calories  pgtype.Numeric
	Proteins  pgtype.Numeric
	Fats      pgtype.Numeric
	Carbs     pgtype.Numeric
	UserLogin pgtype.Text
}
