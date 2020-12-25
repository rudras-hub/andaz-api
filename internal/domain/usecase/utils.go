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

// GenericError with a custom error message.
type GenericError struct{
	Message string `json:"message"`
}

// MarshallToJson marshalls given type to json and adds to the given io.Writer stream.
func MarshallToJson(v interface{}, w io.Writer) error{
	jEncoder := json.NewEncoder(w)
	return jEncoder.Encode(v)
}

// UnmarshallFromJson unmarshalls json string from input io.Reader stream to the given type.
func UnmarshallFromJson(v interface{}, r io.Reader) error{
	jDecoder := json.NewDecoder(r)
	return jDecoder.Decode(v)
}

// respondWithInternalServerError logs an error and returns to  http.StatusInternalServerError to the stream.
func respondWithInternalServerError(log entity.Logger, rw http.ResponseWriter, errorMessage string){
	log.Errorf(errorMessage)
	rw.WriteHeader(http.StatusInternalServerError)
	MarshallToJson(&GenericError{Message: errorMessage}, rw)
}
