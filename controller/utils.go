package controller

import (
	"go-template/errutils"
	"go-template/model"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func ReadRequest(req interface{}, ctx *fiber.Ctx) (id int64, err error) {

	if ctx.Method() != "GET" {
		err = ctx.BodyParser(req)
		if err != nil {
			return
		}
	}

	if ctx.Params("id") != "" {
		idInt, errs := ReadRequestParamsID(ctx)
		if errs != nil {
			err = errs
			return
		}
		id = int64(idInt)
	}

	return
}

func ReadRequestParamsID(ctx *fiber.Ctx) (id int64, err error) {
	idInt, errs := strconv.Atoi(ctx.Params("id"))
	if errs != nil {
		err = errs
		return
	}
	id = int64(idInt)

	return
}

func ResponseErr(ctx *fiber.Ctx, err errutils.ErrorModel) error {
	r := model.Response{
		StatusCode: err.StatusCode,
		Message:    err.ErrMessage,
		ErrDetail:  err.ErrDetail,
	}

	ctx.Locals("err_file", err.ErrFile)
	return r.SetResponseJson(ctx)
}

func ResponseSucces(ctx *fiber.Ctx, data any) error {
	r := model.Response{
		StatusCode: fiber.StatusOK,
		Message:    "success",
		Data:       data,
	}

	return r.SetResponseJson(ctx)
}
