package main

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/pandubhaskara/Go-API/db"
	"github.com/pandubhaskara/Go-API/handlers/product"
)

const (
	Port = "5000"
)

func init() {
	db.Connect()
}

func main() {
	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, "/articles")
	})
	router.GET("/products", product.List)
	router.POST("/products", product.Create)
	router.POST("/delete/producrs/:_id", product.Delete)

	port := Port
	if len(os.Getenv("PORT")) > 0 {
		port = os.Getenv("PORT")
	}
	router.Run(":" + port)
}
