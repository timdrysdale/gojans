// Code generated by go-swagger; DO NOT EDIT.

package uma_user_managed_access

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/timdrysdale/gojans/go-swagger/gluu-4.2.1/oxauth/models"
)

// HostRsrcPrReader is a Reader for the HostRsrcPr structure.
type HostRsrcPrReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *HostRsrcPrReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewHostRsrcPrCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 500:
		result := NewHostRsrcPrInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewHostRsrcPrCreated creates a HostRsrcPrCreated with default headers values
func NewHostRsrcPrCreated() *HostRsrcPrCreated {
	return &HostRsrcPrCreated{}
}

/*HostRsrcPrCreated handles this case with default header values.

OK
*/
type HostRsrcPrCreated struct {
	Payload []*models.UMAPermissionList
}

func (o *HostRsrcPrCreated) Error() string {
	return fmt.Sprintf("[POST /host/rsrc_pr][%d] hostRsrcPrCreated  %+v", 201, o.Payload)
}

func (o *HostRsrcPrCreated) GetPayload() []*models.UMAPermissionList {
	return o.Payload
}

func (o *HostRsrcPrCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewHostRsrcPrInternalServerError creates a HostRsrcPrInternalServerError with default headers values
func NewHostRsrcPrInternalServerError() *HostRsrcPrInternalServerError {
	return &HostRsrcPrInternalServerError{}
}

/*HostRsrcPrInternalServerError handles this case with default header values.

Invalid parameters provided to endpoint.
*/
type HostRsrcPrInternalServerError struct {
	Payload *models.ErrorResponse
}

func (o *HostRsrcPrInternalServerError) Error() string {
	return fmt.Sprintf("[POST /host/rsrc_pr][%d] hostRsrcPrInternalServerError  %+v", 500, o.Payload)
}

func (o *HostRsrcPrInternalServerError) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *HostRsrcPrInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
