package web

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

var (
	webrouter *gin.Engine
)

func SetUpWebRoutes(router *gin.Engine, routePath string) {
	router.LoadHTMLGlob("templates/*")
	router.StaticFile("/styles.css", "./resources/styles.css")
	router.StaticFile("/favicon.ico", "./resources/favicon.ico")
	router.GET("/index.html", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl", gin.H{
			"title":    "The Library",
			"category": "Books",
		})
	})
	router.GET("/book/:uid/index.html", func(c *gin.Context) {
		c.HTML(http.StatusOK, "book.tmpl", gin.H{
			"title":    "The Library",
			"category": "Books",
		})
	})
}

func SetupWebRouter(router *gin.Engine) {
	webrouter = router
	webRouteHeader := "/"
	SetUpWebRoutes(webrouter, webRouteHeader)
}
