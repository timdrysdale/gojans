// Code generated by go-swagger; DO NOT EDIT.

package token_introspection

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/timdrysdale/gojans/pkg/jansAuth/models"
)

// PostRptStatusReader is a Reader for the PostRptStatus structure.
type PostRptStatusReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostRptStatusReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostRptStatusOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 405:
		result := NewPostRptStatusMethodNotAllowed()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPostRptStatusInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPostRptStatusOK creates a PostRptStatusOK with default headers values
func NewPostRptStatusOK() *PostRptStatusOK {
	return &PostRptStatusOK{}
}

/*PostRptStatusOK handles this case with default header values.

OK
*/
type PostRptStatusOK struct {
	Payload *models.RptIntrospectionResponse1
}

func (o *PostRptStatusOK) Error() string {
	return fmt.Sprintf("[POST /rpt/status][%d] postRptStatusOK  %+v", 200, o.Payload)
}

func (o *PostRptStatusOK) GetPayload() *models.RptIntrospectionResponse1 {
	return o.Payload
}

func (o *PostRptStatusOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RptIntrospectionResponse1)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostRptStatusMethodNotAllowed creates a PostRptStatusMethodNotAllowed with default headers values
func NewPostRptStatusMethodNotAllowed() *PostRptStatusMethodNotAllowed {
	return &PostRptStatusMethodNotAllowed{}
}

/*PostRptStatusMethodNotAllowed handles this case with default header values.

Introspection of RPT is not allowed.
*/
type PostRptStatusMethodNotAllowed struct {
	Payload *models.ErrorResponse
}

func (o *PostRptStatusMethodNotAllowed) Error() string {
	return fmt.Sprintf("[POST /rpt/status][%d] postRptStatusMethodNotAllowed  %+v", 405, o.Payload)
}

func (o *PostRptStatusMethodNotAllowed) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *PostRptStatusMethodNotAllowed) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostRptStatusInternalServerError creates a PostRptStatusInternalServerError with default headers values
func NewPostRptStatusInternalServerError() *PostRptStatusInternalServerError {
	return &PostRptStatusInternalServerError{}
}

/*PostRptStatusInternalServerError handles this case with default header values.

Invalid parameters provided to endpoint.
*/
type PostRptStatusInternalServerError struct {
	Payload *models.ErrorResponse
}

func (o *PostRptStatusInternalServerError) Error() string {
	return fmt.Sprintf("[POST /rpt/status][%d] postRptStatusInternalServerError  %+v", 500, o.Payload)
}

func (o *PostRptStatusInternalServerError) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *PostRptStatusInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
