package main

import (
	"booking-system/internal/controller"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	bookingController := controller.NewBookingController()

	// Routing
	r.POST("/book", bookingController.BookServiceHandler)      // Untuk pemesanan layanan
	r.GET("/booking/:id", bookingController.GetBookingHandler) // Untuk melihat booking berdasarkan ID
	r.GET("/bookings", bookingController.GetBookingHandler)    // Semua bookings (jika ada logic untuk ini juga)

	// Jalankan server
	if err := r.Run(":8080"); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}
