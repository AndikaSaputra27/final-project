package entity

// Booking struct represents a booking entity
type Booking struct {
	ID     string  `json:"id"`
	Name   string  `json:"name"`
	Date   string  `json:"date"`
	Amount float64 `json:"amount"`
	Status string  `json:"status"` // Added Status field
}
