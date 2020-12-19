// Code generated by go-swagger; DO NOT EDIT.

package fido2

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/timdrysdale/gojans/pkg/jansAuth/models"
)

// GetFIDO2ConfigurationReader is a Reader for the GetFIDO2Configuration structure.
type GetFIDO2ConfigurationReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetFIDO2ConfigurationReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetFIDO2ConfigurationOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewGetFIDO2ConfigurationForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 406:
		result := NewGetFIDO2ConfigurationNotAcceptable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetFIDO2ConfigurationInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetFIDO2ConfigurationOK creates a GetFIDO2ConfigurationOK with default headers values
func NewGetFIDO2ConfigurationOK() *GetFIDO2ConfigurationOK {
	return &GetFIDO2ConfigurationOK{}
}

/*GetFIDO2ConfigurationOK handles this case with default header values.

OK
*/
type GetFIDO2ConfigurationOK struct {
	Payload *models.FIDO2configuration
}

func (o *GetFIDO2ConfigurationOK) Error() string {
	return fmt.Sprintf("[GET /fido2/configuration][%d] getFido2ConfigurationOK  %+v", 200, o.Payload)
}

func (o *GetFIDO2ConfigurationOK) GetPayload() *models.FIDO2configuration {
	return o.Payload
}

func (o *GetFIDO2ConfigurationOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.FIDO2configuration)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetFIDO2ConfigurationForbidden creates a GetFIDO2ConfigurationForbidden with default headers values
func NewGetFIDO2ConfigurationForbidden() *GetFIDO2ConfigurationForbidden {
	return &GetFIDO2ConfigurationForbidden{}
}

/*GetFIDO2ConfigurationForbidden handles this case with default header values.

Invalid details provided hence access denied.
*/
type GetFIDO2ConfigurationForbidden struct {
	Payload *models.ErrorResponse
}

func (o *GetFIDO2ConfigurationForbidden) Error() string {
	return fmt.Sprintf("[GET /fido2/configuration][%d] getFido2ConfigurationForbidden  %+v", 403, o.Payload)
}

func (o *GetFIDO2ConfigurationForbidden) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *GetFIDO2ConfigurationForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetFIDO2ConfigurationNotAcceptable creates a GetFIDO2ConfigurationNotAcceptable with default headers values
func NewGetFIDO2ConfigurationNotAcceptable() *GetFIDO2ConfigurationNotAcceptable {
	return &GetFIDO2ConfigurationNotAcceptable{}
}

/*GetFIDO2ConfigurationNotAcceptable handles this case with default header values.

Request Not Acceptable.
*/
type GetFIDO2ConfigurationNotAcceptable struct {
	Payload *models.ErrorResponse
}

func (o *GetFIDO2ConfigurationNotAcceptable) Error() string {
	return fmt.Sprintf("[GET /fido2/configuration][%d] getFido2ConfigurationNotAcceptable  %+v", 406, o.Payload)
}

func (o *GetFIDO2ConfigurationNotAcceptable) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *GetFIDO2ConfigurationNotAcceptable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetFIDO2ConfigurationInternalServerError creates a GetFIDO2ConfigurationInternalServerError with default headers values
func NewGetFIDO2ConfigurationInternalServerError() *GetFIDO2ConfigurationInternalServerError {
	return &GetFIDO2ConfigurationInternalServerError{}
}

/*GetFIDO2ConfigurationInternalServerError handles this case with default header values.

Internal error occured. Please check log file for details.
*/
type GetFIDO2ConfigurationInternalServerError struct {
	Payload *models.ErrorResponse
}

func (o *GetFIDO2ConfigurationInternalServerError) Error() string {
	return fmt.Sprintf("[GET /fido2/configuration][%d] getFido2ConfigurationInternalServerError  %+v", 500, o.Payload)
}

func (o *GetFIDO2ConfigurationInternalServerError) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *GetFIDO2ConfigurationInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
