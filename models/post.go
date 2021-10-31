package models

import (
	"time"

	"github.com/daonham/go-fiber/database"
	"github.com/daonham/go-fiber/forms"
)

type Post struct {
	Id        int       `json:"id"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	Published bool      `json:"published"`
	UserId    int       `json:"userId"`
}

func CreatePost(create *forms.PostForm) (id *int64, message string, err error) {
	db, err := database.ConnectDB()

	if err != nil {
		return nil, "Cannot connect to database", err
	}

	insert, err := db.Prepare("INSERT INTO Post(title, content, published, userId) VALUES(?,?,?,?)")

	if err != nil {
		return nil, "Cannot insert post", err
	}

	defer db.Close()

	res, err := insert.Exec(create.Title, create.Content, create.Published, create.UserId)

	if err != nil {
		return nil, "Cannot insert post", err
	}

	lastId, err := res.LastInsertId()

	if err != nil {
		return nil, "Cannot get lastId", err
	}

	return &lastId, "Create post successfully", nil
}
