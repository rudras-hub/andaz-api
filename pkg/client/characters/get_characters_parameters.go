// Code generated by go-swagger; DO NOT EDIT.

package characters

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

// NewGetCharactersParams creates a new GetCharactersParams object
// with the default values initialized.
func NewGetCharactersParams() *GetCharactersParams {

	return &GetCharactersParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetCharactersParamsWithTimeout creates a new GetCharactersParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetCharactersParamsWithTimeout(timeout time.Duration) *GetCharactersParams {

	return &GetCharactersParams{

		timeout: timeout,
	}
}

// NewGetCharactersParamsWithContext creates a new GetCharactersParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetCharactersParamsWithContext(ctx context.Context) *GetCharactersParams {

	return &GetCharactersParams{

		Context: ctx,
	}
}

// NewGetCharactersParamsWithHTTPClient creates a new GetCharactersParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetCharactersParamsWithHTTPClient(client *http.Client) *GetCharactersParams {

	return &GetCharactersParams{
		HTTPClient: client,
	}
}

/*GetCharactersParams contains all the parameters to send to the API endpoint
for the get characters operation typically these are written to a http.Request
*/
type GetCharactersParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get characters params
func (o *GetCharactersParams) WithTimeout(timeout time.Duration) *GetCharactersParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get characters params
func (o *GetCharactersParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get characters params
func (o *GetCharactersParams) WithContext(ctx context.Context) *GetCharactersParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get characters params
func (o *GetCharactersParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get characters params
func (o *GetCharactersParams) WithHTTPClient(client *http.Client) *GetCharactersParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get characters params
func (o *GetCharactersParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *GetCharactersParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}