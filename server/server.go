package server

import (
	"github.com/gin-gonic/gin"
	"library/book"
)

var (
	router = gin.Default()
)

func SetupRouter() {
	apiRouteHeader := "/api/"
	router.Use(allowCorsOptions())
	book.SetUpBookRoutes(router, apiRouteHeader)
	router.Run(":8445")
}
