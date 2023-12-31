package main

import (
	"net/http"

	"github.com/go-playground/validator"
	"github.com/rs/zerolog/log"

	"golang/config"
	"golang/controller"
	"golang/helper"
	"golang/model"
	"golang/repository"
	"golang/router"
	"golang/service"
)

func main() {

	log.Info().Msg("Starting Server")

	//Database
	db := config.DatabaseConnection()
	validate := validator.New()

	db.Table("tags").AutoMigrate(&model.Tags{})

	// Repository
	tagsRepository := repository.NewTagsRepositoryImpl(db)

	// Service
	tagsService := service.NewTagsServiceImpl(tagsRepository, validate)

	// Controller
	TagsController := controller.NewTagsController(tagsService)

	// Router
	routes := router.NewRouter(TagsController)

	server := &http.Server{
		Addr:    ":8888",
		Handler: routes,
	}

	err := server.ListenAndServe()
	helper.ErrorPanic(err)

}
