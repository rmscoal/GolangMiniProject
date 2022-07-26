package admin

import (
	"encoding/json"
	"example/helloWorldServer/handler"
	"net/http"
	"time"
)

func HealthCheckHandler(w http.ResponseWriter, r *http.Request) {
	// Setting the response content to be application/json
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	// Get the current local time
	t := time.Now()
	// Defining a custom response to the user
	resp := &handler.CustomResponse{
		Status:           "OK",
		StatusCode:       http.StatusOK,
		Message:          "Healthy",
		ReadableDateTime: t.Format(time.UnixDate),
		DateTime:         t,
	}

	// Returns response to the user
	json.NewEncoder(w).Encode(resp)
}
