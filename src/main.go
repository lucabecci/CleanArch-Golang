package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/lucabecci/CleanArch-Golang/src/infraestructure/db"
	"github.com/lucabecci/CleanArch-Golang/src/infraestructure/router"
	"github.com/lucabecci/CleanArch-Golang/src/interface/controllers"
	"github.com/lucabecci/CleanArch-Golang/src/interface/repository"
	"github.com/lucabecci/CleanArch-Golang/src/usecases"
)

var (
	httpRouter = router.NewMuxRouter()
	dbHandler  db.DBHandler
)

func getAnimeController() controllers.AnimeController {
	animeRepo := repository.NewAnimeRepo(dbHandler)
	animeInteractor := usecases.NewAnimeInteractor(animeRepo)
	animeController := controllers.NewAnimeController(animeInteractor)
	return *animeController
}

func getAuthorController() controllers.AuthorController {
	authorRepo := repository.NewAuthorRepo(dbHandler)
	authorInteractor := usecases.NewAuthorInteractor(authorRepo)
	authorController := controllers.NewAuthorController(authorInteractor)
	return *authorController
}

func main() {
	httpRouter.GET("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "App is running")
	})

	var err error
	dbHandler, err = db.NewDBHandler("mongodb://localhost:27017", "animestore")
	if err != nil {
		log.Println("Unable to connect the DB")
		return
	}
	animeController := getAnimeController()
	authorController := getAuthorController()

	httpRouter.POST("/anime", animeController.Add)
	httpRouter.GET("/anime", animeController.FindAll)
	httpRouter.POST("/author", authorController.Add)

	httpRouter.SERVE(":8080")
}
