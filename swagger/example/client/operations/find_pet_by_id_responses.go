// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/liuchintao/gostudy/swagger/example/models"
)

// FindPetByIDReader is a Reader for the FindPetByID structure.
type FindPetByIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *FindPetByIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewFindPetByIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewFindPetByIDDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewFindPetByIDOK creates a FindPetByIDOK with default headers values
func NewFindPetByIDOK() *FindPetByIDOK {
	return &FindPetByIDOK{}
}

/*FindPetByIDOK handles this case with default header values.

pet response
*/
type FindPetByIDOK struct {
	Payload *models.Pet
}

func (o *FindPetByIDOK) Error() string {
	return fmt.Sprintf("[GET /pets/{id}][%d] findPetByIdOK  %+v", 200, o.Payload)
}

func (o *FindPetByIDOK) GetPayload() *models.Pet {
	return o.Payload
}

func (o *FindPetByIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Pet)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewFindPetByIDDefault creates a FindPetByIDDefault with default headers values
func NewFindPetByIDDefault(code int) *FindPetByIDDefault {
	return &FindPetByIDDefault{
		_statusCode: code,
	}
}

/*FindPetByIDDefault handles this case with default header values.

unexpected error
*/
type FindPetByIDDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the find pet by id default response
func (o *FindPetByIDDefault) Code() int {
	return o._statusCode
}

func (o *FindPetByIDDefault) Error() string {
	return fmt.Sprintf("[GET /pets/{id}][%d] find pet by id default  %+v", o._statusCode, o.Payload)
}

func (o *FindPetByIDDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *FindPetByIDDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
