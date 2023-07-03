package service

import (
	"errors"
	"fmt"
	"forum/entity"
)

func (s *service) CreatePost(postC *entity.PostCreate) error {
	if len(postC.Title) < 1 || len(postC.Content) < 1 {
		return errors.New("Title or content empty")
	}
	post := entity.Post{
		Title:     postC.Title,
		Content:   postC.Content,
		UserID:    postC.UserID,
		CreatedAt: postC.CreatedAt,
	}
	err, n := s.repo.CreatePost(&post)
	if err != nil {
		return err
	}

	err = s.repo.CreateCategory(postC.Cat_Name, n)
	if err != nil {
		return err
	}

	return nil
}

func (s *service) UpdatePost(post *entity.Post) error {
	if len(post.Content) < 1 {
		return errors.New("Title or content empty")
	}

	return s.repo.UpdatePost(post)
}

func (s *service) GetMostLikedTenPost() ([]entity.Post, error) {
	return s.repo.GetMostLikedTenPost()
}

func (s *service) GetPostByID(id int) (*entity.Post, error) {
	fmt.Println(id)
	post, err := s.repo.GetPostByID(id)
	if err != nil {
		fmt.Println("11")
		return nil, err
	}
	post.Comments, err = s.repo.GetCommentByPostID(id)
	if err != nil {
		fmt.Println("22")
		return nil, err
	}
	fmt.Println(post)

	return post, nil
}
