package usecase

import "github.com/aszanky/newords-go-be/internal/models"

func (u *usecase) AddNewWords(inp models.Word) (err error) {
	err = u.repository.AddNewWords(inp.Word, inp.Indonesian, inp.Notes)
	if err != nil {
		return
	}

	return nil
}
