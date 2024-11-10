package repository

import (
	"log"

	sq "github.com/Masterminds/squirrel"
	"github.com/jmoiron/sqlx"

	"github.com/devanfer02/go-blog/domain"
)

type BlogRepository interface {
	FetchAllBlogs() ([]domain.Blog, error)
	FetchBlogByID(id int) (domain.Blog, error)
	InsertBlog(blog *domain.Blog) error
	UpdateBlog(blog *domain.Blog) error
	DeleteBlog(blog *domain.Blog) error
}

type pgsqlBlogRepository struct {
	conn *sqlx.DB
}

func NewPgsqlBlogRepository(conn *sqlx.DB) BlogRepository {
	return &pgsqlBlogRepository{conn}
}

func (r *pgsqlBlogRepository) FetchAllBlogs() ([]domain.Blog, error) {
	var (
		query sq.SelectBuilder
		sql   string
		err   error
		blogs []domain.Blog = make([]domain.Blog, 0)
	)

	query = sq.Select("*").From("blogs")

	sql, _, err = query.ToSql()

	if err != nil {
		log.Printf("[BLOG REPOSITORY][FetchAllBlogs] ERR: %v\n", err.Error())
		return nil, err 
	}

	err = r.conn.Select(&blogs, sql)

	if err != nil {
		log.Printf("[BLOG REPOSITORY][FetchAllBlogs] ERR: %v\n", err.Error())
		return nil, err 
	}

	return blogs, nil 
}

func (r *pgsqlBlogRepository) FetchBlogByID(id int) (domain.Blog, error) {
	return domain.Blog{}, nil
}

func (r *pgsqlBlogRepository) InsertBlog(blog *domain.Blog) error {
	return nil
}

func (r *pgsqlBlogRepository) UpdateBlog(blog *domain.Blog) error {
	return nil
}

func (r *pgsqlBlogRepository) DeleteBlog(blog *domain.Blog) error {
	return nil
}
