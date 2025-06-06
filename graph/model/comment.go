package model

import "time"

type Comment struct {
	ID        string     `json:"id"`
	User      *User      `json:"user"`
	PostID    string     `json:"postId"`
	Body      string     `json:"body"`
	CreatedAt time.Time  `json:"createdAt"`
	UpdatedAt *time.Time `json:"updatedAt,omitempty"`
}
