package users

import (
	"encoding/json"
	"example/helloWorldServer/handler"
	"io/ioutil"
	"net/http"
	"reflect"
	"time"

	"github.com/gorilla/context"
)

func PatchUsersHandlers(w http.ResponseWriter, r *http.Request) {
	// set the time
	t := time.Now()

	// set user id from token
	UserReqId := context.Get(r, "ID")

	// checks if the user id is valid
	if UserReqId == nil || UserReqId == "newUserDetected" {
		resp := handler.CustomResponseForUserUnauthorized{
			Status:           "Unauthorized",
			StatusCode:       http.StatusUnauthorized,
			Message:          "You are not authorized to acces this site. Please sign up or sign in before hand.",
			ReadableDateTime: t.Format(time.UnixDate),
			DateTime:         t,
			UserID:           UserReqId,
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(resp)
		return
	}

	// get r.Body
	var reqBodyData handler.UsersWithoutID
	if err := json.NewDecoder(r.Body).Decode(&reqBodyData); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// read users file json
	content, err := ioutil.ReadFile("./database/user.json")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	//define users
	var users []handler.Users
	var user handler.Users

	//unmarshal users json from read content
	json.Unmarshal(content, &users)

	// run user finder in the database
	index, found := databaseUserFinder(users, UserReqId)

	if !found {
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

	// setting the user struct
	user = users[index]

	// updating data process
	updatedUser := updateUserData(user, reqBodyData)

	users[index] = updatedUser

	// prepare response format
	resp := handler.CustomResponseForUserWithBooks{
		Status:           "OK",
		StatusCode:       http.StatusOK,
		Message:          "Hello, there!",
		ReadableDateTime: t.Format(time.UnixDate),
		DateTime:         t,
		UserID:           UserReqId,
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resp)
}

func updateUserData(user handler.Users, reqBodyData handler.UsersWithoutID) handler.Users {
	v := reflect.ValueOf(&reqBodyData).Elem()
	valueOfUser := reflect.ValueOf(&user)
	typeOfUser := reflect.TypeOf(user)
	for i := 0; i < v.NumField(); i++ {
		if valueOfUser.Kind() == reflect.Struct {
			f := valueOfUser.FieldByName(typeOfUser.Field(i).Name)
			if f.IsValid() {
				if f.CanSet() {
					if f.Kind() == reflect.String {
						x := v.Field(i)
						f.SetString(x.String())
					}
				}
			}
		}
	}
	return user
}

func databaseUserFinder(users []handler.Users, UserReqId interface{}) (int, bool) {
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
			userIndex = -1
			break
		}
		continue
	}
	return userIndex, userFound
}
