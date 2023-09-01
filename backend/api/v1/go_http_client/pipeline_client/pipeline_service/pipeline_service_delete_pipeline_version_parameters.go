// Code generated by go-swagger; DO NOT EDIT.

package pipeline_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// NewPipelineServiceDeletePipelineVersionParams creates a new PipelineServiceDeletePipelineVersionParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPipelineServiceDeletePipelineVersionParams() *PipelineServiceDeletePipelineVersionParams {
	return &PipelineServiceDeletePipelineVersionParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPipelineServiceDeletePipelineVersionParamsWithTimeout creates a new PipelineServiceDeletePipelineVersionParams object
// with the ability to set a timeout on a request.
func NewPipelineServiceDeletePipelineVersionParamsWithTimeout(timeout time.Duration) *PipelineServiceDeletePipelineVersionParams {
	return &PipelineServiceDeletePipelineVersionParams{
		timeout: timeout,
	}
}

// NewPipelineServiceDeletePipelineVersionParamsWithContext creates a new PipelineServiceDeletePipelineVersionParams object
// with the ability to set a context for a request.
func NewPipelineServiceDeletePipelineVersionParamsWithContext(ctx context.Context) *PipelineServiceDeletePipelineVersionParams {
	return &PipelineServiceDeletePipelineVersionParams{
		Context: ctx,
	}
}

// NewPipelineServiceDeletePipelineVersionParamsWithHTTPClient creates a new PipelineServiceDeletePipelineVersionParams object
// with the ability to set a custom HTTPClient for a request.
func NewPipelineServiceDeletePipelineVersionParamsWithHTTPClient(client *http.Client) *PipelineServiceDeletePipelineVersionParams {
	return &PipelineServiceDeletePipelineVersionParams{
		HTTPClient: client,
	}
}

/*
PipelineServiceDeletePipelineVersionParams contains all the parameters to send to the API endpoint

	for the pipeline service delete pipeline version operation.

	Typically these are written to a http.Request.
*/
type PipelineServiceDeletePipelineVersionParams struct {

	/* VersionID.

	   The ID of the pipeline version to be deleted.
	*/
	VersionID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the pipeline service delete pipeline version params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PipelineServiceDeletePipelineVersionParams) WithDefaults() *PipelineServiceDeletePipelineVersionParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the pipeline service delete pipeline version params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PipelineServiceDeletePipelineVersionParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the pipeline service delete pipeline version params
func (o *PipelineServiceDeletePipelineVersionParams) WithTimeout(timeout time.Duration) *PipelineServiceDeletePipelineVersionParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the pipeline service delete pipeline version params
func (o *PipelineServiceDeletePipelineVersionParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the pipeline service delete pipeline version params
func (o *PipelineServiceDeletePipelineVersionParams) WithContext(ctx context.Context) *PipelineServiceDeletePipelineVersionParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the pipeline service delete pipeline version params
func (o *PipelineServiceDeletePipelineVersionParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the pipeline service delete pipeline version params
func (o *PipelineServiceDeletePipelineVersionParams) WithHTTPClient(client *http.Client) *PipelineServiceDeletePipelineVersionParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the pipeline service delete pipeline version params
func (o *PipelineServiceDeletePipelineVersionParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithVersionID adds the versionID to the pipeline service delete pipeline version params
func (o *PipelineServiceDeletePipelineVersionParams) WithVersionID(versionID string) *PipelineServiceDeletePipelineVersionParams {
	o.SetVersionID(versionID)
	return o
}

// SetVersionID adds the versionId to the pipeline service delete pipeline version params
func (o *PipelineServiceDeletePipelineVersionParams) SetVersionID(versionID string) {
	o.VersionID = versionID
}

// WriteToRequest writes these params to a swagger request
func (o *PipelineServiceDeletePipelineVersionParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param version_id
	if err := r.SetPathParam("version_id", o.VersionID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
