package utils

import (
	"fmt"
	"net/http"
	"strconv"
	"web-api-student/model"
)

func AddHandler(write http.ResponseWriter, request *http.Request) {
	// Check
	if request.Method != http.MethodPost {
		http.Error(write, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	write.Header().Set("Content-Type", "text/plain; charset=utf-8")
	if err := request.ParseForm(); err != nil {
		http.Error(write, "Unable read form", http.StatusBadRequest)
		return
	}

	// Subject
	subject := request.FormValue("subject")
	if subject == "" {
		http.Error(write, "Please enter a valid subject", http.StatusAccepted)
		return
	}

	// Grade
	gradeStr := request.FormValue("grade")
	grade, err := strconv.Atoi(gradeStr)
	if err != nil {
		http.Error(write, "Error1000", http.StatusBadRequest)
		return
	}
	if grade >= 13 {
		http.Error(write, "Enter the grade < 12", http.StatusAccepted)
		return
	}

	note := request.FormValue("note")
	if note == "" {
		http.Error(write, "Please enter a valid note", http.StatusAccepted)
		return
	}

	card := model.Card{
		ID:      len(model.Cards) + 1,
		Subject: subject,
		Grade:   grade,
		Note:    note,
	}
	model.Cards = append(model.Cards, card)

	fmt.Fprintf(write, `
	✅ SUBJECT ADDED SUCCESSFULLY!

			ID: %d
			Subject: %s
			Grade: %d/12
			Note: %s
	`, card.ID, card.Subject, card.Grade, card.Note)

}
