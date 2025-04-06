package dto

import (
	"time"

	"github.com/jackc/pgx/v5/pgtype"
)

type User struct {
	Login     string         `json:"login"`
	Password  string         `json:"password"`
	Sex       pgtype.Text    `json:"sex"`
	Birthdate pgtype.Date    `json:"birthdate"`
	Height    pgtype.Numeric `json:"height"`
	Purpose   pgtype.Text    `json:"purpose"`
	CreatedAt time.Time      `json:"created_at"`
}

type Product struct {
	Id        uint32         `json:"id"`
	Name      pgtype.Text    `json:"name"`
	Calories  pgtype.Numeric `json:"calories"`
	Proteins  pgtype.Numeric `json:"proteins"`
	Fats      pgtype.Numeric `json:"fats"`
	Carbs     pgtype.Numeric `json:"carbs"`
	UserLogin pgtype.Text    `json:"user_login"`
	JWTToken  string         `json:"jwt_token"`
}
