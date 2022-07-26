package handler

import "time"

type (
	CustomResponseWithData struct {
		Status           string    `json:"status"`
		StatusCode       int       `json:"statusCode"`
		Message          string    `json:"message"`
		ReadableDateTime string    `json:"readableDateTime"`
		DateTime         time.Time `json:"dateTime"`
		DataResp         []Book    `json:"dataResp"`
	}

	CustomResponseWithSingleData struct {
		Status           string    `json:"status"`
		StatusCode       int       `json:"statusCode"`
		Message          string    `json:"message"`
		ReadableDateTime string    `json:"readableDateTime"`
		DateTime         time.Time `json:"dateTime"`
		DataResp         Book      `json:"dataResp"`
	}

	CustomResponse struct {
		Status           string    `json:"status"`
		StatusCode       int       `json:"statusCode"`
		Message          string    `json:"message"`
		ReadableDateTime string    `json:"readableDateTime"`
		DateTime         time.Time `json:"dateTime"`
	}

	Book struct {
		ID           string            `json:"id"`
		ProductTheme string            `json:"productTheme"`
		ProductDesc  string            `json:"productDesc"`
		Metadata     map[string]string `json:"metadata"`
	}

	CustomResponseForUserUnauthorized struct {
		Status           string      `json:"status"`
		StatusCode       int         `json:"statusCode"`
		Message          string      `json:"message"`
		ReadableDateTime string      `json:"readableDateTime"`
		DateTime         time.Time   `json:"dateTime"`
		UserID           interface{} `json:"userID"`
	}

	CustomResponseForUserWithBooks struct {
		Status           string      `json:"status"`
		StatusCode       int         `json:"statusCode"`
		Message          string      `json:"message"`
		ReadableDateTime string      `json:"readableDateTime"`
		DateTime         time.Time   `json:"dateTime"`
		UserID           interface{} `json:"userID"`
		UserBooks        []Book      `json:"userBooks"`
	}

	Users struct {
		UserID    string `json:"userId"`
		Name      string `json:"name"`
		Email     string `json:"email"`
		UserBooks []Book `json:"userBooks"`
	}

	UsersWithoutID struct {
		Name  string `json:"name"`
		Email string `json:"email"`
	}

	CustomResponseNewUser struct {
		Status           string      `json:"status"`
		StatusCode       int         `json:"statusCode"`
		Message          string      `json:"message"`
		ReadableDateTime string      `json:"readableDateTime"`
		DateTime         time.Time   `json:"dateTime"`
		UserID           interface{} `json:"userID"`
	}
)
