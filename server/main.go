package main

import (
	"os"

	"github.com/Tejas24ytj/Go-Calorie-Tracker-App/routes"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(cors.Default())
	router.POST("/entry/create", routes.AddEntry)
	router.GET("/entries", routes.GetEntries)
	router.GET("/entry/:id/", routes.GetEntryById)
	router.GET("ingredient/:ingredient", routes.GetEntriesByIngredient)

	router.PUT("/entry/update/:id", routes.UpdateEntry)
	router.PUT("ingredient/:ingredient", routes.UpdateIngredient)
	router.DELETE("/entry/delete/:id,routes.DeleteEntry")
	router.RUN(":" + port)
}
