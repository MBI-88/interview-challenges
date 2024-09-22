package server

// GitModel struct. Model for github data
type GitModel struct {
	TotalCount int `json:"total_count"`
	Items      []*Item
}

// Item struct. Data model
type Item struct {
	HTMLURL   string     `json:"html_url"`
	Title     string     `json:"title"`
	User      *User      `json:"user"`
	Milestone *Milestone `json:"milestone"`
}

// User struct. Data of user
type User struct {
	Login   string `json:"login"`
	HTMLURL string `json:"html_url"`
}

// Milestone struct.
type Milestone struct {
	HTMLURL string `json:"html_url"`
	Title   string `json:"title"`
	State   string `json:"state"`
}
