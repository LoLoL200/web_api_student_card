package utils

import (
	"fmt"
	"net/http"
	"strconv"
	"web-api-student/model"
)

func ViewHandler(write http.ResponseWriter, request *http.Request) {
	// Cards
	subjects := model.Cards

	// Getting the value of a parameter with a name
	idStr := request.URL.Query().Get("id")
	//Check
	if idStr == "" {
		http.Error(write, "Missing 'id' parameter", http.StatusBadRequest)
		return
	}

	// Convert a string to an integer
	id, err := strconv.Atoi(idStr)
	// Check
	if err != nil {
		http.Error(write, "Invalid 'id' format, must be an integer", http.StatusBadRequest)
		return
	}

	// Search for an item by ID
	found := false
	for _, subject := range subjects {
		if subject.ID == id {
			found = true
			fmt.Fprintf(write, `
				Subject: %s
				Grade: %d/12
				Note: %s
			`, subject.Subject, subject.Grade, subject.Note)
			return
		}

	}
	// Check
	if !found {
		http.Error(write, "Subject not found", http.StatusNotFound)
	}

}
