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
