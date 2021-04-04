package usecases

import (
	"log"

	"github.com/lucabecci/CleanArch-Golang/src/domain"
)

type AnimeInteractor struct {
	AnimeRepository domain.AnimeRepository
}

func NewAnimeInteractor(repository domain.AnimeRepository) AnimeInteractor {
	return AnimeInteractor{
		AnimeRepository: repository,
	}
}

func (interactor *AnimeInteractor) CreateAnime(anime domain.Anime) error {
	err := interactor.AnimeRepository.SaveAnime(anime)
	if err != nil {
		log.Println(err.Error())
		return err
	}
	return nil
}

func (interactor *AnimeInteractor) FindAll() ([]*domain.Anime, error) {
	results, err := interactor.AnimeRepository.FindAll()
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}
	return results, nil
}
