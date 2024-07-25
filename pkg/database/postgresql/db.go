package postgresql

import (
	"database/sql"

	_ "github.com/lib/pq"
	"guthub.com/server/internal/models"
	"guthub.com/server/pkg/logger"
)

type Storage struct {
	db *sql.DB
}

func New(connStr string) (*Storage, error) {
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return &Storage{db: db}, nil
}

func (s *Storage) Close() {
	s.db.Close()
}

func (s *Storage) GetAllPosts() ([]models.Post, error) {
	rows, err := s.db.Query("SELECT id, name, body, username FROM posts")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var posts []models.Post
	for rows.Next() {
		var post models.Post
		err := rows.Scan(&post.Id, &post.Name, &post.Body, &post.UserName)
		if err != nil {
			return nil, err
		}
		posts = append(posts, post)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return posts, nil
}

func (s *Storage) GetUserByEmail(email string) (*models.User, error) {
	row := s.db.QueryRow("SELECT id, name, email FROM users WHERE email = $1", email)

	var user models.User
	err := row.Scan(&user.Id, &user.Name, &user.Email)
	if err != nil {
		if err == sql.ErrNoRows {
			logger.Logger.Panic("Такого пользователя не существует:" + err.Error())
		}
	}

	return &user, nil
}

func (s *Storage) GetPostById(id int) (*models.Post, error) {
	row := s.db.QueryRow("SELECT name, body, username FROM posts WHERE id = $1", id)
	var post models.Post
	err := row.Scan(&post.Name, &post.Body, &post.UserName)

	if err != nil {
		if err == sql.ErrNoRows {
			logger.Logger.Panic("Такой записи не существует:" + err.Error())
		}
	}

	return &post, nil
}