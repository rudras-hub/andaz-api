// Code generated by go-swagger; DO NOT EDIT.

package characters

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"andaz-api/pkg/models"
)

// GetCharactersReader is a Reader for the GetCharacters structure.
type GetCharactersReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetCharactersReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetCharactersOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetCharactersOK creates a GetCharactersOK with default headers values
func NewGetCharactersOK() *GetCharactersOK {
	return &GetCharactersOK{}
}

/*GetCharactersOK handles this case with default header values.

Successful operation.
*/
type GetCharactersOK struct {
	Payload []*models.Character
}

func (o *GetCharactersOK) Error() string {
	return fmt.Sprintf("[GET /characters][%d] getCharactersOK  %+v", 200, o.Payload)
}

func (o *GetCharactersOK) GetPayload() []*models.Character {
	return o.Payload
}

func (o *GetCharactersOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
