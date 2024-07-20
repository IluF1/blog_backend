package models

import "time"

type Post struct {
	Id       int           `json:"id"`
	Name     string        `json:"name"`
	Body     string        `json:"body"`
	Date     time.Duration `json:"date"`
	UserName string        `json:"user_name"`
}
