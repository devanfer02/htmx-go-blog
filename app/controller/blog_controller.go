package controller

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/devanfer02/go-blog/app/service"
	"github.com/devanfer02/go-blog/domain"
	"github.com/devanfer02/go-blog/pkg/constants"
	// "github.com/devanfer02/go-blog/pkg/helpers"
	"github.com/gin-gonic/gin"
)

type BlogController struct {
	blogSvc service.BlogService
}

// todo: change some of the api routes to give html response

func MountBlogRoutes(app *gin.Engine, blogSvc service.BlogService) {
	blogCtr := &BlogController{blogSvc: blogSvc}

	app.GET("/", blogCtr.Index)
	app.GET("/blogs", blogCtr.ListBlogs)
	app.GET("/blogs/:id", blogCtr.ShowBlog)
	app.GET("/blogs/create", blogCtr.BlogForm)
	app.GET("/blogs/edit/:id", blogCtr.EditBlog)
	app.POST("/blogs", blogCtr.CreateBlog)
	app.PUT("/blogs/:id", blogCtr.UpdateBlog)

	app.DELETE("/blogs/:id", blogCtr.DeleteBlog)
}

func (c *BlogController) Index(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "Base", gin.H{
		"Title":   "HTMX Go Blog",
		"Content": "Home",
		"Navs":    constants.Navs,
	})
}

func (c *BlogController) BlogForm(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "Base", gin.H{
		"Title":   "HTMX Go Blog",
		"Content": "CreateBlog",
		"Navs":    constants.Navs,
	})
}

func (c *BlogController) EditBlog(ctx *gin.Context) {
	var (
		message string = "failed to fetch blog"
		blog    domain.Blog
		err     error = nil

		idParam = ctx.Param("id")
		id      int
	)

	sendResp := func() {
		ctx.HTML(http.StatusOK, "Base", gin.H{
			"Title":   "HTMX Go Blog",
			"Content": "EditBlog",
			"Navs":    constants.Navs,
			"Err":     err,
			"Blog":    blog,
			"Message": message,
		})
	}

	defer sendResp()
	id, err = strconv.Atoi(idParam)

	if err != nil {
		return
	}

	blog, err = c.blogSvc.GetBlogByID(id)

	if err != nil {
		return
	}

	message = "successfully fetch blog"
}

func (c *BlogController) ListBlogs(ctx *gin.Context) {
	var (
		code     int    = 500
		message  string = "failed to fetch all blogs"
		blogs    []domain.Blog
		err      error  = nil
		resQuery string = ctx.Query("result")
	)

	sendResp := func() {
		ctx.HTML(code, "Base", gin.H{
			"Title":    "List Blogs",
			"Content":  "ListBlogs",
			"Navs":     constants.Navs,
			"Err":      err,
			"Message":  message,
			"ResQuery": resQuery,
			"Blogs":    blogs,
		})
	}

	defer sendResp()

	blogs, err = c.blogSvc.GetAllBlogs()
	code = domain.GetCode(err)

	if err != nil {
		return
	}

	message = "successfully fetch all blogs"
}

func (c *BlogController) ShowBlog(ctx *gin.Context) {
	var (
		message string = "failed to fetch blog"
		blog    domain.Blog
		err     error = nil

		idParam  = ctx.Param("id")
		id       int
		resQuery string = ctx.Query("result")
	)

	sendResp := func() {
		ctx.HTML(http.StatusOK, "Base", gin.H{
			"Title":    "HTMX Go Blog",
			"Content":  "ShowBlog",
			"Navs":     constants.Navs,
			"Err":      err,
			"Blog":     blog,
			"ResQuery": resQuery,
			"Message":  message,
		})
	}

	defer sendResp()
	id, err = strconv.Atoi(idParam)

	if err != nil {
		return
	}

	blog, err = c.blogSvc.GetBlogByID(id)

	if err != nil {
		return
	}

	message = "successfully fetch blog"
}

func (c *BlogController) CreateBlog(ctx *gin.Context) {
	var (
		code    int    = 303
		message string = "failed to create blog"
		blog    domain.Blog
		err     error = nil
	)

	sendResp := func() {
		ctx.Redirect(code, fmt.Sprintf("/blogs?result=%s", message))
	}

	defer sendResp()

	if err = ctx.ShouldBind(&blog); err != nil {
		return
	}

	err = c.blogSvc.CreateBlog(&blog)

	if err != nil {
		return
	}

	message = "successfully create blog"
}

func (c *BlogController) UpdateBlog(ctx *gin.Context) {
	var (
		code    int    = 303
		message string = "failed to update blog"
		blog    domain.Blog
		err     error = nil

		idParam = ctx.Param("id")
		id      int
	)

	sendResp := func() {
		ctx.Redirect(code, fmt.Sprintf("/blogs/%v?result=%s", id,message))
		// ctx.Redirect(code, fmt.Sprintf("/blogs/%v?result=%s", id, message))
	}

	defer sendResp()
	id, err = strconv.Atoi(idParam)

	if err != nil {
		return
	}

	if err := ctx.ShouldBind(&blog); err != nil {
		return
	}

	blog.ID = id

	err = c.blogSvc.UpdateBlog(&blog)

	if err != nil {
		return
	}

	message = "successfully update blog"
}

func (c *BlogController) DeleteBlog(ctx *gin.Context) {
	var (
		code    int    = 303
		message string = "failed to delete blog"
		err     error  = nil

		idParam = ctx.Param("id")
		id      int
	)

	sendResp := func() {
		ctx.Redirect(code, fmt.Sprintf("/blogs?result=%s", message))
	}

	defer sendResp()
	id, err = strconv.Atoi(idParam)

	if err != nil {
		return
	}

	err = c.blogSvc.DeleteBlog(id)

	if err != nil {
		return
	}

	message = "successfully delete blog"
}
