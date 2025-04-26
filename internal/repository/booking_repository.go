package repository

import (
	"fmt"

	"github.com/AndikaSaputra27/booking-system/internal/entity"
)

// BookingRepository interface with Save, FindByID, GetAll, and Delete methods
type BookingRepository interface {
	Save(booking entity.Booking) (entity.Booking, error)
	FindByID(id string) (entity.Booking, error)
	GetAll() ([]entity.Booking, error)
	Delete(id string) error // Added Delete method
}

// InMemoryBookingRepository is an in-memory implementation of BookingRepository
type InMemoryBookingRepository struct {
	bookings map[string]entity.Booking
	nextID   int
}

// NewInMemoryBookingRepository creates a new InMemoryBookingRepository instance
func NewInMemoryBookingRepository() *InMemoryBookingRepository {
	return &InMemoryBookingRepository{
		bookings: make(map[string]entity.Booking),
		nextID:   1,
	}
}

// Save saves a booking and returns it with an assigned ID
func (r *InMemoryBookingRepository) Save(booking entity.Booking) (entity.Booking, error) {
	id := fmt.Sprintf("%d", r.nextID)
	booking.ID = id
	r.bookings[id] = booking
	r.nextID++
	return booking, nil
}

// FindByID finds a booking by its ID
func (r *InMemoryBookingRepository) FindByID(id string) (entity.Booking, error) {
	booking, exists := r.bookings[id]
	if !exists {
		return entity.Booking{}, fmt.Errorf("booking with id %s not found", id)
	}
	return booking, nil
}

// GetAll retrieves all bookings
func (r *InMemoryBookingRepository) GetAll() ([]entity.Booking, error) {
	var allBookings []entity.Booking
	for _, booking := range r.bookings {
		allBookings = append(allBookings, booking)
	}
	return allBookings, nil
}

// Delete deletes a booking by its ID
func (r *InMemoryBookingRepository) Delete(id string) error {
	_, exists := r.bookings[id]
	if !exists {
		return fmt.Errorf("booking with id %s not found", id)
	}
	delete(r.bookings, id) // Delete the booking from the map
	return nil
}
