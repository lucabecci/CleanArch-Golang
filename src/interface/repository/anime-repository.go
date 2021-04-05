package repository

import "github.com/lucabecci/CleanArch-Golang/src/domain"

type AnimeRepo struct {
	handler DBHandler
}

func NewAnimeRepo(handler DBHandler) AnimeRepo {
	return AnimeRepo{handler: handler}
}

func (repo AnimeRepo) SaveAnime(anime domain.Anime) error {
	err := repo.handler.SaveAnime(anime)
	if err != nil {
		return err
	}
	return nil
}

func (repo AnimeRepo) FindAll() ([]domain.Anime, error) {
	results, err := repo.handler.FindAllAnimes()
	if err != nil {
		return results, err
	}
	return results, nil
}
