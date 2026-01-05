package utils

import (
	"fmt"
	"net/http"
	"web-api-student/model"
)

func StatsHandler(write http.ResponseWriter, request *http.Request) {
	// Card
	subjects := model.Cards
	subjectCount := len(subjects)
	//Check
	if subjectCount == 0 {
		fmt.Fprintln(write, "There is no data for statistics")
		return
	}
	// AVG
	var sum int = 0
	for _, grades := range subjects {
		sum += grades.Grade
	}
	avgGrade := (sum / subjectCount)

	// Search the lowest and highest score
	minVal := subjects[0].Grade
	maxVal := subjects[0].Grade
	worstSubject := subjects[0].Subject
	bestSubject := subjects[0].Subject
	for _, grade := range subjects {
		if grade.Grade < minVal {
			minVal = grade.Grade
			worstSubject = grade.Subject
		}
		if grade.Grade > maxVal {
			maxVal = grade.Grade
			bestSubject = grade.Subject
		}

	}

	//Enter info
	fmt.Fprintf(write, `
 === СТАТИСТИКА НАВЧАННЯ ===

Всього предметів: %d
Середній бал: %d/12
Найкраща оцінка: %d/12 (%s)
Найгірша оцінка: %d/12 (%s)
 `,
		subjectCount, avgGrade, maxVal, bestSubject, minVal, worstSubject)
}
