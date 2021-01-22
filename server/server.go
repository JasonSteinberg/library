package server

import (
	"github.com/gin-gonic/gin"
)

var (
	router = gin.Default()
)

func SetupRouter() {
	//apiRouteHeader := "/api/"

	router.Run(":8445")
}
