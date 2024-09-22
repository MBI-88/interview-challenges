package github

import (
	"time"
)


// Issue model
type Issue struct {
	Number    int
	HTMLURL   string `json:"html_url"`
	Title     string
	State     string
	User      *User
	CreatedAt time.Time `json:"created_at"`
	Body      string
}

// User model
type User struct {
	Login   string
	HTMLURL string `json:"html_url"`
}
