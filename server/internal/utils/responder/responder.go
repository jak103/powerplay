package responder

import (
	"encoding/json"
	"errors"
	"fmt"
	"path/filepath"
	"runtime"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/utils"
	"github.com/jak103/leaguemanager/internal/utils/locals"
	"github.com/jak103/leaguemanager/internal/utils/log"
)

var (
	errResponse = &ErrorResponse{}
)

type Response struct {
	StatusCode   int             `json:"status_code"`
	StatusString string          `json:"status_string"`
	RequestId    string          `json:"request_id"`
	Message      string          `json:"message,omitempty"`
	Paging       *Cursor         `json:"paging,omitempty"` // needs to be a pointer so omitempty will work
	Data         json.RawMessage `json:"data,omitempty"`
}

type ErrorResponse struct {
	Response
	err error `json:"-"`
}

func (r ErrorResponse) Error() string {
	return r.err.Error()
}

func WithErr(err error) ErrorResponse {
	return ErrorResponse{
		err: err,
	}
}

type Cursor struct {
	PrevPageToken *string `json:"prev_page_token"` // pointer here, but no omitempty: allows {"prev_page_token": null} behavior
	NextPageToken *string `json:"next_page_token"` // pointer here, but no omitempty: allows {"next_page_token": null} behavior
}

type StandardPageToken struct {
	ExpiresAt time.Time `json:"expires_at"`
}

func NewResponse(c *fiber.Ctx, code int, data ...any) *Response {
	var b json.RawMessage
	if len(data) > 0 {
		j, err := json.Marshal(data[0])
		if err != nil {
			log.WithErr(err).Error("Weird, this should never happen")
			return nil
		}
		b = j
	}

	return &Response{
		StatusCode:   code,
		StatusString: utils.StatusMessage(code),
		RequestId:    locals.RequestId(c),
		Data:         b,
	}
}

func NewErrorResponse(c *fiber.Ctx, code int, formatAndArgs ...any) *Response {
	return &Response{
		StatusCode:   code,
		StatusString: utils.StatusMessage(code),
		RequestId:    locals.RequestId(c),
		Message:      formatString(formatAndArgs...),
	}
}

func Ok(c *fiber.Ctx, data ...any) error {
	var d any
	if len(data) > 0 {
		d = data[0]
	}

	return respond(c, nil, fiber.StatusOK, d)
}

func OkMessage(c *fiber.Ctx, formatAndArgs ...any) error {
	if len(formatAndArgs) == 0 {
		return respond(c, nil, fiber.StatusOK, nil, "Success")
	}

	return respond(c, nil, fiber.StatusOK, nil, formatAndArgs...)
}

func Created(c *fiber.Ctx, data ...any) error {
	var d any
	if len(data) > 0 {
		d = data[0]
	}

	return respond(c, nil, fiber.StatusCreated, d)
}

func (r ErrorResponse) NotFound(c *fiber.Ctx, formatAndArgs ...any) error {
	if len(formatAndArgs) == 0 {
		return respond(c, r.err, fiber.StatusNotFound, nil, "Unable to find the requested resource")
	}

	return respond(c, r.err, fiber.StatusNotFound, nil, formatAndArgs...)
}

func NotFound(c *fiber.Ctx, formatAndArgs ...any) error {
	if len(formatAndArgs) == 0 {
		return respond(c, nil, fiber.StatusNotFound, nil, "Unable to find the requested resource")
	}

	return respond(c, nil, fiber.StatusNotFound, nil, formatAndArgs...)
}

func (r ErrorResponse) NotFoundRoute(c *fiber.Ctx, formatAndArgs ...any) error {
	if len(formatAndArgs) == 0 {
		return respond(c, r.err, fiber.StatusNotFound, nil, "Unable to find the requested route: %s", c.OriginalURL())
	}

	return respond(c, r.err, fiber.StatusNotFound, nil, formatAndArgs...)
}

func NotFoundRoute(c *fiber.Ctx) error {
	return respond(c, nil, fiber.StatusNotFound, nil, "Unable to find the requested route: %s", c.OriginalURL())
}

