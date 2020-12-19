// Code generated by go-swagger; DO NOT EDIT.

package uma_user_managed_access

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/timdrysdale/gojans/pkg/jansAuth/models"
)

// UMA2ConfigurationReader is a Reader for the UMA2Configuration structure.
type UMA2ConfigurationReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UMA2ConfigurationReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUMA2ConfigurationOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 500:
		result := NewUMA2ConfigurationInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewUMA2ConfigurationOK creates a UMA2ConfigurationOK with default headers values
func NewUMA2ConfigurationOK() *UMA2ConfigurationOK {
	return &UMA2ConfigurationOK{}
}

/*UMA2ConfigurationOK handles this case with default header values.

OK
*/
type UMA2ConfigurationOK struct {
	Payload *models.UMA2ConfigurationResponse
}

func (o *UMA2ConfigurationOK) Error() string {
	return fmt.Sprintf("[GET /uma2-configuration][%d] uma2ConfigurationOK  %+v", 200, o.Payload)
}

func (o *UMA2ConfigurationOK) GetPayload() *models.UMA2ConfigurationResponse {
	return o.Payload
}

func (o *UMA2ConfigurationOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.UMA2ConfigurationResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUMA2ConfigurationInternalServerError creates a UMA2ConfigurationInternalServerError with default headers values
func NewUMA2ConfigurationInternalServerError() *UMA2ConfigurationInternalServerError {
	return &UMA2ConfigurationInternalServerError{}
}

/*UMA2ConfigurationInternalServerError handles this case with default header values.

Invalid parameters provided to endpoint.
*/
type UMA2ConfigurationInternalServerError struct {
	Payload *models.ErrorResponse
}

func (o *UMA2ConfigurationInternalServerError) Error() string {
	return fmt.Sprintf("[GET /uma2-configuration][%d] uma2ConfigurationInternalServerError  %+v", 500, o.Payload)
}

func (o *UMA2ConfigurationInternalServerError) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *UMA2ConfigurationInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
