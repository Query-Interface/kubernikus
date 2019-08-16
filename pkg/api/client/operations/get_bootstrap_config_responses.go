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

// GetBootstrapConfigReader is a Reader for the GetBootstrapConfig structure.
type GetBootstrapConfigReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetBootstrapConfigReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetBootstrapConfigOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewGetBootstrapConfigDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetBootstrapConfigOK creates a GetBootstrapConfigOK with default headers values
func NewGetBootstrapConfigOK() *GetBootstrapConfigOK {
	return &GetBootstrapConfigOK{}
}

/*GetBootstrapConfigOK handles this case with default header values.

OK
*/
type GetBootstrapConfigOK struct {
	Payload *models.BootstrapConfig
}

func (o *GetBootstrapConfigOK) Error() string {
	return fmt.Sprintf("[GET /api/v1/clusters/{name}/bootstrap][%d] getBootstrapConfigOK  %+v", 200, o.Payload)
}

func (o *GetBootstrapConfigOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.BootstrapConfig)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetBootstrapConfigDefault creates a GetBootstrapConfigDefault with default headers values
func NewGetBootstrapConfigDefault(code int) *GetBootstrapConfigDefault {
	return &GetBootstrapConfigDefault{
		_statusCode: code,
	}
}

/*GetBootstrapConfigDefault handles this case with default header values.

Error
*/
type GetBootstrapConfigDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the get bootstrap config default response
func (o *GetBootstrapConfigDefault) Code() int {
	return o._statusCode
}

func (o *GetBootstrapConfigDefault) Error() string {
	return fmt.Sprintf("[GET /api/v1/clusters/{name}/bootstrap][%d] GetBootstrapConfig default  %+v", o._statusCode, o.Payload)
}

func (o *GetBootstrapConfigDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
