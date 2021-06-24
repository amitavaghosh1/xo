package oracle

// Code generated by xo. DO NOT EDIT.

import (
	"context"
	"database/sql"
	"time"
)

// Book represents a row from 'booktest.books'.
type Book struct {
	BookID    int       `json:"book_id"`   // book_id
	AuthorID  int       `json:"author_id"` // author_id
	Isbn      string    `json:"isbn"`      // isbn
	Title     string    `json:"title"`     // title
	Year      int       `json:"year"`      // year
	Available time.Time `json:"available"` // available
	Tags      string    `json:"tags"`      // tags
	// xo fields
	_exists, _deleted bool
}

// Exists returns true when the Book exists in the database.
func (b *Book) Exists() bool {
	return b._exists
}

// Deleted returns true when the Book has been marked for deletion from
// the database.
func (b *Book) Deleted() bool {
	return b._deleted
}

// Insert inserts the Book to the database.
func (b *Book) Insert(ctx context.Context, db DB) error {
	switch {
	case b._exists: // already exists
		return logerror(&ErrInsertFailed{ErrAlreadyExists})
	case b._deleted: // deleted
		return logerror(&ErrInsertFailed{ErrMarkedForDeletion})
	}
	// insert (primary key generated and returned by database)
	const sqlstr = `INSERT INTO booktest.books (` +
		`author_id, isbn, title, year, available, tags` +
		`) VALUES (` +
		`:1, :2, :3, :4, :5, :6` +
		`) RETURNING book_id /*LASTINSERTID*/ INTO :pk`
	// run
	logf(sqlstr, b.AuthorID, b.Isbn, b.Title, b.Year, b.Available, b.Tags)
	var id int64
	if _, err := db.ExecContext(ctx, sqlstr, b.AuthorID, b.Isbn, b.Title, b.Year, b.Available, b.Tags, sql.Named("pk", sql.Out{Dest: &id})); err != nil {
		return err
	} // set primary key
	b.BookID = int(id)
	// set exists
	b._exists = true
	return nil
}

// Update updates a Book in the database.
func (b *Book) Update(ctx context.Context, db DB) error {
	switch {
	case !b._exists: // doesn't exist
		return logerror(&ErrUpdateFailed{ErrDoesNotExist})
	case b._deleted: // deleted
		return logerror(&ErrUpdateFailed{ErrMarkedForDeletion})
	}
	// update with primary key
	const sqlstr = `UPDATE booktest.books SET ` +
		`author_id = :1, isbn = :2, title = :3, year = :4, available = :5, tags = :6` +
		` WHERE book_id = :7`
	// run
	logf(sqlstr, b.AuthorID, b.Isbn, b.Title, b.Year, b.Available, b.Tags, b.BookID)
	if _, err := db.ExecContext(ctx, sqlstr, b.AuthorID, b.Isbn, b.Title, b.Year, b.Available, b.Tags, b.BookID); err != nil {
		return logerror(err)
	}
	return nil
}

// Save saves the Book to the database.
func (b *Book) Save(ctx context.Context, db DB) error {
	if b.Exists() {
		return b.Update(ctx, db)
	}
	return b.Insert(ctx, db)
}

// Delete deletes the Book from the database.
func (b *Book) Delete(ctx context.Context, db DB) error {
	switch {
	case !b._exists: // doesn't exist
		return nil
	case b._deleted: // deleted
		return nil
	}
	// delete with single primary key
	const sqlstr = `DELETE FROM booktest.books WHERE book_id = :1`
	// run
	logf(sqlstr, b.BookID)
	if _, err := db.ExecContext(ctx, sqlstr, b.BookID); err != nil {
		return logerror(err)
	}
	// set deleted
	b._deleted = true
	return nil
}

// BookByIsbn retrieves a row from 'booktest.books' as a Book.
//
// Generated from index 'books_isbn_key'.
func BookByIsbn(ctx context.Context, db DB, isbn string) (*Book, error) {
	// query
	const sqlstr = `SELECT ` +
		`book_id, author_id, isbn, title, year, available, tags ` +
		`FROM booktest.books ` +
		`WHERE isbn = :1`
	// run
	logf(sqlstr, isbn)
	b := Book{
		_exists: true,
	}
	if err := db.QueryRowContext(ctx, sqlstr, isbn).Scan(&b.BookID, &b.AuthorID, &b.Isbn, &b.Title, &b.Year, &b.Available, &b.Tags); err != nil {
		return nil, logerror(err)
	}
	return &b, nil
}

// BookByBookID retrieves a row from 'booktest.books' as a Book.
//
// Generated from index 'books_pkey'.
func BookByBookID(ctx context.Context, db DB, bookID int) (*Book, error) {
	// query
	const sqlstr = `SELECT ` +
		`book_id, author_id, isbn, title, year, available, tags ` +
		`FROM booktest.books ` +
		`WHERE book_id = :1`
	// run
	logf(sqlstr, bookID)
	b := Book{
		_exists: true,
	}
	if err := db.QueryRowContext(ctx, sqlstr, bookID).Scan(&b.BookID, &b.AuthorID, &b.Isbn, &b.Title, &b.Year, &b.Available, &b.Tags); err != nil {
		return nil, logerror(err)
	}
	return &b, nil
}

// BooksByTitleYear retrieves a row from 'booktest.books' as a Book.
//
// Generated from index 'books_title_idx'.
func BooksByTitleYear(ctx context.Context, db DB, title string, year int) ([]*Book, error) {
	// query
	const sqlstr = `SELECT ` +
		`book_id, author_id, isbn, title, year, available, tags ` +
		`FROM booktest.books ` +
		`WHERE title = :1 AND year = :2`
	// run
	logf(sqlstr, title, year)
	rows, err := db.QueryContext(ctx, sqlstr, title, year)
	if err != nil {
		return nil, logerror(err)
	}
	defer rows.Close()
	// process
	var res []*Book
	for rows.Next() {
		b := Book{
			_exists: true,
		}
		// scan
		if err := rows.Scan(&b.BookID, &b.AuthorID, &b.Isbn, &b.Title, &b.Year, &b.Available, &b.Tags); err != nil {
			return nil, logerror(err)
		}
		res = append(res, &b)
	}
	if err := rows.Err(); err != nil {
		return nil, logerror(err)
	}
	return res, nil
}

// Author returns the Author associated with the Book's AuthorID (author_id).
//
// Generated from foreign key 'books_author_id_fkey'.
func (b *Book) Author(ctx context.Context, db DB) (*Author, error) {
	return AuthorByAuthorID(ctx, db, b.AuthorID)
}
