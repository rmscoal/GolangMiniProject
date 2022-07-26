package admin

import (
	"encoding/json"
	"example/helloWorldServer/handler"
	"io/ioutil"
	"net/http"
	"time"
)

func PostBooksHandler(w http.ResponseWriter, r *http.Request) {
	// get time
	t := time.Now()

	// define request body
	var reqBodyBook handler.Book

	// parse json request
	err := json.NewDecoder(r.Body).Decode(&reqBodyBook)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// define jsonBook for writing file
	var jsonBook []handler.Book
	// read json book file
	content, err := ioutil.ReadFile("./database/book.json")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	// save it in jsonBook variable
	json.Unmarshal(content, &jsonBook)

	// checks if the reqBodyBook is already in jsonBook
	for _, value := range jsonBook {
		if value.ID == reqBodyBook.ID || value.Metadata["name"] == reqBodyBook.Metadata["name"] {
			resp := &handler.CustomResponse{
				Status:           "Fail",
				StatusCode:       http.StatusFound,
				Message:          "A book with the same ID and metadata.name is already registered.",
				ReadableDateTime: t.Format(time.UnixDate),
				DateTime:         t,
			}
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusFound)
			json.NewEncoder(w).Encode(resp)
			return
		}
		continue
	}
	jsonBook = append(jsonBook, reqBodyBook)

	file, _ := json.MarshalIndent(jsonBook, "", "	")

	ioutil.WriteFile("./database/book.json", file, 0644)

	resp := &handler.CustomResponseWithSingleData{
		Status:           "OK",
		StatusCode:       http.StatusOK,
		Message:          "Book added to database",
		ReadableDateTime: t.Format(time.UnixDate),
		DateTime:         t,
		DataResp:         reqBodyBook,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resp)
}
