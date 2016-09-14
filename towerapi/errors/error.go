package errors

import (
	"fmt"

	"github.com/hashicorp/go-multierror"
)

// An APIError reports the error caused by an API request
type APIError struct {
	Errors []string `json:"__all__,omitempty"`
	Detail string   `json:"detail,omitempty"`
}

func (e APIError) Error() string {
	var result *multierror.Error
	if e.Detail != "" {
		result = multierror.Append(result, fmt.Errorf(e.Detail))
	}
	for _, err := range e.Errors {
		result = multierror.Append(result, fmt.Errorf(err))
	}
	if result.ErrorOrNil() != nil {
		return result.GoString()
	} else {
		return ""
	}

}

// BuildError returns any non-nil http-related error (creating the request,
// getting the response, decoding) if any. If the decoded apierr is non-zero
// the apierr is returned. Otherwise, no errors occurred, returns nil.
func BuildError(httperr error, apierr *APIError) error {
	var result *multierror.Error
	if httperr != nil {
		result = multierror.Append(result, httperr)
	}
	return result.ErrorOrNil()
}
