// Code generated by go-swagger; DO NOT EDIT.

package nodes

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"
)

// NewAddRemoteRDSNodeParams creates a new AddRemoteRDSNodeParams object
// with the default values initialized.
func NewAddRemoteRDSNodeParams() *AddRemoteRDSNodeParams {
	var ()
	return &AddRemoteRDSNodeParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewAddRemoteRDSNodeParamsWithTimeout creates a new AddRemoteRDSNodeParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewAddRemoteRDSNodeParamsWithTimeout(timeout time.Duration) *AddRemoteRDSNodeParams {
	var ()
	return &AddRemoteRDSNodeParams{

		timeout: timeout,
	}
}

// NewAddRemoteRDSNodeParamsWithContext creates a new AddRemoteRDSNodeParams object
// with the default values initialized, and the ability to set a context for a request
func NewAddRemoteRDSNodeParamsWithContext(ctx context.Context) *AddRemoteRDSNodeParams {
	var ()
	return &AddRemoteRDSNodeParams{

		Context: ctx,
	}
}

// NewAddRemoteRDSNodeParamsWithHTTPClient creates a new AddRemoteRDSNodeParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewAddRemoteRDSNodeParamsWithHTTPClient(client *http.Client) *AddRemoteRDSNodeParams {
	var ()
	return &AddRemoteRDSNodeParams{
		HTTPClient: client,
	}
}

/*AddRemoteRDSNodeParams contains all the parameters to send to the API endpoint
for the add remote RDS node operation typically these are written to a http.Request
*/
type AddRemoteRDSNodeParams struct {

	/*Body*/
	Body AddRemoteRDSNodeBody

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the add remote RDS node params
func (o *AddRemoteRDSNodeParams) WithTimeout(timeout time.Duration) *AddRemoteRDSNodeParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the add remote RDS node params
func (o *AddRemoteRDSNodeParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the add remote RDS node params
func (o *AddRemoteRDSNodeParams) WithContext(ctx context.Context) *AddRemoteRDSNodeParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the add remote RDS node params
func (o *AddRemoteRDSNodeParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the add remote RDS node params
func (o *AddRemoteRDSNodeParams) WithHTTPClient(client *http.Client) *AddRemoteRDSNodeParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the add remote RDS node params
func (o *AddRemoteRDSNodeParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the add remote RDS node params
func (o *AddRemoteRDSNodeParams) WithBody(body AddRemoteRDSNodeBody) *AddRemoteRDSNodeParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the add remote RDS node params
func (o *AddRemoteRDSNodeParams) SetBody(body AddRemoteRDSNodeBody) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *AddRemoteRDSNodeParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if err := r.SetBodyParam(o.Body); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
