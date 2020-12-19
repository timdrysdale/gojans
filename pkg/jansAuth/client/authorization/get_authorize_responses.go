// Code generated by go-swagger; DO NOT EDIT.

package authorization

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/timdrysdale/gojans/pkg/jansAuth/models"
)

// GetAuthorizeReader is a Reader for the GetAuthorize structure.
type GetAuthorizeReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAuthorizeReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetAuthorizeOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetAuthorizeBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetAuthorizeUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetAuthorizeInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetAuthorizeOK creates a GetAuthorizeOK with default headers values
func NewGetAuthorizeOK() *GetAuthorizeOK {
	return &GetAuthorizeOK{}
}

/*GetAuthorizeOK handles this case with default header values.

OK
*/
type GetAuthorizeOK struct {
}

func (o *GetAuthorizeOK) Error() string {
	return fmt.Sprintf("[GET /authorize][%d] getAuthorizeOK ", 200)
}

func (o *GetAuthorizeOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetAuthorizeBadRequest creates a GetAuthorizeBadRequest with default headers values
func NewGetAuthorizeBadRequest() *GetAuthorizeBadRequest {
	return &GetAuthorizeBadRequest{}
}

/*GetAuthorizeBadRequest handles this case with default header values.

Invalid parameters are provided to endpoint.
*/
type GetAuthorizeBadRequest struct {
	Payload *models.ErrorResponse
}

func (o *GetAuthorizeBadRequest) Error() string {
	return fmt.Sprintf("[GET /authorize][%d] getAuthorizeBadRequest  %+v", 400, o.Payload)
}

func (o *GetAuthorizeBadRequest) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *GetAuthorizeBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAuthorizeUnauthorized creates a GetAuthorizeUnauthorized with default headers values
func NewGetAuthorizeUnauthorized() *GetAuthorizeUnauthorized {
	return &GetAuthorizeUnauthorized{}
}

/*GetAuthorizeUnauthorized handles this case with default header values.

Unauthorized access request.
*/
type GetAuthorizeUnauthorized struct {
	Payload *models.ErrorResponse
}

func (o *GetAuthorizeUnauthorized) Error() string {
	return fmt.Sprintf("[GET /authorize][%d] getAuthorizeUnauthorized  %+v", 401, o.Payload)
}

func (o *GetAuthorizeUnauthorized) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *GetAuthorizeUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAuthorizeInternalServerError creates a GetAuthorizeInternalServerError with default headers values
func NewGetAuthorizeInternalServerError() *GetAuthorizeInternalServerError {
	return &GetAuthorizeInternalServerError{}
}

/*GetAuthorizeInternalServerError handles this case with default header values.

Internal error occured. Please check log file for details.
*/
type GetAuthorizeInternalServerError struct {
	Payload *models.ErrorResponse
}

func (o *GetAuthorizeInternalServerError) Error() string {
	return fmt.Sprintf("[GET /authorize][%d] getAuthorizeInternalServerError  %+v", 500, o.Payload)
}

func (o *GetAuthorizeInternalServerError) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *GetAuthorizeInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
