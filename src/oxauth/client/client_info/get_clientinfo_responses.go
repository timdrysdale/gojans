// Code generated by go-swagger; DO NOT EDIT.

package client_info

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/timdrysdale/gojans/go-swagger/gluu-4.2.1/oxauth/models"
)

// GetClientinfoReader is a Reader for the GetClientinfo structure.
type GetClientinfoReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetClientinfoReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetClientinfoOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetClientinfoBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetClientinfoOK creates a GetClientinfoOK with default headers values
func NewGetClientinfoOK() *GetClientinfoOK {
	return &GetClientinfoOK{}
}

/*GetClientinfoOK handles this case with default header values.

OK
*/
type GetClientinfoOK struct {
	Payload *models.ClientInfoResponse
}

func (o *GetClientinfoOK) Error() string {
	return fmt.Sprintf("[GET /clientinfo][%d] getClientinfoOK  %+v", 200, o.Payload)
}

func (o *GetClientinfoOK) GetPayload() *models.ClientInfoResponse {
	return o.Payload
}

func (o *GetClientinfoOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ClientInfoResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetClientinfoBadRequest creates a GetClientinfoBadRequest with default headers values
func NewGetClientinfoBadRequest() *GetClientinfoBadRequest {
	return &GetClientinfoBadRequest{}
}

/*GetClientinfoBadRequest handles this case with default header values.

Invalid Request are provided to endpoint.
*/
type GetClientinfoBadRequest struct {
	Payload *models.ErrorResponse
}

func (o *GetClientinfoBadRequest) Error() string {
	return fmt.Sprintf("[GET /clientinfo][%d] getClientinfoBadRequest  %+v", 400, o.Payload)
}

func (o *GetClientinfoBadRequest) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *GetClientinfoBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}