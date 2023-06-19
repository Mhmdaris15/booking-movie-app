package models

type Ticket struct {
	ID              string
	MovieID         string
	UserID          string
	ShowtimeID      string
	TransactionDate string
	TotalPrice      float64
	// Add other fields as per your requirements
}
