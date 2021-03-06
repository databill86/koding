package j_machine

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

	"koding/remoteapi/models"
)

// NewJMachineReviveUsersParams creates a new JMachineReviveUsersParams object
// with the default values initialized.
func NewJMachineReviveUsersParams() *JMachineReviveUsersParams {
	var ()
	return &JMachineReviveUsersParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewJMachineReviveUsersParamsWithTimeout creates a new JMachineReviveUsersParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewJMachineReviveUsersParamsWithTimeout(timeout time.Duration) *JMachineReviveUsersParams {
	var ()
	return &JMachineReviveUsersParams{

		timeout: timeout,
	}
}

// NewJMachineReviveUsersParamsWithContext creates a new JMachineReviveUsersParams object
// with the default values initialized, and the ability to set a context for a request
func NewJMachineReviveUsersParamsWithContext(ctx context.Context) *JMachineReviveUsersParams {
	var ()
	return &JMachineReviveUsersParams{

		Context: ctx,
	}
}

/*JMachineReviveUsersParams contains all the parameters to send to the API endpoint
for the j machine revive users operation typically these are written to a http.Request
*/
type JMachineReviveUsersParams struct {

	/*Body
	  body of the request

	*/
	Body models.DefaultSelector
	/*ID
	  Mongo ID of target instance

	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the j machine revive users params
func (o *JMachineReviveUsersParams) WithTimeout(timeout time.Duration) *JMachineReviveUsersParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the j machine revive users params
func (o *JMachineReviveUsersParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the j machine revive users params
func (o *JMachineReviveUsersParams) WithContext(ctx context.Context) *JMachineReviveUsersParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the j machine revive users params
func (o *JMachineReviveUsersParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithBody adds the body to the j machine revive users params
func (o *JMachineReviveUsersParams) WithBody(body models.DefaultSelector) *JMachineReviveUsersParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the j machine revive users params
func (o *JMachineReviveUsersParams) SetBody(body models.DefaultSelector) {
	o.Body = body
}

// WithID adds the id to the j machine revive users params
func (o *JMachineReviveUsersParams) WithID(id string) *JMachineReviveUsersParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the j machine revive users params
func (o *JMachineReviveUsersParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *JMachineReviveUsersParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	r.SetTimeout(o.timeout)
	var res []error

	if err := r.SetBodyParam(o.Body); err != nil {
		return err
	}

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
