package web

import (
	"encoding/json"
	"io/ioutil"
	"library/structs"
	"net/http"
)

func getBookList() []structs.Book {
	var bookList []structs.Book

	resp, err := http.Get("http://127.0.0.1:8445/api/books")

	if err != nil {
		print(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		print(err)
	}
	json.Unmarshal(body, &bookList)

	return bookList
}

func getBookHistory(id string) []structs.History {
	var bookHistory []structs.History

	resp, err := http.Get("http://127.0.0.1:8445/api/book/" + id + "/history")

	if err != nil {
		print(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		print(err)
	}
	json.Unmarshal(body, &bookHistory)

	return bookHistory
}

func getBook(id string) structs.Book {
	var book structs.Book

	resp, err := http.Get("http://127.0.0.1:8445/api/book/" + id)

	if err != nil {
		print(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		print(err)
	}
	json.Unmarshal(body, &book)

	return book
}
