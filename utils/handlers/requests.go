package utils

import (
	"fmt"
	"net/http"
	"web-api-student/model"
)

func HomeHandler(write http.ResponseWriter, request *http.Request) {
	//Card
	subjects := model.Cards
	subjectCount := len(subjects)

	// Check
	if subjectCount == 0 {
		fmt.Fprintln(write, "Small data for statistics")
		return
	}

	// Enter information my cartoreck subject
	fmt.Fprintf(write, " === MY FILE OF SUBJECTS ===\n")
	fmt.Fprintf(write, "Subjects %d\n", subjectCount)

	// Conclusion all items
	for _, subject := range subjects {
		fmt.Fprintf(write, "%d. %s - Grades: %d/12\n",
			subject.ID,
			subject.Subject,
			subject.Grade,
		)
	}
	fmt.Fprintf(write, "All subjects: %d \n", subjectCount)

}
