package usecase

import (
	"andaz-api/internal/domain/entity"
	"fmt"
	"net/http"
)

type CharacterHandler struct{
	logger entity.Logger
	repository entity.CharacterRepositoryReader
}

func NewCharacterHandler(log entity.Logger, db entity.CharacterRepositoryReader) *CharacterHandler{
	return &CharacterHandler{logger: log, repository: db}
}

func (c *CharacterHandler) ListAll(rw http.ResponseWriter, r *http.Request){
	c.logger.Debugf("get all characters")
	rw.Header().Add(contentType, applicationJson)

	characterEntities, err := c.repository.RetrieveAll()
	if err != nil {
		errMessage := fmt.Sprintf("error while fetching characters from repository: %v", err)
		respondWithInternalServerError(c.logger, rw, errMessage)
		return
	}

	var charRespModels []*CharacterResponseModel

	for _, chrctr := range characterEntities{
		respModel := CharacterEntityToResponseModel(chrctr)
		charRespModels = append(charRespModels, respModel)
	}

	err = MarshallToJson(charRespModels, rw)
	if err != nil{
		errMessage := fmt.Sprintf("error while converting character entity models to json: %v", err)
		respondWithInternalServerError(c.logger, rw, errMessage)
	}
}