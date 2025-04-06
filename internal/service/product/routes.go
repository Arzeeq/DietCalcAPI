package product

import (
	"context"
	"dietcalc/internal/dto"
	"dietcalc/internal/storage"
	"dietcalc/utils"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type Handler struct {
	store storage.ProductStorager
}

func NewHandler(store storage.ProductStorager) *Handler {
	return &Handler{store: store}
}

func (h *Handler) handleCreate(w http.ResponseWriter, r *http.Request) {
	// parse request body
	var dto dto.Product
	if err := utils.ParseJSON(r, &dto); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	// check if product is already created
	_, err := h.store.GetById(context.Background(), dto)
	if err == nil {
		utils.WriteError(w, http.StatusConflict, storage.ErrExists)
		return
	}

	// store product
	if err := h.store.Create(context.Background(), dto); err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	// write successful creation http status
	w.WriteHeader(http.StatusCreated)
}

// func (h *Handler) handleGetAll(w http.ResponseWriter, r *http.Request) {
// 	users, err := h.store.GetAll(context.Background())
// 	if err != nil {
// 		utils.WriteError(w, http.StatusBadRequest, err)
// 		return
// 	}

// 	utils.WriteJSON(w, http.StatusOK, users)
// }

func NewRouter(h *Handler) *chi.Mux {
	r := chi.NewRouter()
	r.Post("/", h.handleCreate)
	return r
}
