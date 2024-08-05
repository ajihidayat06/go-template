package errutils

import (
	"errors"

	"github.com/gofiber/fiber/v2"
)

type ErrorModel struct {
	Err        error
	ErrCode    string
	ErrMessage string
	ErrDetail  string
	StatusCode int
}

func (e *ErrorModel) Error() error {
	return e.Err
}

func GenerateErr(err ErrorModel) ErrorModel {
	return ErrorModel{
		Err:        err.Err,
		ErrCode:    err.ErrCode,
		ErrMessage: err.ErrMessage,
		StatusCode: err.StatusCode,
		ErrDetail:  err.ErrDetail,
	}
}

func NilErr() ErrorModel {
	return ErrorModel{}
}

func GenerateCustomErr(err error, msg string) ErrorModel {
	return GenerateErr(ErrorModel{
		Err:        err,
		ErrMessage: msg,
		StatusCode: fiber.ErrInternalServerError.Code,
		ErrDetail:  err.Error(),
	})
}

func GenerateErrBadRequest(msg string) ErrorModel {
	return GenerateErr(ErrorModel{
		Err:        errors.New(fiber.ErrBadRequest.Message),
		ErrMessage: msg,
		StatusCode: fiber.StatusBadRequest,
		ErrDetail:  msg,
	})
}

func GenerateErrInternalServerError(err error, msg string) ErrorModel {
	return GenerateErr(ErrorModel{
		Err:        err,
		ErrMessage: msg,
		StatusCode: fiber.StatusInternalServerError,
		ErrDetail:  err.Error(),
	})
}

func GenerateErrUnknown(err error) ErrorModel {
	return GenerateErr(ErrorModel{
		Err:        err,
		ErrMessage: "unkown error, please contact customer service",
		StatusCode: fiber.StatusBadRequest,
		ErrDetail:  err.Error(),
	})
}
