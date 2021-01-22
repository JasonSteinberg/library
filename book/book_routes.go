package book

import "github.com/gin-gonic/gin"

func SetUpBookRoutes(router *gin.Engine, routePath string) {
	router.GET(routePath+"books", getBookList)
	router.POST(routePath+"book", createBook)
	router.GET(routePath+"book/:uid", readBook)
}
