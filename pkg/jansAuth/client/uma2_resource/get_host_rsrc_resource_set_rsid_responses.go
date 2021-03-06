// Code generated by go-swagger; DO NOT EDIT.

package uma2_resource

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/timdrysdale/gojans/pkg/jansAuth/models"
)

// GetHostRsrcResourceSetRsidReader is a Reader for the GetHostRsrcResourceSetRsid structure.
type GetHostRsrcResourceSetRsidReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetHostRsrcResourceSetRsidReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetHostRsrcResourceSetRsidOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 500:
		result := NewGetHostRsrcResourceSetRsidInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetHostRsrcResourceSetRsidOK creates a GetHostRsrcResourceSetRsidOK with default headers values
func NewGetHostRsrcResourceSetRsidOK() *GetHostRsrcResourceSetRsidOK {
	return &GetHostRsrcResourceSetRsidOK{}
}

/*GetHostRsrcResourceSetRsidOK handles this case with default header values.

OK
*/
type GetHostRsrcResourceSetRsidOK struct {
	Payload *models.UMAResourceWithID
}

func (o *GetHostRsrcResourceSetRsidOK) Error() string {
	return fmt.Sprintf("[GET /host/rsrc/resource_set/{rsid}][%d] getHostRsrcResourceSetRsidOK  %+v", 200, o.Payload)
}

func (o *GetHostRsrcResourceSetRsidOK) GetPayload() *models.UMAResourceWithID {
	return o.Payload
}

func (o *GetHostRsrcResourceSetRsidOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.UMAResourceWithID)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetHostRsrcResourceSetRsidInternalServerError creates a GetHostRsrcResourceSetRsidInternalServerError with default headers values
func NewGetHostRsrcResourceSetRsidInternalServerError() *GetHostRsrcResourceSetRsidInternalServerError {
	return &GetHostRsrcResourceSetRsidInternalServerError{}
}

/*GetHostRsrcResourceSetRsidInternalServerError handles this case with default header values.

Invalid parameters provided to endpoint.
*/
type GetHostRsrcResourceSetRsidInternalServerError struct {
	Payload *models.ErrorResponse
}

func (o *GetHostRsrcResourceSetRsidInternalServerError) Error() string {
	return fmt.Sprintf("[GET /host/rsrc/resource_set/{rsid}][%d] getHostRsrcResourceSetRsidInternalServerError  %+v", 500, o.Payload)
}

func (o *GetHostRsrcResourceSetRsidInternalServerError) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *GetHostRsrcResourceSetRsidInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
