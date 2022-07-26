package admin

import (
	"encoding/json"
	"example/helloWorldServer/handler"
	"io/ioutil"
	"net/http"
	"time"
)

func remove(newJsonBook []handler.Book, index int) []handler.Book {
	newJsonBook[index] = newJsonBook[len(newJsonBook)-1]
	return newJsonBook[:len(newJsonBook)-1]
}

func DeleteBooksHandler(w http.ResponseWriter, r *http.Request) {
	// get time
	t := time.Now()

	// get the id of the book that wants to be deleted
	id := r.URL.Query().Get("id")

	// define book json type
	var jsonBook []handler.Book

	// reads json book
	content, err := ioutil.ReadFile("./database/book.json")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.Unmarshal(content, &jsonBook)

	// get the length of the jsonBook

	// finds the id of the book that wants to be deleted
	var jsonDeletedBook handler.Book
	for index, value := range jsonBook {
		if value.ID == id {
			jsonDeletedBook = jsonBook[index]
			jsonBook = remove(jsonBook, index)
		}
		if index == len(jsonBook)-1 {
			// sends back response to user for not found book id
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusNotFound)
			resp := &handler.CustomResponse{
				Status:           "ID Not Found",
				StatusCode:       http.StatusNotFound,
				Message:          "There is no book id that matches!",
				ReadableDateTime: t.Format(time.UnixDate),
				DateTime:         t,
			}
			json.NewEncoder(w).Encode(resp)
			return
		}
		continue
	}

	// save deleted jsonBook to json book
	file, _ := json.MarshalIndent(jsonBook, "", "	")
	ioutil.WriteFile("./database/book.json", file, 0644)

	resp := &handler.CustomResponseWithSingleData{
		Status:           "OK",
		StatusCode:       http.StatusOK,
		Message:          "Successfully deleted book with the given id.",
		ReadableDateTime: t.Format(time.UnixDate),
		DateTime:         t,
		DataResp:         jsonDeletedBook,
	}

	// sends back response to user
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resp)
}
