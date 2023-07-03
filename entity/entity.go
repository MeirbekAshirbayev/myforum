package entity

import (
	"errors"
	"time"
)

var (
	ErrNoRecord = errors.New("models: no matching record found")
	SessionMap  = make(map[string]User)
)

const (
	Session           = "Session"
	LoggedIn loggedIn = "LoggedIn"
)

type loggedIn string

type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Post struct {
	ID           int        `json:"id"`
	Title        string     `json:"title"`
	Content      string     `json:"content"`
	UserID       int        `json:"user_id"`
	CreatedAt    time.Time  `json:"created_at"`
	UpdateAt     *time.Time `json:"update_at"`
	CountLike    int        `json:"count_like"`
	CountDislike int        `json:"count_dislike"`
	Comments     []Comment  `json:"comments"`
}

type Comment struct {
	ID        int       `json:"id"`
	Content   string    `json:"content"`
	PostID    int       `json:"post_id"`
	UserID    int       `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
}

type VotePost struct {
	PostID  int  `json:"post_id"`
	UserID  int  `json:"user_id"`
	Like    bool `json:"like"`
	Dislike bool `json:"dislike"`
}

type VoteComment struct {
	CommentID int  `json:"comment_id"`
	UserID    int  `json:"user_id"`
	Like      bool `json:"like"`
	Dislike   bool `json:"dislike"`
}

type Category struct {
	PostID int    `json:"post_id"`
	Name   string `json:"name"`
}

type PostCreate struct {
	ID        int        `json:"id"`
	Title     string     `json:"title"`
	Cat_Name  []string   `json:"cat_name"`
	Content   string     `json:"content"`
	UserID    int        `json:"user_id"`
	CreatedAt time.Time  `json:"created_at"`
	UpdateAt  *time.Time `json:"update_at"`
}
