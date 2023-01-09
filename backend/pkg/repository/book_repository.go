package repository

import (
	"bookstore/pkg/db"
	"bookstore/pkg/model"
	"context"
)

type BookRepository struct {
	conn *db.Conn
}

func NewBookRepository(conn *db.Conn) *BookRepository {
	return &BookRepository{conn: conn}
}

func (r *BookRepository) GetBookByID(ctx context.Context, id int) (*model.Book, error) {
	row := r.conn.QueryRowContext(ctx, "SELECT * FROM books WHERE id = $1", id)
	var book model.Book
	if err := row.Scan(&book.ID, &book.Title, &book.Author); err != nil {
		return nil, err
	}
	return &book, nil
}

func (r *BookRepository) CreateBook(ctx context.Context, book *model.Book) (int, error) {
	row := r.conn.QueryRowContext(ctx, "INSERT INTO books (title, author) VALUES ($1, $2) RETURNING id", book.Title, book.Author)
	var id int
	if err := row.Scan(&id); err != nil {
		return 0, err
	}
	return id, nil
}

func (r *BookRepository) UpdateBook(ctx context.Context, book *model.Book) error {
	_, err := r.conn.ExecContext(ctx, "UPDATE books SET title = $1, author = $2 WHERE id = $3", book.Title, book.Author, book.ID)
	return err
}

func (r *BookRepository) DeleteBook(ctx context.Context, id int) error {
	_, err := r.conn.ExecContext(ctx, "DELETE FROM books WHERE id = $1", id)
	return err
}