// Code generated by go-swagger; DO NOT EDIT.

package client_registration

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/timdrysdale/gojans/pkg/jansAuth/models"
)

// DeleteRegisterReader is a Reader for the DeleteRegister structure.
type DeleteRegisterReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteRegisterReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDeleteRegisterNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDeleteRegisterBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewDeleteRegisterUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewDeleteRegisterInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDeleteRegisterNoContent creates a DeleteRegisterNoContent with default headers values
func NewDeleteRegisterNoContent() *DeleteRegisterNoContent {
	return &DeleteRegisterNoContent{}
}

/*DeleteRegisterNoContent handles this case with default header values.

OK
*/
type DeleteRegisterNoContent struct {
}

func (o *DeleteRegisterNoContent) Error() string {
	return fmt.Sprintf("[DELETE /register][%d] deleteRegisterNoContent ", 204)
}

func (o *DeleteRegisterNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteRegisterBadRequest creates a DeleteRegisterBadRequest with default headers values
func NewDeleteRegisterBadRequest() *DeleteRegisterBadRequest {
	return &DeleteRegisterBadRequest{}
}

/*DeleteRegisterBadRequest handles this case with default header values.

Invalid parameters provided to endpoint.
*/
type DeleteRegisterBadRequest struct {
	Payload *models.ErrorResponse
}

func (o *DeleteRegisterBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /register][%d] deleteRegisterBadRequest  %+v", 400, o.Payload)
}

func (o *DeleteRegisterBadRequest) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *DeleteRegisterBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteRegisterUnauthorized creates a DeleteRegisterUnauthorized with default headers values
func NewDeleteRegisterUnauthorized() *DeleteRegisterUnauthorized {
	return &DeleteRegisterUnauthorized{}
}

/*DeleteRegisterUnauthorized handles this case with default header values.

Invalid parameters provided to endpoint.
*/
type DeleteRegisterUnauthorized struct {
	Payload *models.ErrorResponse
}

func (o *DeleteRegisterUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /register][%d] deleteRegisterUnauthorized  %+v", 401, o.Payload)
}

func (o *DeleteRegisterUnauthorized) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *DeleteRegisterUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteRegisterInternalServerError creates a DeleteRegisterInternalServerError with default headers values
func NewDeleteRegisterInternalServerError() *DeleteRegisterInternalServerError {
	return &DeleteRegisterInternalServerError{}
}

/*DeleteRegisterInternalServerError handles this case with default header values.

Internal error occured. Please check log file for details.
*/
type DeleteRegisterInternalServerError struct {
	Payload *models.ErrorResponse
}

func (o *DeleteRegisterInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /register][%d] deleteRegisterInternalServerError  %+v", 500, o.Payload)
}

func (o *DeleteRegisterInternalServerError) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *DeleteRegisterInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
