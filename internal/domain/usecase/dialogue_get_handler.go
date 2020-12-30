package usecase

import (
	"andaz-api/internal/domain/entity"
	"context"
	"net/http"
)

const(
	// charNameQueryKey defines the key to use to retrieve character name from request query.
	charNameQueryKey = "name"
)

// charContextKey defines the type to use as key to retrieve entity.Character from context.
type charContextKey struct{}

// DialogueGetHandler handles http GET requests for entity.Dialogue(s).
type DialogueGetHandler struct{
	logger             entity.Logger
	charRepository     entity.CharacterRepositoryReader
	dialogueRepository entity.DialogueRepositoryReader
}

// NewDialogueGetHandler returns a new instance of DialogueGetHandler with the given
// entity.Logger, entity.CharacterRepositoryReader and entity.DialogueRepositoryReader input interfaces.
func NewDialogueGetHandler(log entity.Logger, db entity.DialogueRepositoryReader) *DialogueGetHandler{
	return &DialogueGetHandler{logger: log, dialogueRepository: db}
}
// ValidateCharNameMiddleware is the middleware function to verify if the name in request
// query is a valid character of the movie.
func (d *DialogueGetHandler) ValidateCharNameMiddleware(handler http.Handler) http.Handler{
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		d.logger.Debug("validate character name")
		rw.Header().Add(contentType, applicationJson)

		requestedName := r.URL.Query().Get(charNameQueryKey)

		char, err :=  d.charRepository.RetrieveByName(requestedName)
		if err == nil{
			ctx := context.WithValue(r.Context(), charContextKey{}, char)
			r.WithContext(ctx)
			handler.ServeHTTP(rw, r)
			return
		}

		if _, ok := err.(*entity.CharacterWithNameNotFoundError); ok{
			respondWithValidationError(d.logger, rw, err.Error())
			return
		}

		respondWithInternalServerError(d.logger, rw, err.Error())

	})
}

