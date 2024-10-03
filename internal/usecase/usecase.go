package usecase

import (
	"github.com/aszanky/newords-go-be/internal/models"
	"github.com/aszanky/newords-go-be/internal/repository"
)

type Usecase interface {
	AddNewWords(inp models.Word) (err error)
}

type usecase struct {
	repository repository.Repository
}

func NewUsecase(
	rep repository.Repository,
) Usecase {
	return &usecase{
		repository: rep,
	}
}
