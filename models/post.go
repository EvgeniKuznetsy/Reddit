package models

import (
	"time"
)

type Post struct {
	Id         string    `json:"id" db:"id"`
	Author     string    `json:"author" db:"author"`
	Caption    string    `json:"caption" db:"Caption"`
	Body       string    `json:"body" db:"Body"`
	CreateDate time.Time `json:"create_date" db:"create_date"`
	Delated    bool      `json:"" db:"delated"`
}
type OutputPostList struct {
	Post       []Post
	TotalCount int
}

type InputPost struct {
	Author  string
	Caption string
	Body    string
}
type OutPost struct {
	Id         string
	CreateDate time.Time
}
type InputUpdatesPost struct {
	Caption string
	Body    string
}
