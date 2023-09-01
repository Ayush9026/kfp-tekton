// Code generated by go-swagger; DO NOT EDIT.

package pipeline_upload_model

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/validate"
)

// V1ResourceType v1 resource type
//
// swagger:model v1ResourceType
type V1ResourceType string

func NewV1ResourceType(value V1ResourceType) *V1ResourceType {
	return &value
}

// Pointer returns a pointer to a freshly-allocated V1ResourceType.
func (m V1ResourceType) Pointer() *V1ResourceType {
	return &m
}

const (

	// V1ResourceTypeUNKNOWNRESOURCETYPE captures enum value "UNKNOWN_RESOURCE_TYPE"
	V1ResourceTypeUNKNOWNRESOURCETYPE V1ResourceType = "UNKNOWN_RESOURCE_TYPE"

	// V1ResourceTypeEXPERIMENT captures enum value "EXPERIMENT"
	V1ResourceTypeEXPERIMENT V1ResourceType = "EXPERIMENT"

	// V1ResourceTypeJOB captures enum value "JOB"
	V1ResourceTypeJOB V1ResourceType = "JOB"

	// V1ResourceTypePIPELINE captures enum value "PIPELINE"
	V1ResourceTypePIPELINE V1ResourceType = "PIPELINE"

	// V1ResourceTypePIPELINEVERSION captures enum value "PIPELINE_VERSION"
	V1ResourceTypePIPELINEVERSION V1ResourceType = "PIPELINE_VERSION"

	// V1ResourceTypeNAMESPACE captures enum value "NAMESPACE"
	V1ResourceTypeNAMESPACE V1ResourceType = "NAMESPACE"
)

// for schema
var v1ResourceTypeEnum []interface{}

func init() {
	var res []V1ResourceType
	if err := json.Unmarshal([]byte(`["UNKNOWN_RESOURCE_TYPE","EXPERIMENT","JOB","PIPELINE","PIPELINE_VERSION","NAMESPACE"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		v1ResourceTypeEnum = append(v1ResourceTypeEnum, v)
	}
}

func (m V1ResourceType) validateV1ResourceTypeEnum(path, location string, value V1ResourceType) error {
	if err := validate.EnumCase(path, location, value, v1ResourceTypeEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this v1 resource type
func (m V1ResourceType) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateV1ResourceTypeEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this v1 resource type based on context it is used
func (m V1ResourceType) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}
