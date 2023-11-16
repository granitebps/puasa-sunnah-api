package utils

import (
	"time"

	"github.com/ansel1/merry/v2"
	"github.com/gofiber/fiber/v2"
	"github.com/newrelic/go-agent/v3/integrations/nrpkgerrors"
	"github.com/newrelic/go-agent/v3/newrelic"
)

type JSONResponse struct {
	Success   bool        `json:"success"`
	Message   string      `json:"message"`
	Data      interface{} `json:"data"`
	Timestamp int64       `json:"timestamp"`
}

func ReturnSuccessResponse(c *fiber.Ctx, code int, msg string, data interface{}) error {
	c.Set(fiber.HeaderContentType, fiber.MIMEApplicationJSON)
	return c.Status(code).JSON(JSONResponse{
		Success:   true,
		Message:   msg,
		Data:      data,
		Timestamp: time.Now().UnixMilli(),
	})
}

func ReturnErrorResponse(c *fiber.Ctx, err error, data interface{}) error {
	txn := newrelic.FromContext(c.UserContext())
	txn.NoticeError(nrpkgerrors.Wrap(err))

	msg := merry.UserMessage(err)
	if msg == "" {
		msg = err.Error()
	}

	c.Set(fiber.HeaderContentType, fiber.MIMEApplicationJSON)
	return c.Status(merry.HTTPCode(err)).JSON(JSONResponse{
		Success:   false,
		Message:   msg,
		Data:      data,
		Timestamp: time.Now().UnixMilli(),
	})
}
