package admin

import (
	"encoding/json"
	"example/helloWorldServer/handler"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

// function AllBooksHandler
func GetBooksHandler(w http.ResponseWriter, r *http.Request) {
	// define the json book
	var jsonBook []handler.Book

	// get book list in form of json
	content, err := ioutil.ReadFile("./database/book.json")
	if err != nil {
		log.Fatalf("There is an error happening %s", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// unmarshal jsonBook
	json.Unmarshal(content, &jsonBook)

	// get the id of the query param
	id := r.URL.Query().Get("id")
	if id != "" {
		// filtering out jsonBook based on the id from the query
		// firstly define a new json
		var jsonBookResp []handler.Book
		for _, value := range jsonBook {
			if value.ID == id {
				jsonBookResp = append(jsonBookResp, value)
			}
			continue
		}

		if len(jsonBookResp) == 0 {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusBadRequest)
			t := time.Now()
			resp := &handler.CustomResponse{
				Status:           "Bad Query",
				StatusCode:       http.StatusBadRequest,
				Message:          "There was an error in your query",
				ReadableDateTime: t.Format(time.UnixDate),
				DateTime:         t,
			}
			json.NewEncoder(w).Encode(resp)
			return
		}

		jsonBook = jsonBookResp
	}

	// get the current time
	t := time.Now()

	// define the repsonse
	resp := &handler.CustomResponseWithData{
		Status:           "OK",
		StatusCode:       http.StatusOK,
		Message:          "Query data successful!",
		ReadableDateTime: t.Format(time.UnixDate),
		DateTime:         t,
		DataResp:         jsonBook,
	}

	// set header response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	// returns repsonse to the users
	json.NewEncoder(w).Encode(resp)
}
