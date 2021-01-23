package book

import "github.com/gin-gonic/gin"

func SetUpBookRoutes(router *gin.Engine, routePath string) {
	router.GET(routePath+"books", getBookList)
	router.POST(routePath+"book", createBook)
	router.GET(routePath+"book/:uid", readBook)
	router.PUT(routePath+"book/:uid", updateBook)
	router.DELETE(routePath+"book/:uid", deleteBook)
	router.POST(routePath+"book/:uid/checkout", checkoutBook)
	router.POST(routePath+"book/:uid/checkin", checkinBook)
}
