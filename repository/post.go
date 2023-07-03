package repository

import (
	"fmt"
	"forum/entity"
	"time"
)

func (r repo) CreatePost(post *entity.Post) (error, int) {
	stmt, err := r.db.Prepare(`
		INSERT INTO Post (Title, Content, UserID, CreatedAt, UpdateAt)
		VALUES (?, ?, ?, ?, ?)
	`)
	if err != nil {
		return err, 0
	}
	defer stmt.Close()

	stm, err := stmt.Exec(
		post.Title,
		post.Content,
		post.UserID,
		time.Now(),
		post.UpdateAt,
	)
	if err != nil {
		return err, 0
	}

	n, err := stm.LastInsertId()
	post.ID = int(n)
	fmt.Println(post.ID)

	return nil, post.ID
}

func (r repo) UpdatePost(post *entity.Post) error {
	stmt, err := r.db.Prepare(`
		UPDATE Post
		SET Content = ?, UpdateAt = ?
		WHERE ID = ?
	`)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(
		post.Content,
		time.Now(),
		post.ID,
	)
	if err != nil {
		return err
	}

	return nil
}

func (r repo) GetMostLikedTenPost() ([]entity.Post, error) {
	fmt.Println("get repo")
	// Execute the SQLite query
	rows, err := r.db.Query(`
	SELECT Post.ID, Post.Title, Post.Content, Post.UserID, Post.CreatedAt, Post.UpdateAt, COUNT(VotePost.Like) AS LikeCount
	FROM Post
	JOIN VotePost ON Post.ID = VotePost.PostID
	GROUP BY Post.ID
	ORDER BY LikeCount DESC
	LIMIT 10;
	`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// Create a slice to hold the post results
	posts := []entity.Post{}

	// Iterate over the query results
	for rows.Next() {
		var post entity.Post

		// Scan the row values into the Post struct
		err := rows.Scan(&post.ID, &post.Title, &post.Content, &post.UserID, &post.CreatedAt, &post.UpdateAt, &post.CountLike)
		if err != nil {
			return nil, err
		}

		// Append the post to the slice
		posts = append(posts, post)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}
	return posts, nil
}

func (r repo) DeletePost(postID int) error {
	stmt, err := r.db.Prepare(`
		DELETE FROM Post
		WHERE ID = ?
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

func (r *repo) GetPostByID(postID int) (*entity.Post, error) {
	post := &entity.Post{}
	query := `
		SELECT ID, Title, Content, UserID, CreatedAt, UpdateAT
		FROM Post
		WHERE ID = ?
	`
	err := r.db.QueryRow(query, postID).Scan(
		&post.ID,
		&post.Title,
		&post.Content,
		&post.UserID,
		&post.CreatedAt,
		&post.UpdateAt,
	)
	if err != nil {
		return nil, err
	}
	return post, nil
}
