package users

import (
	"encoding/json"
	"example/helloWorldServer/handler"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/google/uuid"
)

func PostUserHandler(w http.ResponseWriter, r *http.Request) {
	t := time.Now()
	userID := uuid.New()

	// define jsonBook for writing file
	var users []handler.Users
	var newUser handler.Users

	// input userID
	newUser.UserID = userID.String()

	// read json book file
	content, err := ioutil.ReadFile("./database/user.json")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	json.Unmarshal(content, &users)

	// input file into user.json
	users = append(users, newUser)
	file, _ := json.MarshalIndent(users, "", "	")
	ioutil.WriteFile("./database/user.json", file, 0644)

	resp := &handler.CustomResponseNewUser{
		Status:           "OK - NewUser",
		StatusCode:       http.StatusFound,
		Message:          "Welcome new user",
		ReadableDateTime: t.Format(time.UnixDate),
		DateTime:         t,
		UserID:           userID,
	}

	// sends back response to user
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resp)
}
