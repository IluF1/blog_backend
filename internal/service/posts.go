package service

import (
	"encoding/json"
	"net/http"

	_ "github.com/lib/pq"
	"go.uber.org/zap"
	"guthub.com/server/pkg/database/postgresql"
	"guthub.com/server/pkg/logger"
)

func GetPosts(w http.ResponseWriter, r *http.Request) {

	db, err := postgresql.New()
	if err != nil {
		logger.Logger.Error("Failed to connect to the database", zap.Error(err))
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)

	}
	defer db.Close()

	posts, err := db.GetAllPosts()
	if err != nil {
		logger.Logger.Error("Failed to get posts", zap.Error(err))
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}

	convert, err := json.Marshal(posts)
	if err != nil {
		logger.Logger.Error("Failed to marshal posts", zap.Error(err))
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)

	}
	
	w.Write(convert)
}

// func GetPostById(w http.ResponseWriter, r *http.Request) {
// 	db, err := postgresql.New("postgres://postgres:12345@localhost:5432/blog?sslmode=disable")
// 	if err != nil {
// 		logger.Logger.Fatal("Failed to connect to the database:" + err.Error())
// 	}
// 	defer db.Close()

// 	id := r.URL.Query().Get("id")
// 	post, err := db.GetPostById()
// 	if err != nil {
// 		logger.Logger.Fatal("Failed to get post:" + err.Error())
// 		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
// 	}

// 	convert, err := json.Marshal(post)
// 	if err != nil {
// 		logger.Logger.Fatal("Failed to convert post to json format:" + err.Error())
// 		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
// 	}
// 	w.Write(convert)

// }
