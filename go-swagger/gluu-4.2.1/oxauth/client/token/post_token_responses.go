// Code generated by go-swagger; DO NOT EDIT.

package token

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/timdrysdale/gojans/go-swagger/gluu-4.2.1/oxauth/models"
)

// PostTokenReader is a Reader for the PostToken structure.
type PostTokenReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostTokenReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostTokenOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPostTokenBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPostTokenUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPostTokenForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPostTokenInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPostTokenOK creates a PostTokenOK with default headers values
func NewPostTokenOK() *PostTokenOK {
	return &PostTokenOK{}
}

/*PostTokenOK handles this case with default header values.

OK
*/
type PostTokenOK struct {
	Payload *models.TokenResponse
}

func (o *PostTokenOK) Error() string {
	return fmt.Sprintf("[POST /token][%d] postTokenOK  %+v", 200, o.Payload)
}

func (o *PostTokenOK) GetPayload() *models.TokenResponse {
	return o.Payload
}

func (o *PostTokenOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.TokenResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostTokenBadRequest creates a PostTokenBadRequest with default headers values
func NewPostTokenBadRequest() *PostTokenBadRequest {
	return &PostTokenBadRequest{}
}

/*PostTokenBadRequest handles this case with default header values.

Invalid parameters provided to endpoint.
*/
type PostTokenBadRequest struct {
	Payload *models.ErrorResponse
}

func (o *PostTokenBadRequest) Error() string {
	return fmt.Sprintf("[POST /token][%d] postTokenBadRequest  %+v", 400, o.Payload)
}

func (o *PostTokenBadRequest) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *PostTokenBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostTokenUnauthorized creates a PostTokenUnauthorized with default headers values
func NewPostTokenUnauthorized() *PostTokenUnauthorized {
	return &PostTokenUnauthorized{}
}

/*PostTokenUnauthorized handles this case with default header values.

Unauthorized access request.
*/
type PostTokenUnauthorized struct {
	Payload *models.ErrorResponse
}

func (o *PostTokenUnauthorized) Error() string {
	return fmt.Sprintf("[POST /token][%d] postTokenUnauthorized  %+v", 401, o.Payload)
}

func (o *PostTokenUnauthorized) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *PostTokenUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostTokenForbidden creates a PostTokenForbidden with default headers values
func NewPostTokenForbidden() *PostTokenForbidden {
	return &PostTokenForbidden{}
}

/*PostTokenForbidden handles this case with default header values.

Invalid details provided hence access denied.
*/
type PostTokenForbidden struct {
	Payload *models.ErrorResponse
}

func (o *PostTokenForbidden) Error() string {
	return fmt.Sprintf("[POST /token][%d] postTokenForbidden  %+v", 403, o.Payload)
}

func (o *PostTokenForbidden) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *PostTokenForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostTokenInternalServerError creates a PostTokenInternalServerError with default headers values
func NewPostTokenInternalServerError() *PostTokenInternalServerError {
	return &PostTokenInternalServerError{}
}

/*PostTokenInternalServerError handles this case with default header values.

Internal error occured. Please check log file for details.
*/
type PostTokenInternalServerError struct {
	Payload *models.ErrorResponse
}

func (o *PostTokenInternalServerError) Error() string {
	return fmt.Sprintf("[POST /token][%d] postTokenInternalServerError  %+v", 500, o.Payload)
}

func (o *PostTokenInternalServerError) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *PostTokenInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
