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
