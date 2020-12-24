package client

import "fmt"

type ResponseError struct {
	Err error
	StatusCode int
}

func (r *ResponseError) Error() string {
	return fmt.Sprintf("status %d: err %v", r.StatusCode, r.Err)
}
