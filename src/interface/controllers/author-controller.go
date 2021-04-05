package controllers

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/lucabecci/CleanArch-Golang/src/domain"
	"github.com/lucabecci/CleanArch-Golang/src/usecases"
)

type AuthorController struct {
	authorInteractor usecases.AuthorInteractor
}

func NewAuthorController(interactor usecases.AuthorInteractor) *AuthorController {
	return &AuthorController{authorInteractor: interactor}
}

func (controller *AuthorController) Add(res http.ResponseWriter, req http.Request) {
	res.Header().Set("Content-type", "application/json")
	var author domain.Author

	err := json.NewDecoder(req.Body).Decode(&author)
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(res).Encode(errors.New("invalid payload"))
		return
	}

	err2 := controller.authorInteractor.CreateAuthor(author)
	if err2 != nil {
		res.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(res).Encode(err2.Error())
		return
	}

	res.WriteHeader(http.StatusOK)
	json.NewEncoder(res).Encode("Author saved")
}
