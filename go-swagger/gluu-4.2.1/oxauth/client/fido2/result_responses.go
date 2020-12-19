// Code generated by go-swagger; DO NOT EDIT.

package fido2

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/timdrysdale/gojans/go-swagger/gluu-4.2.1/oxauth/models"
)

// ResultReader is a Reader for the Result structure.
type ResultReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ResultReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewResultOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewResultForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewResultInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewResultOK creates a ResultOK with default headers values
func NewResultOK() *ResultOK {
	return &ResultOK{}
}

/*ResultOK handles this case with default header values.

OK
*/
type ResultOK struct {
}

func (o *ResultOK) Error() string {
	return fmt.Sprintf("[POST /fido2/assertion/result][%d] resultOK ", 200)
}

func (o *ResultOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewResultForbidden creates a ResultForbidden with default headers values
func NewResultForbidden() *ResultForbidden {
	return &ResultForbidden{}
}

/*ResultForbidden handles this case with default header values.

Invalid details provided hence access denied.
*/
type ResultForbidden struct {
	Payload *models.ErrorResponse
}

func (o *ResultForbidden) Error() string {
	return fmt.Sprintf("[POST /fido2/assertion/result][%d] resultForbidden  %+v", 403, o.Payload)
}

func (o *ResultForbidden) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *ResultForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewResultInternalServerError creates a ResultInternalServerError with default headers values
func NewResultInternalServerError() *ResultInternalServerError {
	return &ResultInternalServerError{}
}

/*ResultInternalServerError handles this case with default header values.

Internal error occured. Please check log file for details.
*/
type ResultInternalServerError struct {
	Payload *models.ErrorResponse
}

func (o *ResultInternalServerError) Error() string {
	return fmt.Sprintf("[POST /fido2/assertion/result][%d] resultInternalServerError  %+v", 500, o.Payload)
}

func (o *ResultInternalServerError) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *ResultInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
