// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"
)

// NewListSitesParams creates a new ListSitesParams object
// with the default values initialized.
func NewListSitesParams() *ListSitesParams {

	return &ListSitesParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewListSitesParamsWithTimeout creates a new ListSitesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewListSitesParamsWithTimeout(timeout time.Duration) *ListSitesParams {

	return &ListSitesParams{

		timeout: timeout,
	}
}

// NewListSitesParamsWithContext creates a new ListSitesParams object
// with the default values initialized, and the ability to set a context for a request
func NewListSitesParamsWithContext(ctx context.Context) *ListSitesParams {

	return &ListSitesParams{

		Context: ctx,
	}
}

// NewListSitesParamsWithHTTPClient creates a new ListSitesParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewListSitesParamsWithHTTPClient(client *http.Client) *ListSitesParams {

	return &ListSitesParams{
		HTTPClient: client,
	}
}

/*ListSitesParams contains all the parameters to send to the API endpoint
for the list sites operation typically these are written to a http.Request
*/
type ListSitesParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the list sites params
func (o *ListSitesParams) WithTimeout(timeout time.Duration) *ListSitesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the list sites params
func (o *ListSitesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the list sites params
func (o *ListSitesParams) WithContext(ctx context.Context) *ListSitesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the list sites params
func (o *ListSitesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the list sites params
func (o *ListSitesParams) WithHTTPClient(client *http.Client) *ListSitesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the list sites params
func (o *ListSitesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *ListSitesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
