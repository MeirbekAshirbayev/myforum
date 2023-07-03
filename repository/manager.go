package repository

import (
	"database/sql"
	"forum/entity"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

type IDB interface {
	CreateUser(user *entity.User) error
	CheckUsername(username string) (bool, error)
	CheckEmail(email string) (bool, error)
	GetUserByEmail(email string) (*entity.User, error)

	CreatePost(post *entity.Post) (error, int)
	UpdatePost(post *entity.Post) error
	GetMostLikedTenPost() ([]entity.Post, error)
	DeletePost(postID int) error
	GetPostByID(postID int) (*entity.Post, error)

	CreateComment(comment *entity.Comment) error
	GetCommentByPostID(postID int) ([]entity.Comment, error)
	DeleteByID(commentID int) error
	DeleteByPostID(postID int) error

	CreateCategory(names []string, postID int) error
	DeleteCategoryByPostID(postID int) error

	CreateLikePost(vote *entity.VotePost, n int) error
	CreateDislikePost(vote *entity.VotePost, n int) error
	CheckLikeByPostIDAndUserID(postID, userID int) (bool, error)
	CheckDislikeByPostIDAndUserID(postID, userID int) (bool, error)
	DeleteByPostIDAndUserID(postID, userID int) error
}

type repo struct {
	log *log.Logger
	db  *sql.DB
}

func New(l *log.Logger) IDB {
	db, err := sql.Open("sqlite3", "database.db")
	if err != nil {
		log.Fatal(err)
	}

	createTable(db)

	return &repo{log: l, db: db}
}

func createTable(db *sql.DB) {
	// Create User table if it doesn't exist
	_, err := db.Exec(`CREATE TABLE IF NOT EXISTS User (
			ID INTEGER PRIMARY KEY AUTOINCREMENT,
			Username TEXT NOT NULL,
			Email TEXT NOT NULL,
			Password TEXT NOT NULL
		);`)
	if err != nil {
		log.Fatal(err)
	}

	// Create Post table if it doesn't exist
	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS Post (
			ID INTEGER PRIMARY KEY AUTOINCREMENT,
			Title TEXT NOT NULL,
			Content TEXT NOT NULL,
			UserID INTEGER NOT NULL,
			CreatedAt DATETIME DEFAULT CURRENT_TIMESTAMP,
			UpdateAt DATETIME,
			FOREIGN KEY (UserID) REFERENCES User (ID)
		);`)
	if err != nil {
		log.Fatal(err)
	}

	// Create Comment table if it doesn't exist
	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS Comment (
			ID INTEGER PRIMARY KEY AUTOINCREMENT,
			Content TEXT NOT NULL,
			PostID INTEGER NOT NULL,
			UserID INTEGER NOT NULL,
			CreatedAt DATETIME DEFAULT CURRENT_TIMESTAMP,
			FOREIGN KEY (PostID) REFERENCES Post (ID),
			FOREIGN KEY (UserID) REFERENCES User (ID)
		);`)
	if err != nil {
		log.Fatal(err)
	}

	// Create VotePost table if it doesn't exist
	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS VotePost (
			PostID INTEGER,
			UserID INTEGER,
			Like INTEGER,
			Dislike INTEGER,
			FOREIGN KEY (PostID) REFERENCES Post (ID),
			FOREIGN KEY (UserID) REFERENCES User (ID)
		);`)
	if err != nil {
		log.Fatal(err)
	}

	// Create VoteComment table if it doesn't exist
	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS VoteComment (
			CommentID INTEGER,
			UserID INTEGER,
			Like INTEGER,
			Dislike INTEGER,
			FOREIGN KEY (CommentID) REFERENCES Comment (ID),
			FOREIGN KEY (UserID) REFERENCES User (ID)
		);`)
	if err != nil {
		log.Fatal(err)
	}

	// Create Category table if it doesn't exist
	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS Category (
			PostID INTEGER,
			Name TEXT NOT NULL,
			FOREIGN KEY (PostID) REFERENCES Post (ID)
		);`)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Tables created successfully!")
}
