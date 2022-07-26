package middleware

import (
	"encoding/json"
	"example/helloWorldServer/handler"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/gorilla/context"

	"github.com/golang-jwt/jwt"
)

func Verification() Middleware {
	return func(f http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			AuthHeader := r.Header.Get("Authorization")
			token := strings.Split(AuthHeader, "Bearer ")[1]

			at(time.Unix(0, 0), func() {
				token, err := jwt.ParseWithClaims(token, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
					return SignKey, nil
				})
				// Checks if user is in user.json by having cases.
				// Case 1 identifies that user is in user.json,
				// Case 2 identifies that the function databaseChecker
				if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
					switch databaseChecker(claims.ID) {
					case 1:
						context.Set(r, "ID", claims.ID)
					case -1:
						http.Error(w, "Invalid ID", http.StatusBadRequest)
					case 0:
						context.Set(r, "ID", "newUserDetected")
					case 2:
						context.Set(r, "ID", nil)
					}

				} else {
					http.Error(w, err.Error(), http.StatusBadGateway)
				}
			})

			f(w, r)
		}
	}
}

// Override time value for tests.  Restore default value after.
func at(t time.Time, f func()) {
	jwt.TimeFunc = func() time.Time {
		return t
	}
	f()
	jwt.TimeFunc = time.Now
}

func databaseChecker(id string) int {
	var users []handler.Users
	content, err := ioutil.ReadFile("./database/user.json")
	if err != nil {
		log.Fatalf("Error reading")
		return -1
	}

	json.Unmarshal(content, &users)

	for index, value := range users {
		if value.UserID == id {
			return 1
		} else {
			if index == len(users)-1 {
				return 0
			}
		}
	}
	return 2
}
