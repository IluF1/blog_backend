package service

import (
	"encoding/json"
	"net/http"

	"guthub.com/server/pkg/database/postgresql"
	"guthub.com/server/pkg/logger"
)

type emailRequest struct {
	Email string `json:"email"`
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	var req *emailRequest
	db, err := postgresql.New()
	if err != nil {
		logger.Logger.Fatal(err.Error())
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		logger.Logger.Error(err.Error())
	}
	defer db.Close()

	user, err := db.GetUserByEmail(req.Email)
	if err != nil {
		logger.Logger.Fatal(err.Error())
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
	convert, err := json.Marshal(user)
	if err != nil {
		logger.Logger.Fatal(err.Error())
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}

	w.Write(convert)
}

