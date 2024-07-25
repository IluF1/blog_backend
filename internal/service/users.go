package service

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"guthub.com/server/pkg/database/postgresql"
	"guthub.com/server/pkg/logger"
)


func GetUser(w http.ResponseWriter, r *http.Request) {
	db, err := postgresql.New("postgres://postgres:12345@localhost:5432/blog?sslmode=disable")
	if err != nil {
		logger.Logger.Fatal(err.Error())
	}
	vars := mux.Vars(r)
	email := vars["email"]

	defer db.Close()

	user, err := db.GetUserByEmail(email)
	if err != nil {
		logger.Logger.Fatal(err.Error())
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
	convert, err := json.Marshal(user)
	if err != nil {
		logger.Logger.Fatal(err.Error())
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(convert)
}