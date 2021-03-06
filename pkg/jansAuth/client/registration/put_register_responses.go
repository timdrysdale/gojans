// Code generated by go-swagger; DO NOT EDIT.

package registration

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/timdrysdale/gojans/pkg/jansAuth/models"
)

// PutRegisterReader is a Reader for the PutRegister structure.
type PutRegisterReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PutRegisterReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPutRegisterOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPutRegisterBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPutRegisterInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPutRegisterOK creates a PutRegisterOK with default headers values
func NewPutRegisterOK() *PutRegisterOK {
	return &PutRegisterOK{}
}

/*PutRegisterOK handles this case with default header values.

OK
*/
type PutRegisterOK struct {
	Payload *models.RegisterResponseParam
}

func (o *PutRegisterOK) Error() string {
	return fmt.Sprintf("[PUT /register][%d] putRegisterOK  %+v", 200, o.Payload)
}

func (o *PutRegisterOK) GetPayload() *models.RegisterResponseParam {
	return o.Payload
}

func (o *PutRegisterOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RegisterResponseParam)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutRegisterBadRequest creates a PutRegisterBadRequest with default headers values
func NewPutRegisterBadRequest() *PutRegisterBadRequest {
	return &PutRegisterBadRequest{}
}

/*PutRegisterBadRequest handles this case with default header values.

Invalid parameters provided to endpoint.
*/
type PutRegisterBadRequest struct {
	Payload *models.ErrorResponse
}

func (o *PutRegisterBadRequest) Error() string {
	return fmt.Sprintf("[PUT /register][%d] putRegisterBadRequest  %+v", 400, o.Payload)
}

func (o *PutRegisterBadRequest) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *PutRegisterBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutRegisterInternalServerError creates a PutRegisterInternalServerError with default headers values
func NewPutRegisterInternalServerError() *PutRegisterInternalServerError {
	return &PutRegisterInternalServerError{}
}

/*PutRegisterInternalServerError handles this case with default header values.

Internal error occured. Please check log file for details.
*/
type PutRegisterInternalServerError struct {
	Payload *models.ErrorResponse
}

func (o *PutRegisterInternalServerError) Error() string {
	return fmt.Sprintf("[PUT /register][%d] putRegisterInternalServerError  %+v", 500, o.Payload)
}

func (o *PutRegisterInternalServerError) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *PutRegisterInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
