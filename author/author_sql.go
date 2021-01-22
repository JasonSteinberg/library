package author

import (
	"library/database"
	"log"
	"strings"
)

func getAuthorID(name string) int {
	db := database.GetSqlReadDB()
	rows, err := db.Query(`select id from author where name=?;`, name)
	defer rows.Close()
	if err != nil {
		log.Println("Error getting author ", err)
		return 0
	}
	if !rows.Next() {
		return 0
	}
	var aid int
	rows.Scan(&aid)
	return aid
}

func createAuthor(name string) int {
	db := database.GetSqlWriteDB()
	result, err := db.Exec(`insert into author (name) values (?);`, name)
	if err != nil {
		log.Println("Error creating author ", err)
		return 0
	}
	i, err := result.LastInsertId() // unsafe but ok here
	if err != nil {
		log.Println("Error getting author id ", err)
	}
	return int(i)
}

func getOrCreateAuthor(name string) int {
	aid := getAuthorID(name)
	if aid == 0 {
		aid = createAuthor(name)
	}
	return aid
}

func linkAuthorToBook(authorID, bookID int) {
	db := database.GetSqlWriteDB()
	_, err := db.Exec(`insert into book_author (author_id, book_id) values (?,?);`, authorID, bookID)
	if err != nil {
		log.Println("Error linking author to book", err)
		return
	}
}

func LinkAuthorsToBook(name string, bookID int) {
	authors := strings.Split(name, ",")

	for _, v := range authors {
		aid := getOrCreateAuthor(v)
		linkAuthorToBook(aid, bookID)
	}
}

func RemoveLinksToBook(bookID int) {
	db := database.GetSqlWriteDB()
	_, err := db.Exec(`delete from book_author where book_id=?;`, bookID)
	if err != nil {
		log.Println("Error deleting links to book", err)
		return
	}
}

func UpdateLinkAuthorsToBook(name string, bookID int) {
	RemoveLinksToBook(bookID)
	LinkAuthorsToBook(name, bookID)
}
