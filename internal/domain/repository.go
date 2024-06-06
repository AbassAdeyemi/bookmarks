package domain

import (
	"context"
	"github.com/AbassAdeyemi/bookmarks/internal/config"
	"github.com/jackc/pgx/v5"
	"time"
)

type Bookmark struct {
	ID        int       `json:"id"`
	Title     string    `json:"title"`
	Url       string    `json:"url"`
	CreatedAt time.Time `json:"created_at"`
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
	query := "SELECT id, title, url, created_at FROM bookmarks WHERE id=$1"
	var bookmark = Bookmark{}
	row := r.db.QueryRow(ctx, query, id)
	err := row.Scan(&bookmark.ID, &bookmark.Title, &bookmark.Url, &bookmark.CreatedAt)
	if err != nil {
		r.logger.Errorf("Error fetching bookmark with id: %v", id)
		return nil, err
	}

	return &bookmark, nil
}

func (r bookmarkRepo) Create(ctx context.Context, b Bookmark) (*Bookmark, error) {
	query := "insert into bookmarks(title, url, created_at) values ($1, $2, $3) RETURNING id"
	var lastInsertID int
	err := r.db.QueryRow(ctx, query, b.Title, b.Url, b.CreatedAt).Scan(&lastInsertID)
	if err != nil {
		r.logger.Errorf("Error while inserting bookmark: %v", err)
		return nil, err
	}
	b.ID = lastInsertID
	return &b, nil
}

func (r bookmarkRepo) Update(ctx context.Context, b Bookmark) error {
	existingBookmark, err := r.GetByID(ctx, b.ID)
	if err != nil {
		return err
	}

	if len(b.Title) != 0 {
		existingBookmark.Title = b.Title
	}
	if len(b.Url) != 0 {
		existingBookmark.Url = b.Url
	}
	query := "update bookmarks set title=$1, url=$2 where id=$3"

	_, err = r.db.Exec(ctx, query, b.Title, b.Url, b.ID)

	if err != nil {
		r.logger.Errorf("Error updating bookmark with id: %v", b.ID)
		return err
	}

	return nil
}

func (r bookmarkRepo) Delete(ctx context.Context, id int) error {
	query := "delete from bookmarks where id=$1"

	_, err := r.db.Exec(ctx, query, id)
	if err != nil {
		r.logger.Errorf("Error deleting bookmark with id: %v", id)
		return err
	}
	return nil
}
