package user

import (
	"context"
	"dietcalc/internal/config"
	"dietcalc/internal/dto"
	"dietcalc/internal/service/auth"
	"dietcalc/internal/storage"
	"dietcalc/utils"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
)

type Handler struct {
	store storage.UserStorager
}

func NewHandler(store storage.UserStorager) *Handler {
	return &Handler{store: store}
}

func (h *Handler) handleCreate(w http.ResponseWriter, r *http.Request) {
	// parse request body
	var dto dto.User
	if err := utils.ParseJSON(r, &dto); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	// check if user is already created
	_, err := h.store.GetByLogin(context.Background(), dto)
	if err == nil {
		utils.WriteError(w, http.StatusConflict, storage.ErrExists)
		return
	}

	// set creation time and hash user password
	dto.CreatedAt = time.Now()
	pass, err := auth.HashPassword(dto.Password)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}
	dto.Password = pass

	// store user
	if err := h.store.Create(context.Background(), dto); err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	// write successful creation http status
	w.WriteHeader(http.StatusCreated)
}

func (h *Handler) handleLogin(w http.ResponseWriter, r *http.Request) {
	// parse request body
	var dto dto.User
	if err := utils.ParseJSON(r, &dto); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	// get user from storage
	user, err := h.store.GetByLogin(context.Background(), dto)
	if err != nil {
		utils.WriteError(w, http.StatusNotFound, storage.ErrNotFound)
		return
	}

	// compare password
	if !auth.ComparePasswords(user.Password, []byte(dto.Password)) {
		utils.WriteError(w, http.StatusBadRequest, storage.ErrWrongPassword)
		return
	}

	// create JWT
	token, err := auth.CreateJWT([]byte(config.Cfg.JWTSecret), dto.Login)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSON(w, http.StatusOK, map[string]string{"token": token})
}

func (h *Handler) handleGetAll(w http.ResponseWriter, r *http.Request) {
	users, err := h.store.GetAll(context.Background())
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	utils.WriteJSON(w, http.StatusOK, users)
}

func NewRouter(h *Handler) *chi.Mux {
	r := chi.NewRouter()
	r.Post("/", h.handleCreate)
	r.Post("/login", h.handleLogin)
	r.Get("/", h.handleGetAll)
	return r
}
