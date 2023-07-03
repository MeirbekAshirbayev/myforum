package server

import (
	"fmt"
	router "forum/custom_route"
	"forum/handler"
	"forum/middleware"
	"forum/repository"
	"forum/service"
	"log"
	"net/http"
)

type Server struct {
	log     *log.Logger
	handler handler.Handler
}

func NewServer(l *log.Logger) {
	route := router.NewRouter()
	repo := repository.New(l)

	svc := service.New(l, repo)

	handler := handler.New(l, svc)

	server := &Server{
		log:     l,
		handler: handler,
	}

	server.Routers(route)

	http.Handle("/", middleware.VerifyUser(route))
	fmt.Println("http://localhost:4000")
	http.ListenAndServe(":4000", nil)
}
