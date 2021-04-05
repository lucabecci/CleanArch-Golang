package repository

import "github.com/lucabecci/CleanArch-Golang/src/domain"

type DBHandler interface {
	FindAllAnimes() ([]domain.Anime, error)
	SaveAnime(anime domain.Anime) error
	SaveAuthor(author domain.Author) error
}
