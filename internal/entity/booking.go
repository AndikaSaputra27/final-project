package entity

type Booking struct {
	ID        uint   `json:"id" gorm:"primaryKey"`
	UserID    uint   `json:"user_id"`
	ServiceID uint   `json:"service_id"`
	Time      string `json:"time"`
	Status    string `json:"status"`
}
