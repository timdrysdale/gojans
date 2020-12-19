// Code generated by go-swagger; DO NOT EDIT.

package fido2

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/timdrysdale/gojans/src/jansAuth/models"
)

// AttestationResultReader is a Reader for the AttestationResult structure.
type AttestationResultReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AttestationResultReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewAttestationResultOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 500:
		result := NewAttestationResultInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewAttestationResultOK creates a AttestationResultOK with default headers values
func NewAttestationResultOK() *AttestationResultOK {
	return &AttestationResultOK{}
}

/*AttestationResultOK handles this case with default header values.

Invalid details provided hence access denied.
*/
type AttestationResultOK struct {
	Payload *models.ErrorResponse
}

func (o *AttestationResultOK) Error() string {
	return fmt.Sprintf("[POST /fido2/attestation/result][%d] attestationResultOK  %+v", 200, o.Payload)
}

func (o *AttestationResultOK) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *AttestationResultOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAttestationResultInternalServerError creates a AttestationResultInternalServerError with default headers values
func NewAttestationResultInternalServerError() *AttestationResultInternalServerError {
	return &AttestationResultInternalServerError{}
}

/*AttestationResultInternalServerError handles this case with default header values.

Internal error occured. Please check log file for details.
*/
type AttestationResultInternalServerError struct {
	Payload *models.ErrorResponse
}

func (o *AttestationResultInternalServerError) Error() string {
	return fmt.Sprintf("[POST /fido2/attestation/result][%d] attestationResultInternalServerError  %+v", 500, o.Payload)
}

func (o *AttestationResultInternalServerError) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *AttestationResultInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}