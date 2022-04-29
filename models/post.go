package models

import (
	"time"
)

type Post struct {
	Id         string    `json:"id" db:"item_id"`
	Author     string    `json:"author" db:"author"`
	Caption    string    `json:"caption" db:"Caption"`
	Body       string    `json:"body" db:"body"`
	CreateDate time.Time `json:"create_date" db:"create_date"`
	Deleted    bool      `json:"-" db:"deleted"`
}
type OutputPostList struct {
	Post       []Post `json:"posts"`
	TotalCount int    `json:"total_count"`
}

type InputPost struct {
	Author  string `binding:"required"`
	Caption string `binding:"required"`
	Body    string `binding:"required"`
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

type UriGetPostList struct {
	Page  string `uri:"page"`
	Limit string `uri:"limit"`
}
