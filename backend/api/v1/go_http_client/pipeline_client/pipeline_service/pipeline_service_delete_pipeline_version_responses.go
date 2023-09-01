// Code generated by go-swagger; DO NOT EDIT.

package pipeline_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/kubeflow/pipelines/backend/api/v1/go_http_client/pipeline_model"
)

// PipelineServiceDeletePipelineVersionReader is a Reader for the PipelineServiceDeletePipelineVersion structure.
type PipelineServiceDeletePipelineVersionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PipelineServiceDeletePipelineVersionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPipelineServiceDeletePipelineVersionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewPipelineServiceDeletePipelineVersionDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPipelineServiceDeletePipelineVersionOK creates a PipelineServiceDeletePipelineVersionOK with default headers values
func NewPipelineServiceDeletePipelineVersionOK() *PipelineServiceDeletePipelineVersionOK {
	return &PipelineServiceDeletePipelineVersionOK{}
}

/*
PipelineServiceDeletePipelineVersionOK describes a response with status code 200, with default header values.

A successful response.
*/
type PipelineServiceDeletePipelineVersionOK struct {
	Payload interface{}
}

// IsSuccess returns true when this pipeline service delete pipeline version o k response has a 2xx status code
func (o *PipelineServiceDeletePipelineVersionOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this pipeline service delete pipeline version o k response has a 3xx status code
func (o *PipelineServiceDeletePipelineVersionOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pipeline service delete pipeline version o k response has a 4xx status code
func (o *PipelineServiceDeletePipelineVersionOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this pipeline service delete pipeline version o k response has a 5xx status code
func (o *PipelineServiceDeletePipelineVersionOK) IsServerError() bool {
	return false
}

// IsCode returns true when this pipeline service delete pipeline version o k response a status code equal to that given
func (o *PipelineServiceDeletePipelineVersionOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the pipeline service delete pipeline version o k response
func (o *PipelineServiceDeletePipelineVersionOK) Code() int {
	return 200
}

func (o *PipelineServiceDeletePipelineVersionOK) Error() string {
	return fmt.Sprintf("[DELETE /apis/v1/pipeline_versions/{version_id}][%d] pipelineServiceDeletePipelineVersionOK  %+v", 200, o.Payload)
}

func (o *PipelineServiceDeletePipelineVersionOK) String() string {
	return fmt.Sprintf("[DELETE /apis/v1/pipeline_versions/{version_id}][%d] pipelineServiceDeletePipelineVersionOK  %+v", 200, o.Payload)
}

func (o *PipelineServiceDeletePipelineVersionOK) GetPayload() interface{} {
	return o.Payload
}

func (o *PipelineServiceDeletePipelineVersionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPipelineServiceDeletePipelineVersionDefault creates a PipelineServiceDeletePipelineVersionDefault with default headers values
func NewPipelineServiceDeletePipelineVersionDefault(code int) *PipelineServiceDeletePipelineVersionDefault {
	return &PipelineServiceDeletePipelineVersionDefault{
		_statusCode: code,
	}
}

/*
PipelineServiceDeletePipelineVersionDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type PipelineServiceDeletePipelineVersionDefault struct {
	_statusCode int

	Payload *pipeline_model.GooglerpcStatus
}

// IsSuccess returns true when this pipeline service delete pipeline version default response has a 2xx status code
func (o *PipelineServiceDeletePipelineVersionDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this pipeline service delete pipeline version default response has a 3xx status code
func (o *PipelineServiceDeletePipelineVersionDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this pipeline service delete pipeline version default response has a 4xx status code
func (o *PipelineServiceDeletePipelineVersionDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this pipeline service delete pipeline version default response has a 5xx status code
func (o *PipelineServiceDeletePipelineVersionDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this pipeline service delete pipeline version default response a status code equal to that given
func (o *PipelineServiceDeletePipelineVersionDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the pipeline service delete pipeline version default response
func (o *PipelineServiceDeletePipelineVersionDefault) Code() int {
	return o._statusCode
}

func (o *PipelineServiceDeletePipelineVersionDefault) Error() string {
	return fmt.Sprintf("[DELETE /apis/v1/pipeline_versions/{version_id}][%d] PipelineService_DeletePipelineVersion default  %+v", o._statusCode, o.Payload)
}

func (o *PipelineServiceDeletePipelineVersionDefault) String() string {
	return fmt.Sprintf("[DELETE /apis/v1/pipeline_versions/{version_id}][%d] PipelineService_DeletePipelineVersion default  %+v", o._statusCode, o.Payload)
}

func (o *PipelineServiceDeletePipelineVersionDefault) GetPayload() *pipeline_model.GooglerpcStatus {
	return o.Payload
}

func (o *PipelineServiceDeletePipelineVersionDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(pipeline_model.GooglerpcStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
