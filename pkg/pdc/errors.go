package pdc

import (
	"errors"
	"fmt"
	"net/http"
)

var (
	ErrLoginFailed = errors.New("error: login failed, JSESSIONID cookie not found in response")
)

func generateError(res *http.Response) error {
	defer res.Body.Close()

	body, _ := decodeBody(res.Body)

	return fmt.Errorf("error: HTTP %v %v: %v", res.StatusCode, http.StatusText(res.StatusCode), body)
}
