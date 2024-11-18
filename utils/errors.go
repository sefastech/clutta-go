package utils

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog"
)

func logError(c echo.Context, err error, logger zerolog.Logger) {
	var (
		method = c.Request().Method
		uri    = c.Request().RequestURI
	)

	logger.Error().Err(err).Str("method", method).Str("uri", uri)
}

func errorResponse(c echo.Context, status int, message any, logger zerolog.Logger) {
	env := Envelope{"error": message, "statusCode": status}

	err := WriteJSON(c, status, env, nil)
	if err != nil {
		logError(c, err, logger)
		c.Response().WriteHeader(500)
	}
}

func ServerErrorResponse(c echo.Context, err error, logger zerolog.Logger) {
	logError(c, err, logger)
	message := "The server encountered a problem and could not process your request"
	errorResponse(c, http.StatusInternalServerError, message, logger)
}

func NotFoundResponse(c echo.Context, logger zerolog.Logger) {
	message := "The requested resource could not be found"
	errorResponse(c, http.StatusNotFound, message, logger)
}

func MethodNotAllowedResponse(c echo.Context, logger zerolog.Logger) {
	message := fmt.Sprintf("The %s method is not supported for this resource", c.Request().Method)
	errorResponse(c, http.StatusMethodNotAllowed, message, logger)
}

func BadRequestResponse(c echo.Context, err error, logger zerolog.Logger) {
	errorResponse(c, http.StatusBadRequest, err.Error(), logger)
}

func FailedValidationResponse(c echo.Context, errors map[string]string, logger zerolog.Logger) {
	errorResponse(c, http.StatusUnprocessableEntity, errors, logger)
}

func EditConflictResponse(c echo.Context, logger zerolog.Logger) {
	message := "Unable to update the record due to an edit conflict, please try again"
	errorResponse(c, http.StatusConflict, message, logger)
}

func RateLimitExceededResponse(c echo.Context, logger zerolog.Logger) {
	message := "Rate limit exceeded"
	errorResponse(c, http.StatusTooManyRequests, message, logger)
}

func InvalidAuthenticationTokenResponse(c echo.Context, logger zerolog.Logger) {
	c.Response().Header().Set("WWW-Authenticate", "Bearer")
	message := "Unauthorized"
	errorResponse(c, http.StatusUnauthorized, message, logger)
}
