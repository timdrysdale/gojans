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

// PostRegisterReader is a Reader for the PostRegister structure.
type PostRegisterReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostRegisterReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostRegisterOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPostRegisterBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPostRegisterInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPostRegisterOK creates a PostRegisterOK with default headers values
func NewPostRegisterOK() *PostRegisterOK {
	return &PostRegisterOK{}
}

/*PostRegisterOK handles this case with default header values.

OK
*/
type PostRegisterOK struct {
	Payload *models.RegisterResponseParam
}

func (o *PostRegisterOK) Error() string {
	return fmt.Sprintf("[POST /register][%d] postRegisterOK  %+v", 200, o.Payload)
}

func (o *PostRegisterOK) GetPayload() *models.RegisterResponseParam {
	return o.Payload
}

func (o *PostRegisterOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RegisterResponseParam)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostRegisterBadRequest creates a PostRegisterBadRequest with default headers values
func NewPostRegisterBadRequest() *PostRegisterBadRequest {
	return &PostRegisterBadRequest{}
}

/*PostRegisterBadRequest handles this case with default header values.

Invalid parameters provided to endpoint.
*/
type PostRegisterBadRequest struct {
	Payload *models.ErrorResponse
}

func (o *PostRegisterBadRequest) Error() string {
	return fmt.Sprintf("[POST /register][%d] postRegisterBadRequest  %+v", 400, o.Payload)
}

func (o *PostRegisterBadRequest) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *PostRegisterBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostRegisterInternalServerError creates a PostRegisterInternalServerError with default headers values
func NewPostRegisterInternalServerError() *PostRegisterInternalServerError {
	return &PostRegisterInternalServerError{}
}

/*PostRegisterInternalServerError handles this case with default header values.

Internal error occured. Please check log file for details.
*/
type PostRegisterInternalServerError struct {
	Payload *models.ErrorResponse
}

func (o *PostRegisterInternalServerError) Error() string {
	return fmt.Sprintf("[POST /register][%d] postRegisterInternalServerError  %+v", 500, o.Payload)
}

func (o *PostRegisterInternalServerError) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *PostRegisterInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
