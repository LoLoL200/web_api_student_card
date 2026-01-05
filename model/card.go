package model

type Card struct {
	ID      int
	Subject string
	Grade   int
	Note    string
}

var Cards = []Card{
	{1, "Math", 11, "Hard subject"},
	{2, "Physics", 9, "Eusy subject"},
	{3, "IT", 5, "Interesting subject"},
}
