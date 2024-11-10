package controller

import (
	"github.com/devanfer02/go-blog/app/service"
	"github.com/devanfer02/go-blog/domain"
	"github.com/gin-gonic/gin"
)

type BlogController struct {
	blogSvc service.BlogService
}

func MountBlogRoutes(app *gin.Engine, blogSvc service.BlogService) {
	blogCtr := &BlogController{blogSvc: blogSvc}

	app.GET("/blogs", blogCtr.FetchAllBlogs)
}

func (c *BlogController) FetchAllBlogs(ctx *gin.Context) {
	var (
		code    int    = 500
		status  string = "fail"
		message string = "failed to fetch all blogs"
		blogs   []domain.Blog
		err     error = nil 
	)

	sendResp := func() {
		ctx.JSON(code, gin.H{
			"code":    code,
			"status":  status,
			"message": message,
			"data":    blogs,
		
		})
	}

	defer sendResp()

	blogs, err = c.blogSvc.GetAllBlogs()

	if err != nil {
		return
	}

	code = 200
	status = "success"
	message = "successfully fetch all blogs"
}
