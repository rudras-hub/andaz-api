// Code generated by go-swagger; DO NOT EDIT.

package dialogue

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// GetDialoguesByCharacterReader is a Reader for the GetDialoguesByCharacter structure.
type GetDialoguesByCharacterReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetDialoguesByCharacterReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetDialoguesByCharacterOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetDialoguesByCharacterBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetDialoguesByCharacterNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetDialoguesByCharacterOK creates a GetDialoguesByCharacterOK with default headers values
func NewGetDialoguesByCharacterOK() *GetDialoguesByCharacterOK {
	return &GetDialoguesByCharacterOK{}
}

/*GetDialoguesByCharacterOK handles this case with default header values.

Successful operation.
*/
type GetDialoguesByCharacterOK struct {
	Payload []string
}

func (o *GetDialoguesByCharacterOK) Error() string {
	return fmt.Sprintf("[GET /dialogue/character/{characterName}][%d] getDialoguesByCharacterOK  %+v", 200, o.Payload)
}

func (o *GetDialoguesByCharacterOK) GetPayload() []string {
	return o.Payload
}

func (o *GetDialoguesByCharacterOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetDialoguesByCharacterBadRequest creates a GetDialoguesByCharacterBadRequest with default headers values
func NewGetDialoguesByCharacterBadRequest() *GetDialoguesByCharacterBadRequest {
	return &GetDialoguesByCharacterBadRequest{}
}

/*GetDialoguesByCharacterBadRequest handles this case with default header values.

Invalid name supplied
*/
type GetDialoguesByCharacterBadRequest struct {
}

func (o *GetDialoguesByCharacterBadRequest) Error() string {
	return fmt.Sprintf("[GET /dialogue/character/{characterName}][%d] getDialoguesByCharacterBadRequest ", 400)
}

func (o *GetDialoguesByCharacterBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetDialoguesByCharacterNotFound creates a GetDialoguesByCharacterNotFound with default headers values
func NewGetDialoguesByCharacterNotFound() *GetDialoguesByCharacterNotFound {
	return &GetDialoguesByCharacterNotFound{}
}

/*GetDialoguesByCharacterNotFound handles this case with default header values.

Character name not found
*/
type GetDialoguesByCharacterNotFound struct {
}

func (o *GetDialoguesByCharacterNotFound) Error() string {
	return fmt.Sprintf("[GET /dialogue/character/{characterName}][%d] getDialoguesByCharacterNotFound ", 404)
}

func (o *GetDialoguesByCharacterNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}