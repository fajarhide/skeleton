package helper

import (
	"github.com/labstack/echo"
)

const (
	APP = "account"
)

type (
	// Response - structure response
	Response struct {
		Ctx          echo.Context `json:"-"`
		RequestID    string       `json:"request_id,omitempty"`
		Code         int          `json:"code"`
		Message      string       `json:"message"`
		ErrorMessage string       `json:"error_message,omitempty"`
		Data         interface{}  `json:"data,omitempty"`
		App          string       `json:"app"`
	}
)

// NewResponse - instantiate new Response
func NewResponse(code int, msg string, data interface{}) *Response {
	return &Response{
		Code:    code,
		Message: msg,
		Data:    data,
	}
}

// WriteResponse - write response to the client
func (r *Response) WriteResponse(ctx echo.Context) error {
	r.App = APP
	return ctx.JSON(r.Code, r)
}

