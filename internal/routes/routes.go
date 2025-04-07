package routes

import (
	"ToGoList/internal/handlers"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	api := r.Group("/api/books")
	{
		api.GET("", handlers.GetBooks)
		api.GET("/:id", handlers.GetBookById)
		api.POST("", handlers.AddBook)
		api.PATCH("/update/:id", handlers.UpdateBook)
		api.DELETE("/delete/:id", handlers.DeleteBook)
	}
}
