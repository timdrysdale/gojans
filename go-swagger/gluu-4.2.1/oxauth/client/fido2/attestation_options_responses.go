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

// AttestationOptionsReader is a Reader for the AttestationOptions structure.
type AttestationOptionsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AttestationOptionsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewAttestationOptionsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewAttestationOptionsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewAttestationOptionsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewAttestationOptionsOK creates a AttestationOptionsOK with default headers values
func NewAttestationOptionsOK() *AttestationOptionsOK {
	return &AttestationOptionsOK{}
}

/*AttestationOptionsOK handles this case with default header values.

OK
*/
type AttestationOptionsOK struct {
	Payload *models.CredentialCreationOptions
}

func (o *AttestationOptionsOK) Error() string {
	return fmt.Sprintf("[POST /fido2/attestation/options][%d] attestationOptionsOK  %+v", 200, o.Payload)
}

func (o *AttestationOptionsOK) GetPayload() *models.CredentialCreationOptions {
	return o.Payload
}

func (o *AttestationOptionsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CredentialCreationOptions)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAttestationOptionsForbidden creates a AttestationOptionsForbidden with default headers values
func NewAttestationOptionsForbidden() *AttestationOptionsForbidden {
	return &AttestationOptionsForbidden{}
}

/*AttestationOptionsForbidden handles this case with default header values.

Invalid details provided hence access denied.
*/
type AttestationOptionsForbidden struct {
	Payload *models.ErrorResponse
}

func (o *AttestationOptionsForbidden) Error() string {
	return fmt.Sprintf("[POST /fido2/attestation/options][%d] attestationOptionsForbidden  %+v", 403, o.Payload)
}

func (o *AttestationOptionsForbidden) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *AttestationOptionsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAttestationOptionsInternalServerError creates a AttestationOptionsInternalServerError with default headers values
func NewAttestationOptionsInternalServerError() *AttestationOptionsInternalServerError {
	return &AttestationOptionsInternalServerError{}
}

/*AttestationOptionsInternalServerError handles this case with default header values.

Internal error occured. Please check log file for details.
*/
type AttestationOptionsInternalServerError struct {
	Payload *models.ErrorResponse
}

func (o *AttestationOptionsInternalServerError) Error() string {
	return fmt.Sprintf("[POST /fido2/attestation/options][%d] attestationOptionsInternalServerError  %+v", 500, o.Payload)
}

func (o *AttestationOptionsInternalServerError) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *AttestationOptionsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
