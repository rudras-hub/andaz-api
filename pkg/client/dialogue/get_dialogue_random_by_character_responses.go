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

// GetDialogueRandomByCharacterReader is a Reader for the GetDialogueRandomByCharacter structure.
type GetDialogueRandomByCharacterReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetDialogueRandomByCharacterReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetDialogueRandomByCharacterOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetDialogueRandomByCharacterBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetDialogueRandomByCharacterNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetDialogueRandomByCharacterOK creates a GetDialogueRandomByCharacterOK with default headers values
func NewGetDialogueRandomByCharacterOK() *GetDialogueRandomByCharacterOK {
	return &GetDialogueRandomByCharacterOK{}
}

/*GetDialogueRandomByCharacterOK handles this case with default header values.

Successful operation.
*/
type GetDialogueRandomByCharacterOK struct {
	Payload string
}

func (o *GetDialogueRandomByCharacterOK) Error() string {
	return fmt.Sprintf("[GET /dialogue/random/byCharacter][%d] getDialogueRandomByCharacterOK  %+v", 200, o.Payload)
}

func (o *GetDialogueRandomByCharacterOK) GetPayload() string {
	return o.Payload
}

func (o *GetDialogueRandomByCharacterOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetDialogueRandomByCharacterBadRequest creates a GetDialogueRandomByCharacterBadRequest with default headers values
func NewGetDialogueRandomByCharacterBadRequest() *GetDialogueRandomByCharacterBadRequest {
	return &GetDialogueRandomByCharacterBadRequest{}
}

/*GetDialogueRandomByCharacterBadRequest handles this case with default header values.

Invalid name supplied.
*/
type GetDialogueRandomByCharacterBadRequest struct {
}

func (o *GetDialogueRandomByCharacterBadRequest) Error() string {
	return fmt.Sprintf("[GET /dialogue/random/byCharacter][%d] getDialogueRandomByCharacterBadRequest ", 400)
}

func (o *GetDialogueRandomByCharacterBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetDialogueRandomByCharacterNotFound creates a GetDialogueRandomByCharacterNotFound with default headers values
func NewGetDialogueRandomByCharacterNotFound() *GetDialogueRandomByCharacterNotFound {
	return &GetDialogueRandomByCharacterNotFound{}
}

/*GetDialogueRandomByCharacterNotFound handles this case with default header values.

Character name not found.
*/
type GetDialogueRandomByCharacterNotFound struct {
}

func (o *GetDialogueRandomByCharacterNotFound) Error() string {
	return fmt.Sprintf("[GET /dialogue/random/byCharacter][%d] getDialogueRandomByCharacterNotFound ", 404)
}

func (o *GetDialogueRandomByCharacterNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
