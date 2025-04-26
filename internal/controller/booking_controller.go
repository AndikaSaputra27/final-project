package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Tampilkan halaman user
func ShowUserPage(c *gin.Context) {
	c.HTML(http.StatusOK, "user.html", gin.H{
		"title": "Halaman User",
	})
}

// Tampilkan halaman admin
func ShowAdminPage(c *gin.Context) {
	c.HTML(http.StatusOK, "admin.html", gin.H{
		"title": "Halaman Admin",
	})
}
