package social_message

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

// NewSocialMessageSearchParams creates a new SocialMessageSearchParams object
// with the default values initialized.
func NewSocialMessageSearchParams() *SocialMessageSearchParams {
	var ()
	return &SocialMessageSearchParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewSocialMessageSearchParamsWithTimeout creates a new SocialMessageSearchParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewSocialMessageSearchParamsWithTimeout(timeout time.Duration) *SocialMessageSearchParams {
	var ()
	return &SocialMessageSearchParams{

		timeout: timeout,
	}
}

// NewSocialMessageSearchParamsWithContext creates a new SocialMessageSearchParams object
// with the default values initialized, and the ability to set a context for a request
func NewSocialMessageSearchParamsWithContext(ctx context.Context) *SocialMessageSearchParams {
	var ()
	return &SocialMessageSearchParams{

		Context: ctx,
	}
}

/*SocialMessageSearchParams contains all the parameters to send to the API endpoint
for the social message search operation typically these are written to a http.Request
*/
type SocialMessageSearchParams struct {

	/*Body
	  body of the request

	*/
	Body models.DefaultSelector

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the social message search params
func (o *SocialMessageSearchParams) WithTimeout(timeout time.Duration) *SocialMessageSearchParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the social message search params
func (o *SocialMessageSearchParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the social message search params
func (o *SocialMessageSearchParams) WithContext(ctx context.Context) *SocialMessageSearchParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the social message search params
func (o *SocialMessageSearchParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithBody adds the body to the social message search params
func (o *SocialMessageSearchParams) WithBody(body models.DefaultSelector) *SocialMessageSearchParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the social message search params
func (o *SocialMessageSearchParams) SetBody(body models.DefaultSelector) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *SocialMessageSearchParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	r.SetTimeout(o.timeout)
	var res []error

	if err := r.SetBodyParam(o.Body); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
