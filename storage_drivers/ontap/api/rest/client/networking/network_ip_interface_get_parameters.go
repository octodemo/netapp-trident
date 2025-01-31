// Code generated by go-swagger; DO NOT EDIT.

package networking

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
	"github.com/go-openapi/swag"
)

// NewNetworkIPInterfaceGetParams creates a new NetworkIPInterfaceGetParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewNetworkIPInterfaceGetParams() *NetworkIPInterfaceGetParams {
	return &NetworkIPInterfaceGetParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewNetworkIPInterfaceGetParamsWithTimeout creates a new NetworkIPInterfaceGetParams object
// with the ability to set a timeout on a request.
func NewNetworkIPInterfaceGetParamsWithTimeout(timeout time.Duration) *NetworkIPInterfaceGetParams {
	return &NetworkIPInterfaceGetParams{
		timeout: timeout,
	}
}

// NewNetworkIPInterfaceGetParamsWithContext creates a new NetworkIPInterfaceGetParams object
// with the ability to set a context for a request.
func NewNetworkIPInterfaceGetParamsWithContext(ctx context.Context) *NetworkIPInterfaceGetParams {
	return &NetworkIPInterfaceGetParams{
		Context: ctx,
	}
}

// NewNetworkIPInterfaceGetParamsWithHTTPClient creates a new NetworkIPInterfaceGetParams object
// with the ability to set a custom HTTPClient for a request.
func NewNetworkIPInterfaceGetParamsWithHTTPClient(client *http.Client) *NetworkIPInterfaceGetParams {
	return &NetworkIPInterfaceGetParams{
		HTTPClient: client,
	}
}

/* NetworkIPInterfaceGetParams contains all the parameters to send to the API endpoint
   for the network ip interface get operation.

   Typically these are written to a http.Request.
*/
type NetworkIPInterfaceGetParams struct {

	/* Fields.

	   Specify the fields to return.
	*/
	FieldsQueryParameter []string

	/* UUID.

	   IP interface UUID
	*/
	UUIDPathParameter string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the network ip interface get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *NetworkIPInterfaceGetParams) WithDefaults() *NetworkIPInterfaceGetParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the network ip interface get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *NetworkIPInterfaceGetParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the network ip interface get params
func (o *NetworkIPInterfaceGetParams) WithTimeout(timeout time.Duration) *NetworkIPInterfaceGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the network ip interface get params
func (o *NetworkIPInterfaceGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the network ip interface get params
func (o *NetworkIPInterfaceGetParams) WithContext(ctx context.Context) *NetworkIPInterfaceGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the network ip interface get params
func (o *NetworkIPInterfaceGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the network ip interface get params
func (o *NetworkIPInterfaceGetParams) WithHTTPClient(client *http.Client) *NetworkIPInterfaceGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the network ip interface get params
func (o *NetworkIPInterfaceGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFieldsQueryParameter adds the fields to the network ip interface get params
func (o *NetworkIPInterfaceGetParams) WithFieldsQueryParameter(fields []string) *NetworkIPInterfaceGetParams {
	o.SetFieldsQueryParameter(fields)
	return o
}

// SetFieldsQueryParameter adds the fields to the network ip interface get params
func (o *NetworkIPInterfaceGetParams) SetFieldsQueryParameter(fields []string) {
	o.FieldsQueryParameter = fields
}

// WithUUIDPathParameter adds the uuid to the network ip interface get params
func (o *NetworkIPInterfaceGetParams) WithUUIDPathParameter(uuid string) *NetworkIPInterfaceGetParams {
	o.SetUUIDPathParameter(uuid)
	return o
}

// SetUUIDPathParameter adds the uuid to the network ip interface get params
func (o *NetworkIPInterfaceGetParams) SetUUIDPathParameter(uuid string) {
	o.UUIDPathParameter = uuid
}

// WriteToRequest writes these params to a swagger request
func (o *NetworkIPInterfaceGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.FieldsQueryParameter != nil {

		// binding items for fields
		joinedFields := o.bindParamFields(reg)

		// query array param fields
		if err := r.SetQueryParam("fields", joinedFields...); err != nil {
			return err
		}
	}

	// path param uuid
	if err := r.SetPathParam("uuid", o.UUIDPathParameter); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindParamNetworkIPInterfaceGet binds the parameter fields
func (o *NetworkIPInterfaceGetParams) bindParamFields(formats strfmt.Registry) []string {
	fieldsIR := o.FieldsQueryParameter

	var fieldsIC []string
	for _, fieldsIIR := range fieldsIR { // explode []string

		fieldsIIV := fieldsIIR // string as string
		fieldsIC = append(fieldsIC, fieldsIIV)
	}

	// items.CollectionFormat: "csv"
	fieldsIS := swag.JoinByFormat(fieldsIC, "csv")

	return fieldsIS
}
