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

// PipelineServiceListPipelinesReader is a Reader for the PipelineServiceListPipelines structure.
type PipelineServiceListPipelinesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PipelineServiceListPipelinesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPipelineServiceListPipelinesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewPipelineServiceListPipelinesDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPipelineServiceListPipelinesOK creates a PipelineServiceListPipelinesOK with default headers values
func NewPipelineServiceListPipelinesOK() *PipelineServiceListPipelinesOK {
	return &PipelineServiceListPipelinesOK{}
}

/*
PipelineServiceListPipelinesOK describes a response with status code 200, with default header values.

A successful response.
*/
type PipelineServiceListPipelinesOK struct {
	Payload *pipeline_model.V1ListPipelinesResponse
}

// IsSuccess returns true when this pipeline service list pipelines o k response has a 2xx status code
func (o *PipelineServiceListPipelinesOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this pipeline service list pipelines o k response has a 3xx status code
func (o *PipelineServiceListPipelinesOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pipeline service list pipelines o k response has a 4xx status code
func (o *PipelineServiceListPipelinesOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this pipeline service list pipelines o k response has a 5xx status code
func (o *PipelineServiceListPipelinesOK) IsServerError() bool {
	return false
}

// IsCode returns true when this pipeline service list pipelines o k response a status code equal to that given
func (o *PipelineServiceListPipelinesOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the pipeline service list pipelines o k response
func (o *PipelineServiceListPipelinesOK) Code() int {
	return 200
}

func (o *PipelineServiceListPipelinesOK) Error() string {
	return fmt.Sprintf("[GET /apis/v1/pipelines][%d] pipelineServiceListPipelinesOK  %+v", 200, o.Payload)
}

func (o *PipelineServiceListPipelinesOK) String() string {
	return fmt.Sprintf("[GET /apis/v1/pipelines][%d] pipelineServiceListPipelinesOK  %+v", 200, o.Payload)
}

func (o *PipelineServiceListPipelinesOK) GetPayload() *pipeline_model.V1ListPipelinesResponse {
	return o.Payload
}

func (o *PipelineServiceListPipelinesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(pipeline_model.V1ListPipelinesResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPipelineServiceListPipelinesDefault creates a PipelineServiceListPipelinesDefault with default headers values
func NewPipelineServiceListPipelinesDefault(code int) *PipelineServiceListPipelinesDefault {
	return &PipelineServiceListPipelinesDefault{
		_statusCode: code,
	}
}

/*
PipelineServiceListPipelinesDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type PipelineServiceListPipelinesDefault struct {
	_statusCode int

	Payload *pipeline_model.GooglerpcStatus
}

// IsSuccess returns true when this pipeline service list pipelines default response has a 2xx status code
func (o *PipelineServiceListPipelinesDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this pipeline service list pipelines default response has a 3xx status code
func (o *PipelineServiceListPipelinesDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this pipeline service list pipelines default response has a 4xx status code
func (o *PipelineServiceListPipelinesDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this pipeline service list pipelines default response has a 5xx status code
func (o *PipelineServiceListPipelinesDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this pipeline service list pipelines default response a status code equal to that given
func (o *PipelineServiceListPipelinesDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the pipeline service list pipelines default response
func (o *PipelineServiceListPipelinesDefault) Code() int {
	return o._statusCode
}

func (o *PipelineServiceListPipelinesDefault) Error() string {
	return fmt.Sprintf("[GET /apis/v1/pipelines][%d] PipelineService_ListPipelines default  %+v", o._statusCode, o.Payload)
}

func (o *PipelineServiceListPipelinesDefault) String() string {
	return fmt.Sprintf("[GET /apis/v1/pipelines][%d] PipelineService_ListPipelines default  %+v", o._statusCode, o.Payload)
}

func (o *PipelineServiceListPipelinesDefault) GetPayload() *pipeline_model.GooglerpcStatus {
	return o.Payload
}

func (o *PipelineServiceListPipelinesDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(pipeline_model.GooglerpcStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
