package handler

import (
	"encoding/json"
	"forum/entity"
	"net/http"
)

func (h Handler) CreateComment(w http.ResponseWriter, r *http.Request) {
	req := entity.Comment{}
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}
	err = h.svc.CreateComment(&req)
	if err != nil {
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("comment jazdym"))
}
