package user

import (
	"context"
	"dietcalc/internal/config"
	"dietcalc/internal/dto"
	"dietcalc/internal/logger"
	"dietcalc/internal/service/auth"
	"dietcalc/internal/storage"
	"encoding/json"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
)

type Handler struct {
	store storage.UserStorager
	log   *logger.MyLogger
}

func NewHandler(store storage.UserStorager, log *logger.MyLogger) *Handler {
	return &Handler{store: store, log: log}
}

func (h *Handler) handleCreate(w http.ResponseWriter, r *http.Request) {
	// parse request body
	var userDto dto.User
	if err := dto.Parse(r.Body, &userDto); err != nil {
		h.log.ReplyHTTPError(w, http.StatusBadRequest, err)
		return
	}

	// check if user is already created
	if h.store.CheckByLogin(context.Background(), userDto.Login) {
		h.log.ReplyHTTPError(w, http.StatusConflict, storage.ErrExists)
		return
	}

	// set creation time and hash user password
	userDto.CreatedAt = time.Now()
	pass, err := auth.HashPassword(userDto.Password)
	if err != nil {
		h.log.ReplyHTTPError(w, http.StatusInternalServerError, err)
		return
	}
	userDto.Password = pass

	// store user
	if err := h.store.Create(context.Background(), userDto); err != nil {
		h.log.ReplyHTTPError(w, http.StatusInternalServerError, err)
		return
	}

	// write successful creation http status
	w.WriteHeader(http.StatusCreated)
}

func (h *Handler) handleLogin(w http.ResponseWriter, r *http.Request) {
	// parse request body
	var userDto dto.User

	if err := dto.Parse(r.Body, &userDto); err != nil {
		h.log.ReplyHTTPError(w, http.StatusBadRequest, err)
		return
	}

	// get user from storage
	user, err := h.store.GetByLogin(context.Background(), userDto.Login)
	if err != nil {
		h.log.ReplyHTTPError(w, http.StatusNotFound, storage.ErrNotFound)
		return
	}

	// compare password
	if !auth.ComparePasswords(user.Password, userDto.Password) {
		h.log.ReplyHTTPError(w, http.StatusBadRequest, storage.ErrWrongPassword)
		return
	}

	// create JWT
	token, err := auth.CreateJWT(
		userDto.Login,
		auth.JWTParams{
			Secret:   []byte(config.Cfg.JWTSecret),
			Duration: config.Cfg.JWTDuration,
		},
	)

	if err != nil {
		h.log.ReplyHTTPError(w, http.StatusInternalServerError, err)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	if err = json.NewEncoder(w).Encode(map[string]string{"jwt_token": token}); err != nil {
		h.log.Error("failed to write jwt_token", logger.ErrAttr(err))
	}
}

func NewRouter(h *Handler) *chi.Mux {
	r := chi.NewRouter()
	r.Post("/", h.handleCreate)
	r.Post("/login", h.handleLogin)
	return r
}
