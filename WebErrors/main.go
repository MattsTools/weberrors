package WebErrors

import "net/http"

type WebError struct {
	ErrorIdentifier string
	ErrorMessage    string
	StatusCode      int // http status codes
}

func (c WebError) Error() string {
	return c.ErrorMessage
}

func (c WebError) MakeSafe() WebError {

	if c.StatusCode == http.StatusInternalServerError {
		return WebError{
			ErrorIdentifier: c.ErrorIdentifier,
			ErrorMessage:    http.StatusText(http.StatusInternalServerError),
			StatusCode:      c.StatusCode,
		}
	}

	return c
}

// This package should be used with caution
// it should not be used to create code that is super coupled together
// it should be used to allow a richer error return to the HTTP client according to REST standard
