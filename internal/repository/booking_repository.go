package repository

import (
	"booking-system/internal/entity"
	"fmt"
)

type BookingRepository interface {
	Save(booking entity.Booking) (entity.Booking, error)
	FindByID(id int) (entity.Booking, error)
}

type InMemoryBookingRepository struct {
	bookings map[int]entity.Booking
	nextID   int
}

func NewInMemoryBookingRepository() *InMemoryBookingRepository {
	return &InMemoryBookingRepository{
		bookings: make(map[int]entity.Booking),
		nextID:   1,
	}
}

func (r *InMemoryBookingRepository) Save(booking entity.Booking) (entity.Booking, error) {
	booking.ID = r.nextID
	r.bookings[r.nextID] = booking
	r.nextID++
	return booking, nil
}

func (r *InMemoryBookingRepository) FindByID(id int) (entity.Booking, error) {
	booking, exists := r.bookings[id]
	if !exists {
		return entity.Booking{}, ErrBookingNotFound
	}
	return booking, nil
}

var ErrBookingNotFound = fmt.Errorf("booking not found")
