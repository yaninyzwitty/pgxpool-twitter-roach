package model

import "time"

type Post struct {
	ID        string     `json:"id"`
	User      *User      `json:"user"`
	Body      string     `json:"body"`
	CreatedAt time.Time  `json:"createdAt"`
	UpdatedAt *time.Time `json:"updatedAt,omitempty"`
	Comments  []*Comment `json:"comments,omitempty"`
}
