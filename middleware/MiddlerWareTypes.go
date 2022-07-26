package middleware

import (
	"net/http"

	"github.com/golang-jwt/jwt"
)

type (
	CustomClaims struct {
		ID    string `json:"id"`
		Name  string `json:"name"`
		Email string `json:"email"`
		jwt.StandardClaims
	}

	Middleware func(http.HandlerFunc) http.HandlerFunc
)

var (
	SignKey = []byte("G0l4angVSC043")
)
