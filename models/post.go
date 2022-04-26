package models

import (
	"time"
)

type Post struct {
	Id         string    `json:"id" db:"iteam_id"`
	Author     string    `json:"author" db:"author"`
	Caption    string    `json:"caption" db:"Caption"`
	Body       string    `json:"body" db:"Body"`
	CreateDate time.Time `json:"create_date" db:"create_date"`
	Delated    bool      `json:"" db:"deleted"`
}
type OutputPostList struct {
	Post       []Post
	TotalCount int
}

type InputPost struct {
	Author  string `binding:"requiring"`
	Caption string `binding:"requiring"`
	Body    string `binding:"requiring"`
}
type OutPost struct {
	Id         string
	CreateDate time.Time
}
type InputUpdatesPost struct {
	Caption string
	Body    string
	Id      string
}
