package controller

import (
	"strconv"

	"github.com/devanfer02/go-blog/app/service"
	"github.com/devanfer02/go-blog/domain"
	"github.com/gin-gonic/gin"
)

type BlogController struct {
	blogSvc service.BlogService
}

func MountBlogRoutes(app *gin.Engine, blogSvc service.BlogService) {
	blogCtr := &BlogController{blogSvc: blogSvc}
	blogR := app.Group("/api/blogs")

	blogR.GET("", blogCtr.FetchAllBlogs)
	blogR.GET("/:id", blogCtr.FetchBlogByID)
	blogR.POST("", blogCtr.CreateBlog)
	blogR.PUT("/:id", blogCtr.UpdateBlog)
	blogR.DELETE("/:id", blogCtr.DeleteBlog)
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
			"err": func() string {
				if err != nil {
					return err.Error()
				}
				return ""
			}(),
		})
	}

	defer sendResp()

	blogs, err = c.blogSvc.GetAllBlogs()
	code = domain.GetCode(err)

	if err != nil {
		return
	}

	status = "success"
	message = "successfully fetch all blogs"
}

func (c *BlogController) FetchBlogByID(ctx *gin.Context) {
	var (
		code    int    = 500
		status  string = "fail"
		message string = "failed to fetch blog"
		blog    domain.Blog
		err     error = nil

		idParam = ctx.Param("id")
		id      int
	)

	sendResp := func() {
		ctx.JSON(code, gin.H{
			"code":    code,
			"status":  status,
			"message": message,
			"data":    blog,
			"err": func() string {
				if err != nil {
					return err.Error()
				}
				return ""
			}(),
		})
	}

	defer sendResp()
	id, err = strconv.Atoi(idParam)

	if err != nil {
		code = 400 
		return 
	}

	blog, err = c.blogSvc.GetBlogByID(id)
	code = domain.GetCode(err)

	if err != nil {
		return
	}

	status = "success"
	message = "successfully fetch blog"
}

func (c *BlogController) CreateBlog(ctx *gin.Context) {
	var (
		code    int    = 500
		status  string = "fail"
		message string = "failed to create blog"
		blog    domain.Blog
		err     error = nil
	)

	sendResp := func() {
		ctx.JSON(code, gin.H{
			"code":    code,
			"status":  status,
			"message": message,
			"err": func() string {
				if err != nil {
					return err.Error()
				}
				return ""
			}(),
		})
	}

	defer sendResp()
	
	if err := ctx.ShouldBindJSON(&blog); err != nil {
		code = 400
		return 
	}

	err = c.blogSvc.CreateBlog(&blog)
	code = domain.GetCode(err)

	if err != nil {
		return
	}

	status = "success"
	message = "successfully fetch blog"
}

func (c *BlogController) UpdateBlog(ctx *gin.Context) {
	var (
		code    int    = 500
		status  string = "fail"
		message string = "failed to update blog"
		blog    domain.Blog
		err     error = nil

		idParam = ctx.Param("id")
		id      int
	)

	sendResp := func() {
		ctx.JSON(code, gin.H{
			"code":    code,
			"status":  status,
			"message": message,
			"err": func() string {
				if err != nil {
					return err.Error()
				}
				return ""
			}(),
		})
	}

	defer sendResp()
	id, err = strconv.Atoi(idParam)

	if err != nil {
		code = 400 
		return 
	}

	if err := ctx.ShouldBindJSON(&blog); err != nil {
		code = 400
		return 
	}

	blog.ID = id 

	err = c.blogSvc.UpdateBlog(&blog)
	code = domain.GetCode(err)

	if err != nil {
		return
	}

	status = "success"
	message = "successfully update blog"
}

func (c *BlogController) DeleteBlog(ctx *gin.Context) {
	var (
		code    int    = 500
		status  string = "fail"
		message string = "failed to delete blog"
		blog    domain.Blog
		err     error = nil

		idParam = ctx.Param("id")
		id      int
	)

	sendResp := func() {
		ctx.JSON(code, gin.H{
			"code":    code,
			"status":  status,
			"message": message,
			"err": func() string {
				if err != nil {
					return err.Error()
				}
				return ""
			}(),
		})
	}

	defer sendResp()
	id, err = strconv.Atoi(idParam)

	if err != nil {
		code = 400 
		return 
	}

	if err := ctx.ShouldBindJSON(&blog); err != nil {
		code = 400
		return 
	}

	blog.ID = id 

	err = c.blogSvc.DeleteBlog(&blog)
	code = domain.GetCode(err)

	if err != nil {
		return
	}

	status = "success"
	message = "successfully delete blog"
}