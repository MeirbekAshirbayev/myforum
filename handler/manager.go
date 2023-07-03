package handler

import (
	"encoding/json"
	"forum/service"
	"log"
	"net/http"
)

type Handler struct {
	log *log.Logger
	svc service.IService
}

func New(l *log.Logger, s service.IService) Handler {
	return Handler{log: l, svc: s}
}

func (h Handler) Home(w http.ResponseWriter, r *http.Request) {
	res, err := h.svc.GetMostLikedTenPost()

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	err = json.NewEncoder(w).Encode(res)
	if err != nil {
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}
}

func (h Handler) TenLikedPost(w http.ResponseWriter, r *http.Request) {
	res, err := h.svc.GetMostLikedTenPost()

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	err = json.NewEncoder(w).Encode(res)
	if err != nil {
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}
}
