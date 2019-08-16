// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/sapcc/kubernikus/pkg/api/models"
)

// GetBoostrapConfigReader is a Reader for the GetBoostrapConfig structure.
type GetBoostrapConfigReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetBoostrapConfigReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetBoostrapConfigOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewGetBoostrapConfigDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetBoostrapConfigOK creates a GetBoostrapConfigOK with default headers values
func NewGetBoostrapConfigOK() *GetBoostrapConfigOK {
	return &GetBoostrapConfigOK{}
}

/*GetBoostrapConfigOK handles this case with default header values.

OK
*/
type GetBoostrapConfigOK struct {
	Payload *models.Credentials
}

func (o *GetBoostrapConfigOK) Error() string {
	return fmt.Sprintf("[GET /api/v1/clusters/{name}/bootstrap][%d] getBoostrapConfigOK  %+v", 200, o.Payload)
}

func (o *GetBoostrapConfigOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Credentials)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetBoostrapConfigDefault creates a GetBoostrapConfigDefault with default headers values
func NewGetBoostrapConfigDefault(code int) *GetBoostrapConfigDefault {
	return &GetBoostrapConfigDefault{
		_statusCode: code,
	}
}

/*GetBoostrapConfigDefault handles this case with default header values.

Error
*/
type GetBoostrapConfigDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the get boostrap config default response
func (o *GetBoostrapConfigDefault) Code() int {
	return o._statusCode
}

func (o *GetBoostrapConfigDefault) Error() string {
	return fmt.Sprintf("[GET /api/v1/clusters/{name}/bootstrap][%d] GetBoostrapConfig default  %+v", o._statusCode, o.Payload)
}

func (o *GetBoostrapConfigDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
