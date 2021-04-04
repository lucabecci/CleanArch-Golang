package domain

type Author struct {
	Name    string `json:"name"`
	Age     int    `json:"age"`
	Country string `json:"country"`
}

type AuthorRepository interface {
	SaveAuthor(author Author) error
}
