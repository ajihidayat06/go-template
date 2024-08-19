package errutils

import (
	"errors"
	"fmt"
	"go-template/common"

	"github.com/gofiber/fiber/v2"
)

type ErrorModel struct {
	Err        error
	ErrFile    string
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
		ErrFile:    common.GetRuntimeCaller(3),
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
		ErrFile:    common.GetRuntimeCaller(3),
		ErrMessage: msg,
		StatusCode: fiber.ErrInternalServerError.Code,
		ErrDetail:  err.Error(),
	})
}

func GenerateErrBadRequest(msg string) ErrorModel {
	return GenerateErr(ErrorModel{
		Err:        errors.New(fiber.ErrBadRequest.Message),
		ErrFile:    common.GetRuntimeCaller(3),
		ErrMessage: msg,
		StatusCode: fiber.StatusBadRequest,
		ErrDetail:  msg,
	})
}

func GenerateErrInternalServerError(err error, msg string) ErrorModel {
	return GenerateErr(ErrorModel{
		Err:        err,
		ErrFile:    common.GetRuntimeCaller(3),
		ErrMessage: msg,
		StatusCode: fiber.StatusInternalServerError,
		ErrDetail:  err.Error(),
	})
}

func GenerateErrUnknown(err error) ErrorModel {
	return GenerateErr(ErrorModel{
		Err:        err,
		ErrFile:    common.GetRuntimeCaller(3),
		ErrMessage: "unkown error, please contact customer service",
		StatusCode: fiber.StatusBadRequest,
		ErrDetail:  err.Error(),
	})
}

func GenerateErrInvalidFormatField(err error, fieldName string) ErrorModel {
	return GenerateErr(ErrorModel{
		Err:        err,
		ErrFile:    common.GetRuntimeCaller(3),
		ErrMessage: fmt.Sprintf("invalid format field %s", fieldName),
		StatusCode: fiber.StatusBadRequest,
		ErrDetail:  err.Error(),
	})
}

func GenerateErrInvalidRequest(err error) ErrorModel {
	return GenerateErr(ErrorModel{
		Err:        err,
		ErrFile:    common.GetRuntimeCaller(3),
		ErrMessage: "invalid request",
		StatusCode: fiber.StatusBadRequest,
		ErrDetail:  err.Error(),
	})
}
