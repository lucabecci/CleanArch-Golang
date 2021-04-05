package usecases

import (
	"log"

	"github.com/lucabecci/CleanArch-Golang/src/domain"
)

type AuthorInteractor struct {
	AuthorRepository domain.AuthorRepository
}

func NewAuthorInteractor(repository domain.AuthorRepository) AuthorInteractor {
	return AuthorInteractor{
		AuthorRepository: repository,
	}
}

func (interactor *AuthorInteractor) CreateAuthor(author domain.Author) error {
	err := interactor.AuthorRepository.SaveAuthor(author)
	if err != nil {
		log.Println(err.Error())
		return err
	}

	return nil
}
