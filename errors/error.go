package errors

import (
	"errors"

	"github.com/ansel1/merry"
	"github.com/gofiber/fiber/v2"
)

var (
	ErrInternalServerError = errors.New("internal server error")
	ErrEndpointNotFound    = errors.New("endpoint not found")
)

// Wrap error with user message and HTTP code.
func WrapUserMessageAndCode(err error, msg string, code int) error {
	if code == 0 {
		code = fiber.StatusInternalServerError
	}
	if msg == "" {
		msg = err.Error()
	}
	return merry.Wrap(err).WithUserMessage(msg).WithHTTPCode(code)
}
