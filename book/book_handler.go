package book

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func createBook(c *gin.Context) {
	var book Book

	if err := c.ShouldBindJSON(&book); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	book.createBook()
}
