// Code generated by go generate; DO NOT EDIT.
package containers

import (
	"net/url"

	"github.com/containers/podman/v3/pkg/bindings/internal/util"
)

// Changed returns true if named field has been set
func (o *ResizeTTYOptions) Changed(fieldName string) bool {
	return util.Changed(o, fieldName)
}

// ToParams formats struct fields to be passed to API service
func (o *ResizeTTYOptions) ToParams() (url.Values, error) {
	return util.ToParams(o)
}

// WithHeight set field Height to given value
func (o *ResizeTTYOptions) WithHeight(value int) *ResizeTTYOptions {
	o.Height = &value
	return o
}

// GetHeight returns value of field Height
func (o *ResizeTTYOptions) GetHeight() int {
	if o.Height == nil {
		var z int
		return z
	}
	return *o.Height
}

// WithWidth set field Width to given value
func (o *ResizeTTYOptions) WithWidth(value int) *ResizeTTYOptions {
	o.Width = &value
	return o
}

// GetWidth returns value of field Width
func (o *ResizeTTYOptions) GetWidth() int {
	if o.Width == nil {
		var z int
		return z
	}
	return *o.Width
}

// WithRunning set field Running to given value
func (o *ResizeTTYOptions) WithRunning(value bool) *ResizeTTYOptions {
	o.Running = &value
	return o
}

// GetRunning returns value of field Running
func (o *ResizeTTYOptions) GetRunning() bool {
	if o.Running == nil {
		var z bool
		return z
	}
	return *o.Running
}