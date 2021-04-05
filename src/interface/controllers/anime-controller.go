package controllers

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/lucabecci/CleanArch-Golang/src/domain"
	"github.com/lucabecci/CleanArch-Golang/src/usecases"
)

type AnimeController struct {
	AnimeInteractor usecases.AnimeInteractor
}

func NewAnimeController(animeInteractor usecases.AnimeInteractor) *AnimeController {
	return &AnimeController{AnimeInteractor: animeInteractor}
}

func (controller *AnimeController) Add(res http.ResponseWriter, req http.Request) {
	res.Header().Set("Content-type", "application/json")
	var anime domain.Anime

	err := json.NewDecoder(req.Body).Decode(&anime)
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(res).Encode(errors.New("payload invalid"))
		return
	}

	err2 := controller.AnimeInteractor.CreateAnime(anime)
	if err2 != nil {
		res.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(res).Encode(err.Error())
		return
	}

	res.WriteHeader(http.StatusOK)
	json.NewEncoder(res).Encode("Anime created")
}

func (controller *AnimeController) FindAll(res http.ResponseWriter, req http.Request) {
	res.Header().Set("Content-type", "application/json")
	results, err := controller.AnimeInteractor.FindAll()

	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(res).Encode(err.Error())
		return
	}

	res.WriteHeader(http.StatusOK)
	json.NewEncoder(res).Encode(results)
}
