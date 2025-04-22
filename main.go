// /booking-system/main.go

package main

import (
	"booking-system/internal/controller"
	"booking-system/internal/repository"
	"booking-system/internal/service"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	// Inisialisasi repository, service, dan controller
	repo := repository.NewInMemoryBookingRepository()
	bookingService := service.NewBookingService(repo)
	bookingController := controller.NewBookingController(bookingService)

	// Setup Gin
	r := gin.Default()

	// Define routes
	r.POST("/book", bookingController.BookServiceHandler)      // Untuk pemesanan layanan
	r.GET("/booking/:id", bookingController.GetBookingHandler) // Untuk melihat booking berdasarkan ID

	// Jalankan server pada port 8080
	if err := r.Run(":8080"); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}
