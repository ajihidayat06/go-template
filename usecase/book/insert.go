package book

import (
	"context"
	"go-template/errutils"
	"go-template/model"
)

func (u *bookUseCaseImpl) InsertBook(ctx context.Context, book model.BookRequest) (response interface{}, err errutils.ErrorModel) {
	// response, err = InsertWithTx(u.DB, u.doInsert, book)
	// if err != nil {
	// 	return
	// }

	if book.Id == 0 {
		return nil, errutils.GenerateErrBadRequest("book id is null")
	}

	return book, errutils.NilErr()
}
