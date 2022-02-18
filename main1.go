package main

import (
	"os"

	"github.com/gin-gonic/gin"
	// "github.com/pandubhaskara/Go-API/db"
	products "github.com/pandubhaskara/Go-API/handlers"
)

const (
	Port = "5000"
)

func init() {
	db.Connect()
}

func main1() {
	router := gin.Default()

	// router.GET("/", func(c *gin.Context) {
	// 	c.Redirect(http.StatusMovedPermanently, "/products")
	// })
	router.GET("/product", products.List)
	router.POST("/product", products.Create)
	router.POST("/delete/product/:_id", products.Delete)

	port := Port
	if len(os.Getenv("PORT")) > 0 {
		port = os.Getenv("PORT")
	}
	router.Run(":" + port)
}
