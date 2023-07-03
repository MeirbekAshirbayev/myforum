package handler

import (
	"encoding/json"
	"forum/entity"
	"net/http"
	"sync"
)

func (h Handler) SignUp(w http.ResponseWriter, r *http.Request) {
	user := entity.User{}
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}
	err = h.svc.Signup(&user)
	if err != nil {
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	cookie, err := r.Cookie("Session")
	if err != nil {
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	var mu sync.Mutex
	mu.Lock()
	entity.SessionMap[cookie.Value] = user
	mu.Unlock()
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("kostym"))
}

func (h Handler) Login(w http.ResponseWriter, r *http.Request) {
	user := entity.User{}
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}
	user1, err := h.svc.Login(user.Email, user.Password)
	if err != nil {
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	cookie, err := r.Cookie("Session")
	if err != nil {
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	var mu sync.Mutex
	mu.Lock()
	entity.SessionMap[cookie.Value] = *user1
	mu.Unlock()
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("kostym"))
}

func (h Handler) Logout(w http.ResponseWriter, r *http.Request) {
	cookie, _ := r.Cookie("Session")

	delete(entity.SessionMap, cookie.Value)

	cookie = &http.Cookie{
		Name:   "Session",
		Value:  "",
		MaxAge: -1,
	}
	http.SetCookie(w, cookie)

	http.Redirect(w, r, "/home", http.StatusSeeOther)
}
