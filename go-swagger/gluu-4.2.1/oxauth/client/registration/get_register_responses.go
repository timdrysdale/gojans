// Code generated by go-swagger; DO NOT EDIT.

package registration

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/timdrysdale/gojans/go-swagger/gluu-4.2.1/oxauth/models"
)

// GetRegisterReader is a Reader for the GetRegister structure.
type GetRegisterReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetRegisterReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetRegisterOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetRegisterBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetRegisterUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetRegisterInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetRegisterOK creates a GetRegisterOK with default headers values
func NewGetRegisterOK() *GetRegisterOK {
	return &GetRegisterOK{}
}

/*GetRegisterOK handles this case with default header values.

OK
*/
type GetRegisterOK struct {
	Payload *models.ClientResponse
}

func (o *GetRegisterOK) Error() string {
	return fmt.Sprintf("[GET /register][%d] getRegisterOK  %+v", 200, o.Payload)
}

func (o *GetRegisterOK) GetPayload() *models.ClientResponse {
	return o.Payload
}

func (o *GetRegisterOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ClientResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRegisterBadRequest creates a GetRegisterBadRequest with default headers values
func NewGetRegisterBadRequest() *GetRegisterBadRequest {
	return &GetRegisterBadRequest{}
}

/*GetRegisterBadRequest handles this case with default header values.

Invalid parameters provided to endpoint.
*/
type GetRegisterBadRequest struct {
	Payload *models.ErrorResponse
}

func (o *GetRegisterBadRequest) Error() string {
	return fmt.Sprintf("[GET /register][%d] getRegisterBadRequest  %+v", 400, o.Payload)
}

func (o *GetRegisterBadRequest) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *GetRegisterBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRegisterUnauthorized creates a GetRegisterUnauthorized with default headers values
func NewGetRegisterUnauthorized() *GetRegisterUnauthorized {
	return &GetRegisterUnauthorized{}
}

/*GetRegisterUnauthorized handles this case with default header values.

Invalid parameters are provided to endpoint.
*/
type GetRegisterUnauthorized struct {
	Payload *models.ErrorResponse
}

func (o *GetRegisterUnauthorized) Error() string {
	return fmt.Sprintf("[GET /register][%d] getRegisterUnauthorized  %+v", 401, o.Payload)
}

func (o *GetRegisterUnauthorized) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *GetRegisterUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRegisterInternalServerError creates a GetRegisterInternalServerError with default headers values
func NewGetRegisterInternalServerError() *GetRegisterInternalServerError {
	return &GetRegisterInternalServerError{}
}

/*GetRegisterInternalServerError handles this case with default header values.

Internal error occured. Please check log file for details.
*/
type GetRegisterInternalServerError struct {
	Payload *models.ErrorResponse
}

func (o *GetRegisterInternalServerError) Error() string {
	return fmt.Sprintf("[GET /register][%d] getRegisterInternalServerError  %+v", 500, o.Payload)
}

func (o *GetRegisterInternalServerError) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *GetRegisterInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}