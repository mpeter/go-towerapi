package base

import (
	"fmt"
)

// An TowerError reports the error caused by an API request
type TowerError struct {
	// HTTP response that caused this error
	// Response *http.Response
	All    []string `json:"__all__,omitempty"`
	Detail string   `json:"detail,omitempty"`
}

func (e TowerError) Error() string {
	//return fmt.Sprintf("tower: %v %+v %v", e.All)
	return fmt.Sprintf("tower: %v %v", e.All, e.Detail)
}
