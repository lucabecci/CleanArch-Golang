package repository

import "github.com/lucabecci/CleanArch-Golang/src/domain"

type AuthorRepo struct {
	handler DBHandler
}

func NewAuthorRepo(handler DBHandler) AuthorRepo {
	return AuthorRepo{
		handler: handler,
	}
}

func (repo AuthorRepo) SaveAuthor(author domain.Author) error {
	err := repo.handler.SaveAuthor(author)
	if err != nil {
		return err
	}
	return nil
}
