package domain

type Anime struct {
	Title        string `json:"title"`
	Author       string `json:"author"`
	Description  string `json:"description"`
	Category     string `json:"category"`
	Calification int    `json:"calification"`
}

type AnimeRepository interface {
	SaveAnime(anime Anime) error
	FindAll() ([]*Anime, error)
}
