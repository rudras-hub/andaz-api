package usecase

import (
	"andaz-api/internal/entity"
	"fmt"
	"net/http"
)

// CharacterGetHandler handles http GET requests for entity.Character(s).
type CharacterGetHandler struct{
	logger entity.Logger
	repository entity.CharacterRepositoryReader
}

// NewCharacterGetHandler returns a new instance of CharacterGetHandler with the given
// entity.Logger and entity.CharacterRepositoryReader input interfaces.
func NewCharacterGetHandler(log entity.Logger, db entity.CharacterRepositoryReader) *CharacterGetHandler {
	return &CharacterGetHandler{logger: log, repository: db}
}

// ListAll handles the http method to list all available movie characters.
func (c *CharacterGetHandler) ListAll(rw http.ResponseWriter, r *http.Request){
	c.logger.Debug("get all characters")
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