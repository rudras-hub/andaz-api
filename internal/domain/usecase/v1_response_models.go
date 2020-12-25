package usecase

import (
	"andaz-api/internal/domain/entity"
)

// CharacterResponseModel struct representing movie character response model.
type CharacterResponseModel struct {
	Name string `json:"name"`
	Alias string `json:"alias"`
}

// DialogueResponseModel struct representing movie dialogue response model.
type DialogueResponseModel struct {
	CharacterName string `json:"characterName"`
	Quote string `json:"quote"`
}

// CharacterEntityToResponseModel converts entity.Character type to CharacterResponseModel.
func CharacterEntityToResponseModel(characterEntityModel *entity.Character)*CharacterResponseModel{
	return &CharacterResponseModel{
		Name: characterEntityModel.Name,
		Alias: characterEntityModel.Alias,
	}
}

// DialogueEntityToResponseModel converts entity.Dialogue to DialogueResponseModel
func DialogueEntityToResponseModel(dialogueEntityModel *entity.Dialogue) *DialogueResponseModel{
	return &DialogueResponseModel{
		CharacterName: dialogueEntityModel.Character.Name,
		Quote: dialogueEntityModel.Quote,
	}
}