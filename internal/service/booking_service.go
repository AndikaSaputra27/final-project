package service

import (
	"errors"
	"strconv"

	"github.com/AndikaSaputra27/booking-system/internal/entity"
	"github.com/AndikaSaputra27/booking-system/internal/repository"
)

// BookingService defines the methods that can be performed on bookings.
type BookingService struct {
	BookingRepository repository.BookingRepository
}

func (s BookingService) CreateService(service *entity.Service) any {
	panic("unimplemented")
}

// NewBookingService creates a new instance of BookingService
func NewBookingService(bookingRepository repository.BookingRepository) *BookingService {
	return &BookingService{
		BookingRepository: bookingRepository,
	}
}

// CreateBooking creates a new booking in the system
func (s *BookingService) CreateBooking(booking entity.Booking) (entity.Booking, error) {
	createdBooking, err := s.BookingRepository.Save(booking)
	if err != nil {
		return entity.Booking{}, err
	}
	return createdBooking, nil
}

// GetBookingByID retrieves a booking by its ID
func (s *BookingService) GetBookingByID(id int) (entity.Booking, error) {
	// Convert the ID to a string if the repository expects a string
	idStr := strconv.Itoa(id)

	booking, err := s.BookingRepository.FindByID(idStr)
	if err != nil {
		return entity.Booking{}, err
	}
	return booking, nil
}

// GetAllBookings retrieves all bookings
func (s *BookingService) GetAllBookings() ([]entity.Booking, error) {
	bookings, err := s.BookingRepository.GetAll()
	if err != nil {
		return nil, err
	}
	return bookings, nil
}

// UpdateBooking updates an existing booking's details
func (s *BookingService) UpdateBooking(id int, updatedBooking entity.Booking) (entity.Booking, error) {
	idStr := strconv.Itoa(id)

	booking, err := s.BookingRepository.FindByID(idStr)
	if err != nil {
		return entity.Booking{}, errors.New("booking not found")
	}

	// Update field dari booking
	booking.Name = updatedBooking.Name
	booking.Date = updatedBooking.Date
	booking.Status = updatedBooking.Status

	// Simpan kembali booking yang sudah diperbarui
	updated, err := s.BookingRepository.Save(booking)
	if err != nil {
		return entity.Booking{}, err
	}

	return updated, nil

}

// DeleteBooking deletes a booking by its ID
func (s *BookingService) DeleteBooking(id int) error {
	// Convert the ID to a string if the repository expects a string
	idStr := strconv.Itoa(id)

	_, err := s.BookingRepository.FindByID(idStr)
	if err != nil {
		return errors.New("booking not found")
	}

	// Delete the booking from the repository
	err = s.BookingRepository.Delete(idStr)
	if err != nil {
		return err
	}

	return nil
}
