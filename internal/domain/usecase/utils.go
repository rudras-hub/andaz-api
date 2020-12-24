package usecase

import (
	"andaz-api/internal/domain/entity"
	"encoding/json"
	"io"
	"net/http"
)

const(
	contentType = "Content-Type"
	applicationJson = "application/json"
)

type GenericError struct{
	Message string `json:"message"`
}

func MarshallToJson(v interface{}, w io.Writer) error{
	jEncoder := json.NewEncoder(w)
	return jEncoder.Encode(v)
}

func UnmarshallFromJson(v interface{}, r io.Reader) error{
	jDecoder := json.NewDecoder(r)
	return jDecoder.Decode(v)
}

func respondWithInternalServerError(log entity.Logger, rw http.ResponseWriter, errorMessage string){
	log.Errorf(errorMessage)
	rw.WriteHeader(http.StatusInternalServerError)
	MarshallToJson(&GenericError{Message: errorMessage}, rw)
}
