package web

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

var (
	webrouter *gin.Engine
)

func SetUpWebRoutes(router *gin.Engine, routePath string) {
	router.LoadHTMLGlob("templates/*")
	router.StaticFile("/styles.css", "./resources/styles.css")
	router.StaticFile("/favicon.ico", "./resources/favicon.ico")
	router.GET("/index.html", func(c *gin.Context) {
		bookList := getBookList()
		c.HTML(http.StatusOK, "index.tmpl", gin.H{
			"title":    "The Library",
			"category": "Books",
			"books":    bookList,
		})
	})
	router.GET("/book/:uid/index.html", bookDetail)
}

func SetupWebRouter(router *gin.Engine) {
	webrouter = router
	webRouteHeader := "/"
	SetUpWebRoutes(webrouter, webRouteHeader)
}

func bookDetail(c *gin.Context) {
	bookID := c.Param("uid")
	_, err := strconv.Atoi(bookID) // it is an int?
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	bookHistory := getBookHistory(bookID)
	book := getBook(bookID)
	c.HTML(http.StatusOK, "book.tmpl", gin.H{
		"title":    "The Library",
		"category": "Books",
		"history":  bookHistory,
		"book":     book,
	})
}
