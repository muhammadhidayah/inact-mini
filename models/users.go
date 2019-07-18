package models

import "time"

type Users struct {
	ID        int64     `json:"id"`
	Username  string    `json:"username"`
	Password  string    `json:"password"`
	LastLogin time.Time `json:"last_login"`
	IsOnline  int       `json:"is_online"`
}
