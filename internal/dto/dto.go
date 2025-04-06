package dto

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
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

func Parse(r io.Reader, payload any) error {
	if r == nil {
		return fmt.Errorf("parsing from nil reader")
	}

	return json.NewDecoder(r).Decode(payload)
}

func Write(w http.ResponseWriter, status int, dto any) error {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(dto)
}
