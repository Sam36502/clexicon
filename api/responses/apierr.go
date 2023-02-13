package responses

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

type APIError struct {
	httpCode int
	Code     string `json:"code"`
	Msg      string `json:"msg"`
}

// All Allowed error codes
const (
	APIERR_OBJECT_EXISTS  = "object_exists"
	APIERR_INVALID_OBJECT = "invalid_object"
	APIERR_INVALID_PARAM  = "invalid_param"
	APIERR_INVALID_ERROR  = "invalid_error"
	APIERR_MODEL_FAIL     = "model_fail"
)

var g_ErrorTemplates = map[string]APIError{
	APIERR_OBJECT_EXISTS: {
		httpCode: http.StatusConflict,
		Code:     APIERR_OBJECT_EXISTS,
		Msg:      "Object with %s '%s' already exists.",
	},
	APIERR_INVALID_OBJECT: {
		httpCode: http.StatusBadRequest,
		Code:     APIERR_INVALID_OBJECT,
		Msg:      "Invalid object parameter provided: %v",
	},
	APIERR_INVALID_PARAM: {
		httpCode: http.StatusBadRequest,
		Code:     APIERR_INVALID_PARAM,
		Msg:      "Invalid parameter '%s' provided: %s",
	},
	APIERR_INVALID_ERROR: {
		httpCode: http.StatusInternalServerError,
		Code:     APIERR_INVALID_ERROR,
		Msg:      "Invalid error code '%s' used internally; please contact the developers if the problem persists.",
	},
	APIERR_MODEL_FAIL: {
		httpCode: http.StatusInternalServerError,
		Code:     APIERR_MODEL_FAIL,
		Msg:      "Failed to make request to data-model: %v",
	},
}

func GetAPIError(ctx echo.Context, code string, msgVars ...interface{}) error {
	if errTmp, ex := g_ErrorTemplates[code]; ex {
		return ctx.JSON(errTmp.httpCode, APIError{
			Code: errTmp.Code,
			Msg:  fmt.Sprintf(errTmp.Msg, msgVars...),
		})
	}

	return GetAPIError(ctx, APIERR_INVALID_ERROR, code)
}
