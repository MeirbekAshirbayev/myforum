package service

import (
	"forum/entity"
	"forum/repository"
	"log"
)

type IService interface {
	Signup(user *entity.User) error
	Login(email string, password string) (*entity.User, error)

	CreatePost(postC *entity.PostCreate) error
	UpdatePost(post *entity.Post) error
	GetMostLikedTenPost() ([]entity.Post, error)
	GetPostByID(postID int) (*entity.Post, error)

	CreateComment(comment *entity.Comment) error
	GetCommentByPostID(postID int) ([]entity.Comment, error)

	CreateLikePost(v *entity.VotePost, n int) error
	CreateDislikePost(v *entity.VotePost, n int) error
}

type service struct {
	log  *log.Logger
	repo repository.IDB
}

func New(l *log.Logger, r repository.IDB) IService {
	return &service{log: l, repo: r}
}
