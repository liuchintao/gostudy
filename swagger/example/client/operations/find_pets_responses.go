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

// FindPetsReader is a Reader for the FindPets structure.
type FindPetsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *FindPetsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewFindPetsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewFindPetsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewFindPetsOK creates a FindPetsOK with default headers values
func NewFindPetsOK() *FindPetsOK {
	return &FindPetsOK{}
}

/*FindPetsOK handles this case with default header values.

pet response
*/
type FindPetsOK struct {
	Payload []*models.Pet
}

func (o *FindPetsOK) Error() string {
	return fmt.Sprintf("[GET /pets][%d] findPetsOK  %+v", 200, o.Payload)
}

func (o *FindPetsOK) GetPayload() []*models.Pet {
	return o.Payload
}

func (o *FindPetsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewFindPetsDefault creates a FindPetsDefault with default headers values
func NewFindPetsDefault(code int) *FindPetsDefault {
	return &FindPetsDefault{
		_statusCode: code,
	}
}

/*FindPetsDefault handles this case with default header values.

unexpected error
*/
type FindPetsDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the find pets default response
func (o *FindPetsDefault) Code() int {
	return o._statusCode
}

func (o *FindPetsDefault) Error() string {
	return fmt.Sprintf("[GET /pets][%d] findPets default  %+v", o._statusCode, o.Payload)
}

func (o *FindPetsDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *FindPetsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
