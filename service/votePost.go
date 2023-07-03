package service

import "forum/entity"

func (s *service) CreateLikePost(v *entity.VotePost, n int) error {
	b, err := s.repo.CheckLikeByPostIDAndUserID(v.PostID, v.UserID)
	if err != nil {
		return err
	}
	if !b {
		s.repo.DeleteByPostIDAndUserID(v.PostID, v.UserID)
		return s.repo.CreateLikePost(v, n)
	}

	return s.repo.DeleteByPostIDAndUserID(v.PostID, v.UserID)
}

func (s *service) CreateDislikePost(v *entity.VotePost, n int) error {
	b, err := s.repo.CheckDislikeByPostIDAndUserID(v.PostID, v.UserID)
	if err != nil {
		return err
	}
	if !b {
		s.repo.DeleteByPostIDAndUserID(v.PostID, v.UserID)
		return s.repo.CreateDislikePost(v, 1)
	} else {
		return s.repo.DeleteByPostIDAndUserID(v.PostID, v.UserID)
	}
}
