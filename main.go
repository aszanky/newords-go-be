package main

import (
	"log"

	"github.com/aszanky/newords-go-be/internal/handler"
	"github.com/aszanky/newords-go-be/internal/repository"
	"github.com/aszanky/newords-go-be/internal/usecase"
	"github.com/aszanky/newords-go-be/pkg/database"
)

func main() {
	database, err := database.NewDatabase()
	if err != nil {
		log.Fatalf("Postgresql init: %s", err)
	} else {
		log.Printf("Postgres connected, Status: %#v", database.Stats())
		log.Println()
	}

	//repository
	repository := repository.NewRepository(database)

	//usecase
	usecase := usecase.NewUsecase(repository)

	//handler
	handler := handler.NewHandler(usecase)

	//RUN SERVER
	log.Println("Starting newords services on PORT 8090")
	err = handler.Start(":8090")
	if err != nil {
		log.Fatal(err)
	}
}
