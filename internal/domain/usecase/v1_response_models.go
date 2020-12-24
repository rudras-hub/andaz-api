package usecase

import (
	"andaz-api/internal/domain/entity"
)

type CharacterResponseModel struct {
	Name string `json:"name"`
	Alias string `json:"alias"`
}

type DialogueResponseModel struct {
	CharacterName string `json:"characterName"`
	Quote string `json:"quote"`
}

func CharacterEntityToResponseModel(characterEntityModel *entity.Character)*CharacterResponseModel{
	return &CharacterResponseModel{
		Name: characterEntityModel.Name,
		Alias: characterEntityModel.Alias,
	}
}

func DialogueEntityToResponseModel(dialogueEntityModel *entity.Dialogue) *DialogueResponseModel{
	return &DialogueResponseModel{
		CharacterName: dialogueEntityModel.Character.Name,
		Quote: dialogueEntityModel.Quote,
	}
}