func (r ErrorResponse) BadRequest(c *fiber.Ctx, formatAndArgs ...any) error {
	return respond(c, r.err, fiber.StatusBadRequest, nil, formatAndArgs...)
}

func BadRequest(c *fiber.Ctx, formatAndArgs ...any) error {
	return respond(c, nil, fiber.StatusBadRequest, nil, formatAndArgs...)
}

func (r ErrorResponse) InternalServerError(c *fiber.Ctx, formatAndArgs ...any) error {
	if len(formatAndArgs) == 0 {
		return respond(c, r.err, fiber.StatusInternalServerError, nil, "An internal server error has occured, please submit a support ticket with the request ID if it persists")
	}

	return respond(c, r.err, fiber.StatusInternalServerError, nil, formatAndArgs...)
}

func InternalServerError(c *fiber.Ctx, formatAndArgs ...any) error {
	if len(formatAndArgs) == 0 {
		return respond(c, nil, fiber.StatusInternalServerError, nil, "An internal server error has occured, please submit a support ticket with the request ID if it persists")
	}

	return respond(c, nil, fiber.StatusInternalServerError, nil, formatAndArgs...)
}

func TooManyRequests(c *fiber.Ctx) error {
	return respond(c, nil, fiber.StatusTooManyRequests, nil, "too many requests for route %s", c.Path())
}

func (r ErrorResponse) TooManyRequests(c *fiber.Ctx) error {
	return respond(c, r.err, fiber.StatusTooManyRequests, nil, "too many requests for route %s", c.Path())
}

// The HTTP 401 Unauthorized response status code indicates that the client
// request has not been completed because it lacks valid authentication credentials for the requested resource.
// This status code is similar to the 403 Forbidden status code, except that in situations resulting
// in this status code, user authentication can allow access to the resource.
func Unauthorized(c *fiber.Ctx, formatAndArgs ...any) error {
	if len(formatAndArgs) == 0 {
		return respond(c, nil, fiber.StatusUnauthorized, nil, "Not authorized to access this resource")
	}

	return respond(c, nil, fiber.StatusUnauthorized, nil, formatAndArgs...)
}

func (r ErrorResponse) Unauthorized(c *fiber.Ctx, formatAndArgs ...any) error {
	if len(formatAndArgs) == 0 {
		return respond(c, r.err, fiber.StatusUnauthorized, nil, "Not authorized to access this resource")
	}

	return respond(c, r.err, fiber.StatusUnauthorized, nil, formatAndArgs...)
}

// The HTTP 403 Forbidden response status code indicates that the server understands the request but refuses to authorize it.
// This status is similar to 401, but for the 403 Forbidden status code, re-authenticating makes no
// difference. The access is tied to the application logic, such as insufficient scopes to a resource.
func Forbidden(c *fiber.Ctx, formatAndArgs ...any) error {
	if len(formatAndArgs) == 0 {
		return respond(c, nil, fiber.StatusForbidden, nil, "Access to this resource is forbidden")
	}

	return respond(c, nil, fiber.StatusForbidden, nil, formatAndArgs...)
}

func (r ErrorResponse) Forbidden(c *fiber.Ctx, formatAndArgs ...any) error {
	if len(formatAndArgs) == 0 {
		return respond(c, r.err, fiber.StatusForbidden, nil, "Access to this resource is forbidden")
	}

	return respond(c, r.err, fiber.StatusForbidden, nil, formatAndArgs...)
}

func UnsupportedMediaType(c *fiber.Ctx, formatAndArgs ...any) error {
	if len(formatAndArgs) == 0 {
		return respond(c, nil, fiber.StatusUnsupportedMediaType, nil, "Content type must be 'application/json'")
	}

	return respond(c, nil, fiber.StatusUnsupportedMediaType, nil, formatAndArgs...)
}

