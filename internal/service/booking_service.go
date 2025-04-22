package service

import (
	"booking-system/internal/entity"
	"booking-system/internal/repository"
)

type BookingService struct {
	repo repository.BookingRepository
}

func NewBookingService(r repository.BookingRepository) *BookingService {
	return &BookingService{repo: r}
}

func (s *BookingService) BookService(booking entity.Booking) (entity.Booking, error) {
	return s.repo.Save(booking)
}

func (s *BookingService) GetBooking(id int) (entity.Booking, error) {
	return s.repo.FindByID(id)
}
