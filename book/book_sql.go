package book

import (
	"github.com/gin-gonic/gin"
	"library/author"
	"library/database"
	"library/structs"
	"log"
	"net/http"
)

type Book struct {
	structs.Book
}

func getBookList(c *gin.Context) {
	var bookList []structs.Book

	db := database.GetSqlReadDB()
	rows, err := db.Query(
		`select books.id, title, description, isbn, group_concat(a.name) as author
				from books
				inner join book_author ba on books.id = ba.book_id
				inner join author a on ba.author_id = a.id
				group by books.id`)
	if err != nil {
		log.Println("getBookList ", err)
	}
	defer rows.Close()

	for rows.Next() {
		book := structs.Book{}
		err := rows.Scan(&book.ID, &book.Title, &book.Description, &book.ISBN, &book.Authors)
		if err != nil {
			log.Println("Oh no from getBookList", err)
		}
		bookList = append(bookList, book)
	}

	c.JSON(http.StatusOK, bookList)
}

func (b *Book) createBook() {
	db := database.GetSqlWriteDB()

	result, err := db.Exec(`insert into books (title, description, isbn)
				values (?,?,?)`, b.Title, b.Description, b.ISBN)
	if err != nil {
		log.Println("createBook has an issue ", err)
	}

	i, err := result.LastInsertId()
	if err != nil {
		log.Println("createBook has an issue ", err)
	}

	author.LinkAuthorsToBook(b.Authors, int(i)) // unsafe cast but okay here
}
