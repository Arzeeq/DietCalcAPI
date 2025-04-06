package product

import (
	"context"
	"dietcalc/internal/dto"
	"dietcalc/internal/logger"
	"dietcalc/internal/storage"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type Handler struct {
	store storage.ProductStorager
	log   *logger.MyLogger
}

func NewHandler(store storage.ProductStorager, logger *logger.MyLogger) *Handler {
	return &Handler{store: store, log: logger}
}

func (h *Handler) handleCreate(w http.ResponseWriter, r *http.Request) {
	// parse request body
	var productDto dto.Product
	if err := dto.Parse(r.Body, &productDto); err != nil {
		h.log.ReplyHTTPError(w, http.StatusBadRequest, err)
		return
	}

	// check if product is already created
	_, err := h.store.GetById(context.Background(), productDto)
	if err == nil {
		h.log.ReplyHTTPError(w, http.StatusConflict, storage.ErrExists)
		return
	}

	// store product
	if err := h.store.Create(context.Background(), productDto); err != nil {
		h.log.ReplyHTTPError(w, http.StatusInternalServerError, err)
		return
	}

	// write successful creation http status
	w.WriteHeader(http.StatusCreated)
}

func NewRouter(h *Handler) *chi.Mux {
	r := chi.NewRouter()
	r.Post("/", h.handleCreate)
	return r
}
