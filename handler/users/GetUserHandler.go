package users

import (
	"encoding/json"
	"example/helloWorldServer/handler"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/gorilla/context"
)

func GetUserHandler(w http.ResponseWriter, r *http.Request) {
	// get time
	t := time.Now()

	UserReqId := context.Get(r, "ID")
	if UserReqId == nil {
		resp := handler.CustomResponseForUserUnauthorized{
			Status:           "Unauthorized",
			StatusCode:       http.StatusUnauthorized,
			Message:          "You are not authorized to access this site.",
			ReadableDateTime: t.Format(time.UnixDate),
			DateTime:         t,
			UserID:           UserReqId,
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(resp)
		return
	}
	if UserReqId == "newUserDetected" {
		http.Redirect(w, r, "http://localhost:8000/users/newuser", http.StatusFound)
		return
	}

	// read users file json
	content, err := ioutil.ReadFile("./database/user.json")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// define users
	var users []handler.Users

	// unmarshal users json from read content
	json.Unmarshal(content, &users)

	// find the user in users
	var userFound bool
	var userIndex int
	for index, value := range users {
		if value.UserID == UserReqId {
			userFound = true
			userIndex = index
			break
		}
		if index == len(users)-1 {
			userFound = false
			return
		}
		continue
	}

	if !userFound {
		resp := handler.CustomResponseForUserUnauthorized{
			Status:           "UserNotFound",
			StatusCode:       http.StatusUnauthorized,
			Message:          "User ID is not registered in database.",
			ReadableDateTime: t.Format(time.UnixDate),
			DateTime:         t,
			UserID:           UserReqId,
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(resp)
		return
	}

	// make a book json list for user
	var books []handler.Book = users[userIndex].UserBooks

	// prepare response format
	resp := handler.CustomResponseForUserWithBooks{
		Status:           "OK",
		StatusCode:       http.StatusOK,
		Message:          "Hello, there!",
		ReadableDateTime: t.Format(time.UnixDate),
		DateTime:         t,
		UserID:           UserReqId,
		UserBooks:        books,
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resp)
}
