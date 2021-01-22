package server

import (
	"github.com/gin-gonic/gin"
)

func allowCorsOptions() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Access-Control-Allow-Methods", "*")
		c.Header("Access-Control-Allow-Origin", `*`)
	}
}
