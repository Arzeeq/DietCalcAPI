package api

import (
	"context"
	"dietcalc/internal/config"
	"dietcalc/internal/logger"
	"dietcalc/internal/service/user"
	postgres "dietcalc/internal/storage/postgres/user"
	"dietcalc/utils"
	"fmt"
	"log/slog"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

func Run() {
	// Setting up logger
	logger := logger.Setup(config.Cfg.Env)
	logger.Info("starting DietCalc", slog.String("env", config.Cfg.Env))

	// Initialize pool of connections
	pool, err := initDB()
	if err != nil {
		logger.Error("failed to init database", utils.ErrAttr(err))
		os.Exit(1)
	}
	defer pool.Close()

	// creating storages
	userStorage := postgres.NewUserStorage(pool)

	// creating handlers
	userHandler := user.NewHandler(userStorage)

	// mounting router
	r := chi.NewRouter()
	r.Mount("/user", user.NewRouter(userHandler))

	// listen port
	logger.Info(fmt.Sprintf("listening %s", config.Cfg.Address))
	http.ListenAndServe(config.Cfg.Address, r)
}

func initDB() (*pgxpool.Pool, error) {
	// form database connection string
	p := &config.Cfg.DBParam
	databaseURL := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", p.User, p.Password, p.Host, p.Port, p.DB)

	// create new pool of connections
	pool, err := pgxpool.New(context.Background(), databaseURL)
	if err != nil {
		return nil, err
	}

	return pool, nil
}
