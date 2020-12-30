package usecase

import (
	"andaz-api/internal/domain/entity"
	"encoding/json"
	"io"
	"net/http"
)

// ErrorType defines custom type of error.
type ErrorType string

// ErrorFormat with a custom error message.
type ErrorFormat struct{
	Type ErrorType `json:"type"`
	Message string `json:"message"`
}

const(
	contentType = "Content-Type"
	applicationJson = "application/json"
	GenericErrorType ErrorType = "Generic Error"
	ValidationErrorType ErrorType = "Validation Error"
)

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

// respondWithInternalServerError logs an error and returns http.StatusInternalServerError to the stream.
func respondWithInternalServerError(log entity.Logger, rw http.ResponseWriter, errorMessage string){
	log.Errorf("Internal Server Error %v", errorMessage)
	rw.WriteHeader(http.StatusInternalServerError)
	MarshallToJson(&ErrorFormat{Type: GenericErrorType, Message: errorMessage}, rw)
}

// respondWithValidationError logs an error and returns http.StatusBadRequest to the stream.
func respondWithValidationError(log entity.Logger, rw http.ResponseWriter, errorMessage string){
	log.Errorf("Validation Error: %v", errorMessage)
	rw.WriteHeader(http.StatusBadRequest)
	MarshallToJson(&ErrorFormat{Type: ValidationErrorType, Message: errorMessage}, rw)
}
