package models

import "time"

type PostItem struct {
	Id       string    `json:"id"`
	Content  string    `json:"content"`
	CreateAt time.Time `json:"created_at"`
}
