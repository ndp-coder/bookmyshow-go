package main

import (
	"bookmyshow/controllers"
	"bookmyshow/database"
	"bookmyshow/midlewares"

	"github.com/gin-gonic/gin"

	"bookmyshow/routs"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		panic("Error loading .env file")
	}

	database.ConnectDB()

	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "bookmy show is running ",
		})

	})

	routs.Authrouts(router)

	protected := router.Group("/")
	protected.Use(midlewares.AuthMiddleware())
	{
		protected.POST("/bookseat", controllers.BookSeat)
		protected.GET("/getmovies", controllers.GetMovies)
		protected.GET("/gethalls", controllers.GetHalls)
		protected.GET("/gettheaters", controllers.GetTheaters)
		protected.POST("/payment", controllers.AddMovieInTheater)
		protected.POST("/AddTheater", controllers.AddTheater)
		protected.POST("/AddHall", controllers.AddHall)
		protected.POST("/AddMovie", controllers.AddMovie)
		protected.DELETE("/DeleteTheater", controllers.DeleteTheater)
		protected.DELETE("/DeleteHall", controllers.DeleteHall)
		protected.DELETE("/DeleteMovie", controllers.DeleteMovie)
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}

	log.Printf("Server running on port %s", port)
	if err := router.Run(":" + port); err != nil {
		log.Fatal("Failed to start server:", err)
	}

}
