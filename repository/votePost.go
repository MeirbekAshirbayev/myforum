package repository

import (
	"database/sql"
	"fmt"
	"forum/entity"
)

func (r repo) CreateLikePost(vote *entity.VotePost, n int) error {
	stmt, err := r.db.Prepare(`
		INSERT INTO VotePost (PostID, UserID, Like, Dislike)
		VALUES (?, ?, ?, ?)
	`)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(vote.PostID, vote.UserID, n, 0)
	if err != nil {
		return err
	}

	return nil
}

func (r repo) CreateDislikePost(vote *entity.VotePost, n int) error {
	stmt, err := r.db.Prepare(`
		INSERT INTO VotePost (PostID, UserID, Like, Dislike)
		VALUES (?, ?, ?, ?)
	`)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(vote.PostID, vote.UserID, 0, n)
	if err != nil {
		return err
	}

	return nil
}

func (r repo) CountLikePost(postID int) (int, error) {
	stmt, err := r.db.Prepare(`
		SELECT COUNT(*) FROM VotePost
		WHERE PostID = ? AND Like = 1
	`)
	if err != nil {
		return 0, err
	}
	defer stmt.Close()

	var count int
	err = stmt.QueryRow(postID).Scan(&count)
	if err != nil {
		return 0, err
	}

	return count, nil
}

func (r repo) CountDislikePost(postID int) (int, error) {
	stmt, err := r.db.Prepare(`
		SELECT COUNT(*) FROM VotePost
		WHERE PostID = ? AND Dislike = 1
	`)
	if err != nil {
		return 0, err
	}
	defer stmt.Close()

	var count int
	err = stmt.QueryRow(postID).Scan(&count)
	if err != nil {
		return 0, err
	}

	return count, nil
}

func (r repo) CheckLikeByPostIDAndUserID(postID, userID int) (bool, error) {
	stmt, err := r.db.Prepare(`
		SELECT Like FROM VotePost
		WHERE PostID = ? AND UserID = ?
	`)
	if err != nil {
		return false, err
	}
	defer stmt.Close()

	var like bool
	err = stmt.QueryRow(postID, userID).Scan(&like)
	if err != nil {
		if err == sql.ErrNoRows {
			return false, nil
		}
		return false, err
	}

	return like, nil
}

func (r repo) CheckDislikeByPostIDAndUserID(postID, userID int) (bool, error) {
	stmt, err := r.db.Prepare(`
		SELECT Dislike FROM VotePost
		WHERE PostID = ? AND UserID = ?
	`)
	if err != nil {
		return false, err
	}
	defer stmt.Close()

	var dislike bool
	err = stmt.QueryRow(postID, userID).Scan(&dislike)
	if err != nil {
		if err == sql.ErrNoRows {
			return false, nil
		}
		return false, err
	}

	return dislike, nil
}

func (r repo) DeleteByPostIDAndUserID(postID, userID int) error {
	stmt, err := r.db.Prepare("DELETE FROM VotePost WHERE UserID = ? AND PostID = ?")
	if err != nil {
		fmt.Println("Error preparing the statement:", err)
		return err
	}
	_, err = stmt.Exec(userID, postID) // Replace 1 and 2 with the actual UserID and PostID values
	if err != nil {
		fmt.Println("Error executing the statement:", err)
		return err
	}
	return nil
}
