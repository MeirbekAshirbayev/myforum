package handler

import (
	"encoding/json"
	"fmt"
	"forum/entity"
	"net/http"
	"strconv"
)

func (h Handler) CreatePost(w http.ResponseWriter, r *http.Request) {
	postC := entity.PostCreate{}

	err := json.NewDecoder(r.Body).Decode(&postC)
	if err != nil {
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	err = h.svc.CreatePost(&postC)
	if err != nil {
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}
}

func (h Handler) UpdatePost(w http.ResponseWriter, r *http.Request) {
	post := entity.Post{}
	err := json.NewDecoder(r.Body).Decode(&post)
	if err != nil {
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	err = h.svc.UpdatePost(&post)
	if err != nil {
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(post)
}

func (h Handler) CreateLikePost(w http.ResponseWriter, r *http.Request) {
	v := entity.VotePost{}
	err := json.NewDecoder(r.Body).Decode(&v)
	if err != nil {
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	if v.Like {
		err := h.svc.CreateLikePost(&v, 1)
		if err != nil {
			fmt.Println(err)
			return
		}
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(v)
}

func (h Handler) CreateDislikePost(w http.ResponseWriter, r *http.Request) {
	v := entity.VotePost{}
	err := json.NewDecoder(r.Body).Decode(&v)
	if err != nil {
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	if v.Dislike {
		err := h.svc.CreateDislikePost(&v, 1)
		if err != nil {
			fmt.Println(err)
			return
		}
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(v)
}

func (h Handler) GetPostByID(w http.ResponseWriter, r *http.Request) {
	n, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}
	post, err := h.svc.GetPostByID(n)
	if err != nil {
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	err = json.NewEncoder(w).Encode(post)
	if err != nil {
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
}
