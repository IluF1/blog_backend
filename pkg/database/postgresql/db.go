package postgresql

import (
	"database/sql"
	"guthub.com/server/internal/models"
	"guthub.com/server/pkg/logger"
)

type storage struct {
	db *sql.DB
}

func New(connStr string) (*storage, error) {
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		logger.Logger.Fatal(err.Error())
	}

	err = db.Ping()
	if err != nil {
		logger.Logger.Fatal(err.Error())
	}

	return &storage{db: db}, nil
}

func (s *storage) Close() {
	s.db.Close()
}

func (s *storage) GetAllPosts() []models.Post {

	rows, err := s.db.Query("SELECT name, id, password, email FROM posts")
	if err != nil {
		logger.Logger.Fatal("Failed query" + err.Error())
	}

	var posts = make([]models.Post, 0)
	for rows.Next() {
		var post models.Post
		err := rows.Scan(&post.Id, &post.Name, &post.Body, &post.UserName)
		if err != nil {
			logger.Logger.Fatal(err.Error())
		}

		posts = append(posts, post)
	}
	return posts
}

func (s *storage) GetUserByEmail(email string) models.User {
	row, err := s.db.Query("SELECT id, name, password FROM users WHERE email = $1", email)
	if err != nil {
		logger.Logger.Fatal(err.Error())
	}
	var user models.User
	for row.Next() {

		err := row.Scan(&user.Id, &user.Name, &user.Password)
		if err != nil {
			logger.Logger.Fatal(err.Error())
		}
	}

	return user
}
