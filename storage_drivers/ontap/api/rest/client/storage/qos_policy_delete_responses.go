// Code generated by go-swagger; DO NOT EDIT.

package storage

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/netapp/trident/storage_drivers/ontap/api/rest/models"
)

// QosPolicyDeleteReader is a Reader for the QosPolicyDelete structure.
type QosPolicyDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *QosPolicyDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 202:
		result := NewQosPolicyDeleteAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewQosPolicyDeleteDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewQosPolicyDeleteAccepted creates a QosPolicyDeleteAccepted with default headers values
func NewQosPolicyDeleteAccepted() *QosPolicyDeleteAccepted {
	return &QosPolicyDeleteAccepted{}
}

/* QosPolicyDeleteAccepted describes a response with status code 202, with default header values.

Accepted
*/
type QosPolicyDeleteAccepted struct {
	Payload *models.JobLinkResponse
}

func (o *QosPolicyDeleteAccepted) Error() string {
	return fmt.Sprintf("[DELETE /storage/qos/policies/{uuid}][%d] qosPolicyDeleteAccepted  %+v", 202, o.Payload)
}
func (o *QosPolicyDeleteAccepted) GetPayload() *models.JobLinkResponse {
	return o.Payload
}

func (o *QosPolicyDeleteAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.JobLinkResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewQosPolicyDeleteDefault creates a QosPolicyDeleteDefault with default headers values
func NewQosPolicyDeleteDefault(code int) *QosPolicyDeleteDefault {
	return &QosPolicyDeleteDefault{
		_statusCode: code,
	}
}

/* QosPolicyDeleteDefault describes a response with status code -1, with default header values.

Error
*/
type QosPolicyDeleteDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the qos policy delete default response
func (o *QosPolicyDeleteDefault) Code() int {
	return o._statusCode
}

func (o *QosPolicyDeleteDefault) Error() string {
	return fmt.Sprintf("[DELETE /storage/qos/policies/{uuid}][%d] qos_policy_delete default  %+v", o._statusCode, o.Payload)
}
func (o *QosPolicyDeleteDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *QosPolicyDeleteDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
