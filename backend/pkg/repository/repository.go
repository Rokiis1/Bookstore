package repository

import (
	"bookstore/pkg/model"
	"context"
)

type Repository interface {
	GetBookByID(ctx context.Context, id int) (*model.Book, error)
	CreateBook(ctx context.Context, book *model.Book) (int, error)
	UpdateBook(ctx context.Context, book *model.Book) error
	DeleteBook(ctx context.Context, id int) error
	// Other resource repository functions go here...
}