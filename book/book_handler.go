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
		return
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

func updateBook(c *gin.Context) {
	var book Book

	bookID := c.Param("uid")
	id, err := strconv.Atoi(bookID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Bad book id!"})
		return
	}
	if err := c.ShouldBindJSON(&book); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	book.ID = id
	if book.ID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Bad book id"})
		return
	}

	err = book.updateBook()
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	c.Status(http.StatusOK)
}

func deleteBook(c *gin.Context) {
	var book Book

	bookID := c.Param("uid")
	i, err := strconv.Atoi(bookID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	book.ID = i
	if book.ID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Bad book id"})
		return
	}

	err = book.deleteBook()
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	c.Status(http.StatusOK)
}

func getBookList(c *gin.Context) {
	bookList, err := getBookListSQL()
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, bookList)
}

func checkoutBook(c *gin.Context) {
	bookID := c.Param("uid")
	id, err := strconv.Atoi(bookID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	status, err := bookStatus(id)
	if status != Available || err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Book is already check out!"})
		return
	}

	checkoutBookSql(id)
	c.Status(http.StatusOK)
}

func checkinBook(c *gin.Context) {
	bookID := c.Param("uid")
	id, err := strconv.Atoi(bookID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	status, err := bookStatus(id)
	if status != Unavailable || err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Book is not checked out!"})
		return
	}

	checkinBookSql(id)
	c.Status(http.StatusOK)
}
