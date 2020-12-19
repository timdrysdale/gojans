// Code generated by go-swagger; DO NOT EDIT.

package fido_u2f

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/timdrysdale/gojans/go-swagger/gluu-4.2.1/oxauth/models"
)

// PostFIDOU2FAuthenticationReader is a Reader for the PostFIDOU2FAuthentication structure.
type PostFIDOU2FAuthenticationReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostFIDOU2FAuthenticationReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostFIDOU2FAuthenticationOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewPostFIDOU2FAuthenticationForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPostFIDOU2FAuthenticationInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPostFIDOU2FAuthenticationOK creates a PostFIDOU2FAuthenticationOK with default headers values
func NewPostFIDOU2FAuthenticationOK() *PostFIDOU2FAuthenticationOK {
	return &PostFIDOU2FAuthenticationOK{}
}

/*PostFIDOU2FAuthenticationOK handles this case with default header values.

OK
*/
type PostFIDOU2FAuthenticationOK struct {
	Payload *models.AuthenticateStatus
}

func (o *PostFIDOU2FAuthenticationOK) Error() string {
	return fmt.Sprintf("[POST /fido/u2f/authentication][%d] postFidoU2fAuthenticationOK  %+v", 200, o.Payload)
}

func (o *PostFIDOU2FAuthenticationOK) GetPayload() *models.AuthenticateStatus {
	return o.Payload
}

func (o *PostFIDOU2FAuthenticationOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.AuthenticateStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostFIDOU2FAuthenticationForbidden creates a PostFIDOU2FAuthenticationForbidden with default headers values
func NewPostFIDOU2FAuthenticationForbidden() *PostFIDOU2FAuthenticationForbidden {
	return &PostFIDOU2FAuthenticationForbidden{}
}

/*PostFIDOU2FAuthenticationForbidden handles this case with default header values.

Invalid parameters provided to endpoint.
*/
type PostFIDOU2FAuthenticationForbidden struct {
	Payload *models.ErrorResponse
}

func (o *PostFIDOU2FAuthenticationForbidden) Error() string {
	return fmt.Sprintf("[POST /fido/u2f/authentication][%d] postFidoU2fAuthenticationForbidden  %+v", 403, o.Payload)
}

func (o *PostFIDOU2FAuthenticationForbidden) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *PostFIDOU2FAuthenticationForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostFIDOU2FAuthenticationInternalServerError creates a PostFIDOU2FAuthenticationInternalServerError with default headers values
func NewPostFIDOU2FAuthenticationInternalServerError() *PostFIDOU2FAuthenticationInternalServerError {
	return &PostFIDOU2FAuthenticationInternalServerError{}
}

/*PostFIDOU2FAuthenticationInternalServerError handles this case with default header values.

Invalid parameters provided to endpoint.
*/
type PostFIDOU2FAuthenticationInternalServerError struct {
	Payload *models.ErrorResponse
}

func (o *PostFIDOU2FAuthenticationInternalServerError) Error() string {
	return fmt.Sprintf("[POST /fido/u2f/authentication][%d] postFidoU2fAuthenticationInternalServerError  %+v", 500, o.Payload)
}

func (o *PostFIDOU2FAuthenticationInternalServerError) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *PostFIDOU2FAuthenticationInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}