func ProxyAuthRequired(c *fiber.Ctx, formatAndArgs ...any) error {
	if len(formatAndArgs) == 0 {
		return respond(c, nil, fiber.StatusProxyAuthRequired, nil, "Authorization required")
	}

	return respond(c, nil, fiber.StatusProxyAuthRequired, nil, formatAndArgs...)
}

func Conflict(c *fiber.Ctx, formatAndArgs ...any) error {
	if len(formatAndArgs) == 0 {
		return respond(c, nil, fiber.StatusConflict, nil, "This resource is conflicted")
	}

	return respond(c, nil, fiber.StatusConflict, nil, formatAndArgs...)
}

func (r ErrorResponse) Conflict(c *fiber.Ctx, formatAndArgs ...any) error {
	if len(formatAndArgs) == 0 {
		return respond(c, r.err, fiber.StatusConflict, nil, "This resource is conflicted")
	}

	return respond(c, r.err, fiber.StatusConflict, nil, formatAndArgs...)
}

func (r ErrorResponse) Locked(c *fiber.Ctx, formatAndArgs ...any) error {
	if len(formatAndArgs) == 0 {
		return respond(c, r.err, fiber.StatusLocked, nil, "This resource is currently locked")
	}

	return respond(c, r.err, fiber.StatusLocked, nil, formatAndArgs...)
}

func Send(c *fiber.Ctx, r *Response) error {
	c.Response().Header.SetContentType("application/json")
	return c.Status(r.StatusCode).JSON(r)
}

func SendBytes(c *fiber.Ctx, code int, data []byte) error {
	return respondBytes(c, code, data)
}

func respondBytes(c *fiber.Ctx, code int, data []byte) error {
	// This function exists just to get the skip count correct for logInfo
	logInfo(c)
	c.Response().Header.SetContentType("application/json")
	return c.Status(code).Send(data)
}

func respond(c *fiber.Ctx, err error, code int, data any, formatAndArgs ...any) error {
	logInfo(c, formatAndArgs...)

	var b json.RawMessage
	if data != nil {
		j, err := json.Marshal(data)
		if err != nil {
			log.WithErr(err).Error("Weird, this should never happen")
			return InternalServerError(c)
		}
		b = j
	}

	if err == nil {
		r := Response{
			StatusCode:   code,
			StatusString: utils.StatusMessage(code),
			RequestId:    locals.RequestId(c),
			Message:      formatString(formatAndArgs...),
			Data:         b,
		}

		c.Response().Header.SetContentType("application/json")
		return c.Status(code).JSON(r)
	} else {
		er := ErrorResponse{
			Response: Response{
				StatusCode:   code,
				StatusString: utils.StatusMessage(code),
				RequestId:    locals.RequestId(c),
				Message:      formatString(formatAndArgs...),
				Data:         b,
			},
			err: err,
		}

		c.Response().Header.SetContentType("application/json")
		c.Status(code).JSON(er)
		return er
	}
}

func logInfo(c *fiber.Ctx, formatAndArgs ...any) {
	log := locals.Logger(c)

	if pc, _, _, ok := runtime.Caller(2); ok {
		name := runtime.FuncForPC(pc).Name()
		name = filepath.Base(name)

		line := fmt.Sprintf("Finished handler from %s()", name)
		if msg := formatString(formatAndArgs...); msg != "" {
			line = fmt.Sprintf("%s with responder.Message: '%s'", line, msg)
		}

		log.Info(line)
	}
}

func formatString(formatAndArgs ...any) string {
	if len(formatAndArgs) > 0 {
		if len(formatAndArgs) > 1 {
			return fmt.Sprintf(formatAndArgs[0].(string), formatAndArgs[1:]...)
		} else {
			return fmt.Sprint(formatAndArgs[0])
		}
	}

	return "" // nothing was passed, return a blank string. This will be the Responder.Message, which is omit empty
}

func GlobalErrorHandler(c *fiber.Ctx, err error) error {
	if errors.As(err, errResponse) {
		return nil
	}

	log := locals.Logger(c)
	log.WithErr(err).Error("Global error handler caught this error")

	// Return statuscode with error message
	return InternalServerError(c)
}
