package service

import (
	"errors"
	"forum/entity"
)

func (s *service) CreateComment(comment *entity.Comment) error {
	if len(comment.Content) < 1 {
		return errors.New("Empty comment string")
	}
	return s.repo.CreateComment(comment)
}

func (s *service) GetCommentByPostID(postID int) ([]entity.Comment, error) {
	return s.repo.GetCommentByPostID(postID)
}
