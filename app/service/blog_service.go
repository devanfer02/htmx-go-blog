package service

import (
	"github.com/devanfer02/go-blog/app/repository"
	"github.com/devanfer02/go-blog/domain"
)

type BlogService interface {
	GetAllBlogs() ([]domain.Blog, error)
	GetBlogByID(id int) (domain.Blog, error)
	CreateBlog(blog *domain.Blog) error 
	UpdateBlog(blog *domain.Blog) error 
	DeleteBlog(blog *domain.Blog) error 
}

type blogService struct {
	blogRepo repository.BlogRepository
}

func NewBlogService(blogRepo repository.BlogRepository) BlogService {
	return &blogService{blogRepo: blogRepo}
}

func(s *blogService) GetAllBlogs() ([]domain.Blog, error) {
	blogs, err := s.blogRepo.FetchAllBlogs()

	if err != nil {
		return nil, err 
	}

	return blogs, err 
}

func(s *blogService) GetBlogByID(id int) (domain.Blog, error) {
	return domain.Blog{}, nil 
}

func(s *blogService) CreateBlog(blog *domain.Blog) error  {
	return nil 
}

func(s *blogService) UpdateBlog(blog *domain.Blog) error  {
	return nil 
}

func(s *blogService) DeleteBlog(blog *domain.Blog) error  {
	return nil 
}
