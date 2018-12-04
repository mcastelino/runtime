// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/kata-containers/runtime/virtcontainers/pkg/fireclient/client/models"
)

// PutGuestVsockByIDReader is a Reader for the PutGuestVsockByID structure.
type PutGuestVsockByIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PutGuestVsockByIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 201:
		result := NewPutGuestVsockByIDCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 204:
		result := NewPutGuestVsockByIDNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewPutGuestVsockByIDBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewPutGuestVsockByIDDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPutGuestVsockByIDCreated creates a PutGuestVsockByIDCreated with default headers values
func NewPutGuestVsockByIDCreated() *PutGuestVsockByIDCreated {
	return &PutGuestVsockByIDCreated{}
}

/*PutGuestVsockByIDCreated handles this case with default header values.

Vsock created
*/
type PutGuestVsockByIDCreated struct {
}

func (o *PutGuestVsockByIDCreated) Error() string {
	return fmt.Sprintf("[PUT /vsocks/{id}][%d] putGuestVsockByIdCreated ", 201)
}

func (o *PutGuestVsockByIDCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPutGuestVsockByIDNoContent creates a PutGuestVsockByIDNoContent with default headers values
func NewPutGuestVsockByIDNoContent() *PutGuestVsockByIDNoContent {
	return &PutGuestVsockByIDNoContent{}
}

/*PutGuestVsockByIDNoContent handles this case with default header values.

Vsock updated
*/
type PutGuestVsockByIDNoContent struct {
}

func (o *PutGuestVsockByIDNoContent) Error() string {
	return fmt.Sprintf("[PUT /vsocks/{id}][%d] putGuestVsockByIdNoContent ", 204)
}

func (o *PutGuestVsockByIDNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPutGuestVsockByIDBadRequest creates a PutGuestVsockByIDBadRequest with default headers values
func NewPutGuestVsockByIDBadRequest() *PutGuestVsockByIDBadRequest {
	return &PutGuestVsockByIDBadRequest{}
}

/*PutGuestVsockByIDBadRequest handles this case with default header values.

Vsock cannot be created due to bad input
*/
type PutGuestVsockByIDBadRequest struct {
	Payload *models.Error
}

func (o *PutGuestVsockByIDBadRequest) Error() string {
	return fmt.Sprintf("[PUT /vsocks/{id}][%d] putGuestVsockByIdBadRequest  %+v", 400, o.Payload)
}

func (o *PutGuestVsockByIDBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutGuestVsockByIDDefault creates a PutGuestVsockByIDDefault with default headers values
func NewPutGuestVsockByIDDefault(code int) *PutGuestVsockByIDDefault {
	return &PutGuestVsockByIDDefault{
		_statusCode: code,
	}
}

/*PutGuestVsockByIDDefault handles this case with default header values.

Internal server error
*/
type PutGuestVsockByIDDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the put guest vsock by ID default response
func (o *PutGuestVsockByIDDefault) Code() int {
	return o._statusCode
}

func (o *PutGuestVsockByIDDefault) Error() string {
	return fmt.Sprintf("[PUT /vsocks/{id}][%d] putGuestVsockByID default  %+v", o._statusCode, o.Payload)
}

func (o *PutGuestVsockByIDDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
