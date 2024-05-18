package domain

import (
	"context"
	"github.com/AbassAdeyemi/bookmarks/internal/config"
	"github.com/jackc/pgx/v5"
	"time"
)

type Bookmark struct {
	ID        string
	Title     string
	Url       string
	CreatedAt time.Time
}

type BookmarkRepository interface {
	GetAll(ctx context.Context) ([]Bookmark, error)
	GetByID(ctx context.Context, id int) (*Bookmark, error)
	Create(ctx context.Context, b Bookmark) (*Bookmark, error)
	Update(ctx context.Context, b Bookmark) error
	Delete(ctx context.Context, id int) error
}

type bookmarkRepo struct {
	db     *pgx.Conn
	logger *config.Logger
}

func NewBookmarkRepository(db *pgx.Conn, logger *config.Logger) BookmarkRepository {
	return bookmarkRepo{
		db:     db,
		logger: logger,
	}
}

func (r bookmarkRepo) GetAll(ctx context.Context) ([]Bookmark, error) {
	query := "SELECT id, title, url, created_at FROM bookmarks"
	rows, err := r.db.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var bookmarks []Bookmark
	for rows.Next() {
		var bookmark = Bookmark{}
		err = rows.Scan(&bookmark.ID, &bookmark.Title, &bookmark.Url, &bookmark.CreatedAt)
		if err != nil {
			return nil, err
		}
		bookmarks = append(bookmarks, bookmark)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return bookmarks, nil
}

func (r bookmarkRepo) GetByID(ctx context.Context, id int) (*Bookmark, error) {
	panic("implement me")
}

func (r bookmarkRepo) Create(ctx context.Context, b Bookmark) (*Bookmark, error) {
	panic("implement me")
}

func (r bookmarkRepo) Update(ctx context.Context, b Bookmark) error {
	panic("implement me")
}

func (r bookmarkRepo) Delete(ctx context.Context, id int) error {
	panic("implement me")
}
