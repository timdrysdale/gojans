// Code generated by go-swagger; DO NOT EDIT.

package user_info

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/timdrysdale/gojans/src/jansAuth/models"
)

// PostUserinfoReader is a Reader for the PostUserinfo structure.
type PostUserinfoReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostUserinfoReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostUserinfoOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPostUserinfoBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPostUserinfoUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPostUserinfoForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPostUserinfoInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPostUserinfoOK creates a PostUserinfoOK with default headers values
func NewPostUserinfoOK() *PostUserinfoOK {
	return &PostUserinfoOK{}
}

/*PostUserinfoOK handles this case with default header values.

OK
*/
type PostUserinfoOK struct {
	Payload interface{}
}

func (o *PostUserinfoOK) Error() string {
	return fmt.Sprintf("[POST /userinfo][%d] postUserinfoOK  %+v", 200, o.Payload)
}

func (o *PostUserinfoOK) GetPayload() interface{} {
	return o.Payload
}

func (o *PostUserinfoOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostUserinfoBadRequest creates a PostUserinfoBadRequest with default headers values
func NewPostUserinfoBadRequest() *PostUserinfoBadRequest {
	return &PostUserinfoBadRequest{}
}

/*PostUserinfoBadRequest handles this case with default header values.

Invalid parameters provided to endpoint.
*/
type PostUserinfoBadRequest struct {
	Payload *models.ErrorResponse
}

func (o *PostUserinfoBadRequest) Error() string {
	return fmt.Sprintf("[POST /userinfo][%d] postUserinfoBadRequest  %+v", 400, o.Payload)
}

func (o *PostUserinfoBadRequest) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *PostUserinfoBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostUserinfoUnauthorized creates a PostUserinfoUnauthorized with default headers values
func NewPostUserinfoUnauthorized() *PostUserinfoUnauthorized {
	return &PostUserinfoUnauthorized{}
}

/*PostUserinfoUnauthorized handles this case with default header values.

Invalid parameters provided to endpoint.
*/
type PostUserinfoUnauthorized struct {
	Payload *models.ErrorResponse
}

func (o *PostUserinfoUnauthorized) Error() string {
	return fmt.Sprintf("[POST /userinfo][%d] postUserinfoUnauthorized  %+v", 401, o.Payload)
}

func (o *PostUserinfoUnauthorized) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *PostUserinfoUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostUserinfoForbidden creates a PostUserinfoForbidden with default headers values
func NewPostUserinfoForbidden() *PostUserinfoForbidden {
	return &PostUserinfoForbidden{}
}

/*PostUserinfoForbidden handles this case with default header values.

Invalid parameters provided to endpoint.
*/
type PostUserinfoForbidden struct {
	Payload *models.ErrorResponse
}

func (o *PostUserinfoForbidden) Error() string {
	return fmt.Sprintf("[POST /userinfo][%d] postUserinfoForbidden  %+v", 403, o.Payload)
}

func (o *PostUserinfoForbidden) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *PostUserinfoForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostUserinfoInternalServerError creates a PostUserinfoInternalServerError with default headers values
func NewPostUserinfoInternalServerError() *PostUserinfoInternalServerError {
	return &PostUserinfoInternalServerError{}
}

/*PostUserinfoInternalServerError handles this case with default header values.

Internal error occured. Please check log file for details.
*/
type PostUserinfoInternalServerError struct {
	Payload *models.ErrorResponse
}

func (o *PostUserinfoInternalServerError) Error() string {
	return fmt.Sprintf("[POST /userinfo][%d] postUserinfoInternalServerError  %+v", 500, o.Payload)
}

func (o *PostUserinfoInternalServerError) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *PostUserinfoInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}