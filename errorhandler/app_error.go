package errorhandler

import (
	"errors"
	"fmt"
	"net/http"
	"strings"
)

type AppError struct {
	StatusCode int    `json:"status_code"`
	RootError  error  `json:"-"`
	Message    string `json:"message"`
	Log        string `json:"-"`
	ErrorKey   string `json:"error_key"`
}

func NewCustomError(statusCode int, root error, msg, key string) *AppError {
	if root != nil {
		return &AppError{
			StatusCode: statusCode,
			RootError:  root,
			Message:    msg,
			Log:        root.Error(),
			ErrorKey:   key,
		}
	}

	rootErr := errors.New(msg)
	return &AppError{
		StatusCode: statusCode,
		RootError:  rootErr,
		Message:    msg,
		Log:        rootErr.Error(),
		ErrorKey:   key,
	}
}

func (ae *AppError) GetRootError() error {
	if err, ok := ae.RootError.(*AppError); ok {
		return err.GetRootError()
	}

	return ae.RootError
}

func (ae *AppError) Error() string {
	return ae.GetRootError().Error()
}

func ErrCannotGetRecord(recordName string, err error) *AppError {
	return NewCustomError(
		http.StatusNotFound,
		err,
		fmt.Sprintf("Cannot get %s", strings.ToLower(recordName)),
		fmt.Sprintf("err_cannot_get_%s", strings.ToLower(recordName)),
	)
}

func ErrBadRequest(err error) *AppError {
	return NewCustomError(
		http.StatusBadRequest,
		err,
		"Invalid request",
		"err_bad_request",
	)
}

func ErrInternal(err error) *AppError {
	return NewCustomError(
		http.StatusInternalServerError,
		err,
		"Something went wrong",
		"err_internal",
	)
}

func ErrUnauthorized(err error) *AppError {
	return NewCustomError(
		http.StatusUnauthorized,
		err,
		"Unauthorized",
		"err_unauthorized",
	)
}

func ErrRecordExisted(recordName string, err error) *AppError {
	return NewCustomError(
		http.StatusBadRequest,
		err,
		fmt.Sprintf("%s existed", strings.ToLower(recordName)),
		fmt.Sprintf("err_%s_existed", strings.ToLower(recordName)),
	)
}

func ErrInvalidRecord(recordName string, err error) *AppError {
	return NewCustomError(
		http.StatusBadRequest,
		err,
		fmt.Sprintf("invalid %s", strings.ToLower(recordName)),
		fmt.Sprintf("err_invalid_%s", strings.ToLower(recordName)),
	)
}
