package repository

import (
	"database/sql"
	"forum/entity"
)

func (r repo) CreateLikeComment(vote *entity.VoteComment) error {
	stmt, err := r.db.Prepare(`
		INSERT INTO VoteComment (CommentID, UserID, Like, Dislike)
		VALUES (?, ?, ?, 0)
		ON CONFLICT (CommentID, UserID) DO UPDATE SET Like = ?
	`)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(vote.CommentID, vote.UserID, 1, 1)
	if err != nil {
		return err
	}

	return nil
}

func (r repo) CreateDislikeComment(vote *entity.VoteComment) error {
	stmt, err := r.db.Prepare(`
		INSERT INTO VoteComment (CommentID, UserID, Like, Dislike)
		VALUES (?, ?, 0, ?)
		ON CONFLICT (CommentID, UserID) DO UPDATE SET Dislike = ?
	`)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(vote.CommentID, vote.UserID, 1, 1)
	if err != nil {
		return err
	}

	return nil
}

func (r repo) CountLikeComment(commentID int) (int, error) {
	stmt, err := r.db.Prepare(`
		SELECT COUNT(*) FROM VoteComment
		WHERE CommentID = ? AND Like = 1
	`)
	if err != nil {
		return 0, err
	}
	defer stmt.Close()

	var count int
	err = stmt.QueryRow(commentID).Scan(&count)
	if err != nil {
		return 0, err
	}

	return count, nil
}

func (r repo) CountDislikeComment(commentID int) (int, error) {
	stmt, err := r.db.Prepare(`
		SELECT COUNT(*) FROM VoteComment
		WHERE CommentID = ? AND Dislike = 1
	`)
	if err != nil {
		return 0, err
	}
	defer stmt.Close()

	var count int
	err = stmt.QueryRow(commentID).Scan(&count)
	if err != nil {
		return 0, err
	}

	return count, nil
}

func (r repo) CheckLikeByCommentIDAndUserID(commentID, userID int) (bool, error) {
	stmt, err := r.db.Prepare(`
		SELECT Like FROM VoteComment
		WHERE CommentID = ? AND UserID = ?
	`)
	if err != nil {
		return false, err
	}
	defer stmt.Close()

	var like bool
	err = stmt.QueryRow(commentID, userID).Scan(&like)
	if err != nil {
		if err == sql.ErrNoRows {
			return false, nil
		}
		return false, err
	}

	return like, nil
}

func (r repo) CheckDislikeByCommentIDAndUserID(commentID, userID int) (bool, error) {
	stmt, err := r.db.Prepare(`
		SELECT Dislike FROM VoteComment
		WHERE CommentID = ? AND UserID = ?
	`)
	if err != nil {
		return false, err
	}
	defer stmt.Close()

	var dislike bool
	err = stmt.QueryRow(commentID, userID).Scan(&dislike)
	if err != nil {
		if err == sql.ErrNoRows {
			return false, nil
		}
		return false, err
	}

	return dislike, nil
}
