package book

import (
	"library/author"
	"library/database"
	"library/structs"
	"log"
)

type Book struct {
	structs.Book
}

func getBookListSQL() ([]structs.Book, error) {
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
		return bookList, err
	}
	defer rows.Close()

	for rows.Next() {
		book := structs.Book{}
		err := rows.Scan(&book.ID, &book.Title, &book.Description, &book.ISBN, &book.Authors)
		if err != nil {
			log.Println("Oh no from getBookList", err)
			return bookList, err
		}
		bookList = append(bookList, book)
	}

	return bookList, nil
}

func (b *Book) createBook() error {
	db := database.GetSqlWriteDB()

	result, err := db.Exec(`insert into books (title, description, isbn)
				values (?,?,?)`, b.Title, b.Description, b.ISBN)
	if err != nil {
		log.Println("createBook has an issue ", err)
		return err
	}

	i, err := result.LastInsertId()
	if err != nil {
		log.Println("createBook has an issue ", err)
		return err
	}

	author.LinkAuthorsToBook(b.Authors, int(i)) // unsafe cast but okay here
	return nil
}

func (b *Book) readBook() error {
	db := database.GetSqlReadDB()

	rows, err := db.Query(`select title, description, isbn, group_concat(a.name) as author
				from books
				inner join book_author ba on books.id = ba.book_id
				inner join author a on ba.author_id = a.id
				where books.id = ?
				group by books.id`, b.ID)
	defer rows.Close()
	if err != nil {
		log.Println("readBook has an issue ", err)
		return err
	}

	if !rows.Next() {
		return nil
	}

	rows.Scan(&b.Title, &b.Description, &b.ISBN, &b.Authors)
	if err != nil {
		log.Println("readBook read issue ", err)
		return err
	}
	return nil
}

func (b *Book) updateBook() error {
	db := database.GetSqlWriteDB()

	result, err := db.Exec(`insert into books (title, description, isbn)
				values (?,?,?)`, b.Title, b.Description, b.ISBN)
	if err != nil {
		log.Println("updateBook has an issue ", err)
		return err
	}

	author.UpdateLinkAuthorsToBook(b.Authors, int(i)) // unsafe cast but okay here
	return nil
}

func (b *Book) deleteBook() error {
	db := database.GetSqlWriteDB()

	_, err := db.Exec(`delete from books where id=?`, b.ID)
	if err != nil {
		log.Println("deleteBook has an issue ", err)
		return err
	}

	author.RemoveLinksToBook(b.ID) // unsafe cast but okay here
	return nil
}
