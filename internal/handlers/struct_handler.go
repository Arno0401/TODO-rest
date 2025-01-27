package handler

import "arno/internal/repository"

type Handler struct {
	rep *repository.Repository
}

func NewHandler(rep *repository.Repository) *Handler {
	return &Handler{
		rep: rep,
	}
}
