package server

import (
	"github.com/gin-gonic/gin"
	"library/book"
)

var (
	router = gin.Default()
)

func SetupRouter() *gin.Engine {
	apiRouteHeader := "/api/"
	router.Use(allowCorsOptions())
	book.SetUpBookRoutes(router, apiRouteHeader)
	return router
}
