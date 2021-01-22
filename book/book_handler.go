package book

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func createBook(c *gin.Context) {
	var book Book

	if err := c.ShouldBindJSON(&book); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := book.createBook()
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
	}
	c.Status(http.StatusOK)
}

func readBook(c *gin.Context) {
	var book Book

	bookID := c.Param("uid")
	id, err := strconv.Atoi(bookID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Bad book id!"})
		return
	}

	book.ID = id
	err = book.readBook()
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, book)
}

func getBookList(c *gin.Context) {
	bookList, err := getBookListSQL()
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, bookList)
}
