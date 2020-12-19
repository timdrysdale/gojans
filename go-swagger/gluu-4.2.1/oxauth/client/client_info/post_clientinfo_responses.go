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

// PostClientinfoReader is a Reader for the PostClientinfo structure.
type PostClientinfoReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostClientinfoReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostClientinfoOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPostClientinfoBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPostClientinfoOK creates a PostClientinfoOK with default headers values
func NewPostClientinfoOK() *PostClientinfoOK {
	return &PostClientinfoOK{}
}

/*PostClientinfoOK handles this case with default header values.

OK
*/
type PostClientinfoOK struct {
	Payload *models.ClientInfoResponse
}

func (o *PostClientinfoOK) Error() string {
	return fmt.Sprintf("[POST /clientinfo][%d] postClientinfoOK  %+v", 200, o.Payload)
}

func (o *PostClientinfoOK) GetPayload() *models.ClientInfoResponse {
	return o.Payload
}

func (o *PostClientinfoOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ClientInfoResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostClientinfoBadRequest creates a PostClientinfoBadRequest with default headers values
func NewPostClientinfoBadRequest() *PostClientinfoBadRequest {
	return &PostClientinfoBadRequest{}
}

/*PostClientinfoBadRequest handles this case with default header values.

Invalid Request are provided to endpoint.
*/
type PostClientinfoBadRequest struct {
	Payload *models.ErrorResponse
}

func (o *PostClientinfoBadRequest) Error() string {
	return fmt.Sprintf("[POST /clientinfo][%d] postClientinfoBadRequest  %+v", 400, o.Payload)
}

func (o *PostClientinfoBadRequest) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *PostClientinfoBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}