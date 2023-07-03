package repository

import (
	"database/sql"
	"forum/entity"
	"time"
)

func (r repo) CreateComment(comment *entity.Comment) error {
	stmt, err := r.db.Prepare(`
		INSERT INTO Comment (Content, PostID, UserID, CreatedAt)
		VALUES (?, ?, ?, ?)
	`)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(
		comment.Content,
		comment.PostID,
		comment.UserID,
		time.Now(),
	)
	if err != nil {
		return err
	}

	return nil
}

func (r repo) DeleteByID(commentID int) error {
	stmt, err := r.db.Prepare(`
		DELETE FROM Comment
		WHERE ID = ?
	`)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(commentID)
	if err != nil {
		return err
	}

	return nil
}

func (r repo) DeleteByPostID(postID int) error {
	stmt, err := r.db.Prepare(`
		DELETE FROM Comment
		WHERE PostID = ?
	`)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(postID)
	if err != nil {
		return err
	}

	return nil
}

func (r repo) GetCommentByPostID(postID int) ([]entity.Comment, error) {
	stmt, err := r.db.Prepare(`
		SELECT *
		FROM Comment
		WHERE PostID = ?
	`)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	rows, err := stmt.Query(postID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var comments []entity.Comment
	for rows.Next() {
		var comment entity.Comment
		err := rows.Scan(
			&comment.ID,
			&comment.Content,
			&comment.PostID,
			&comment.UserID,
			&comment.CreatedAt,
		)
		if err != nil {
			return nil, err
		}
		comments = append(comments, comment)
	}

	return comments, nil
}
