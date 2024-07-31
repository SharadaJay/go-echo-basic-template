package intf

import (
	"basic-echo-app/BookStore/model"
	"context"
)

type BookService interface {
	PrintBookTitle(ctx context.Context, book *model.Book)
	TestBookService(ctx context.Context)
}

type BookDatalayer interface {
	GetBookByID(ctx context.Context, bookID int16)
	TestBookDatalayer(ctx context.Context)
}
