package headerflag

import (
	"errors"
	"net/http"
	"strings"
)

type headerFlag http.Header

func (h *headerFlag) String() string {
	return "headers"
}

func (h *headerFlag) Set(value string) error {
	k, v, ok := strings.Cut(value, "=")
	if !ok {
		return errors.New("header missing =")
	}
	http.Header(*h).Add(k, v)
	return nil
}

func (h *headerFlag) Headers() http.Header {
	return http.Header(*h)
}

func New() *headerFlag {
	return &headerFlag{}
}
