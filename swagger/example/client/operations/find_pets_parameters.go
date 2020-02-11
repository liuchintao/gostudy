// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"
)

// NewFindPetsParams creates a new FindPetsParams object
// with the default values initialized.
func NewFindPetsParams() *FindPetsParams {
	var ()
	return &FindPetsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewFindPetsParamsWithTimeout creates a new FindPetsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewFindPetsParamsWithTimeout(timeout time.Duration) *FindPetsParams {
	var ()
	return &FindPetsParams{

		timeout: timeout,
	}
}

// NewFindPetsParamsWithContext creates a new FindPetsParams object
// with the default values initialized, and the ability to set a context for a request
func NewFindPetsParamsWithContext(ctx context.Context) *FindPetsParams {
	var ()
	return &FindPetsParams{

		Context: ctx,
	}
}

// NewFindPetsParamsWithHTTPClient creates a new FindPetsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewFindPetsParamsWithHTTPClient(client *http.Client) *FindPetsParams {
	var ()
	return &FindPetsParams{
		HTTPClient: client,
	}
}

/*FindPetsParams contains all the parameters to send to the API endpoint
for the find pets operation typically these are written to a http.Request
*/
type FindPetsParams struct {

	/*Limit
	  maximum number of results to return

	*/
	Limit *int32
	/*Tags
	  tags to filter by

	*/
	Tags []string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the find pets params
func (o *FindPetsParams) WithTimeout(timeout time.Duration) *FindPetsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the find pets params
func (o *FindPetsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the find pets params
func (o *FindPetsParams) WithContext(ctx context.Context) *FindPetsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the find pets params
func (o *FindPetsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the find pets params
func (o *FindPetsParams) WithHTTPClient(client *http.Client) *FindPetsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the find pets params
func (o *FindPetsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithLimit adds the limit to the find pets params
func (o *FindPetsParams) WithLimit(limit *int32) *FindPetsParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the find pets params
func (o *FindPetsParams) SetLimit(limit *int32) {
	o.Limit = limit
}

// WithTags adds the tags to the find pets params
func (o *FindPetsParams) WithTags(tags []string) *FindPetsParams {
	o.SetTags(tags)
	return o
}

// SetTags adds the tags to the find pets params
func (o *FindPetsParams) SetTags(tags []string) {
	o.Tags = tags
}

// WriteToRequest writes these params to a swagger request
func (o *FindPetsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Limit != nil {

		// query param limit
		var qrLimit int32
		if o.Limit != nil {
			qrLimit = *o.Limit
		}
		qLimit := swag.FormatInt32(qrLimit)
		if qLimit != "" {
			if err := r.SetQueryParam("limit", qLimit); err != nil {
				return err
			}
		}

	}

	valuesTags := o.Tags

	joinedTags := swag.JoinByFormat(valuesTags, "csv")
	// query array param tags
	if err := r.SetQueryParam("tags", joinedTags...); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}