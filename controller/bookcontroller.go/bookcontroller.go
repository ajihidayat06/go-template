package bookcontroller

import (
	"go-template/config"
	"go-template/controller"
	"go-template/errutils"
	"go-template/model"
	usecase "go-template/usecase/book"

	"github.com/gofiber/fiber/v2"
)

type BookController struct {
	cfg         *config.CfgStruct
	BookUseCase usecase.BookUseCase
}

func NewBookController(bookUseCase *usecase.BookUseCase, cfg *config.CfgStruct) BookController {
	return BookController{
		cfg:         cfg,
		BookUseCase: *bookUseCase,
	}
}

func (r *BookController) InsertBook(ctx *fiber.Ctx) error {
	var BookRequest model.BookRequest

	_, err := controller.ReadRequest(&BookRequest, ctx)
	if err != nil {
		return controller.ResponseErr(ctx, errutils.GenerateErrUnknown(err))
	}

	// validation
	err = r.cfg.Validatior.Struct(BookRequest)
	if err != nil {
		return controller.ResponseErr(ctx, errutils.GenerateErrBadRequest(err.Error()))
	}

	// save data
	response, errMdl := r.BookUseCase.InsertBook(ctx.Context(), model.BookRequest{
		Id:    BookRequest.Id,
		Title: BookRequest.Title,
	})
	if errMdl.Error() != nil {
		return controller.ResponseErr(ctx, errMdl)
	}

	return controller.ResponseSucces(ctx, response)
}
