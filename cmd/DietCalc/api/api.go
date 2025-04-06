package api

import (
	"context"
	"dietcalc/internal/config"
	"dietcalc/internal/logger"
	"dietcalc/internal/service/product"
	"dietcalc/internal/service/user"
	"dietcalc/internal/storage/postgres"
	"fmt"
	"log/slog"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

func Run() {
	// Setting up logger
	log := logger.New(config.Cfg.Env)
	log.Info("starting DietCalc", slog.String("env", config.Cfg.Env))
	defer func() {
		log.Info("DietCalc finished")
	}()

	// Initialize pool of connections
	pool, err := initDB()
	if err != nil {
		log.Error("failed to init database", logger.ErrAttr(err))
		os.Exit(1)
	}
	defer pool.Close()

	// creating storages
	userStorage := postgres.NewUserStorage(pool)
	productStorage := postgres.NewProductStorage(pool)

	// creating handlers
	userHandler := user.NewHandler(userStorage, log)
	productHandler := product.NewHandler(productStorage, log)

	// mounting router
	r := chi.NewRouter()
	r.Route(config.Cfg.Prefix, func(router chi.Router) {
		router.Mount("/user", user.NewRouter(userHandler))
		router.Mount("/product", product.NewRouter(productHandler))
	})

	// listen port
	log.Info(fmt.Sprintf("listening %s%s", config.Cfg.Address, config.Cfg.Prefix))
	err = http.ListenAndServe(config.Cfg.Address, r)
	log.Error("application DietCalc finished with error", logger.ErrAttr(err))
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
