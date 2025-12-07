package routes

import (
	"management_buku/controllers"
	"management_buku/middlewares"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
    api := r.Group("/api")

    api.POST("/users/register", controllers.Register)
	api.POST("/users/login", controllers.Login)

    cat := api.Group("/categories")
    cat.Use(middlewares.JWTAuth())
    {
        cat.GET("", controllers.GetAllCategories)
        cat.POST("", controllers.CreateCategory)
        cat.GET("/:id", controllers.GetCategoryDetail)
		cat.GET("/:id/books", controllers.GetBooksByCategory)
        cat.DELETE("/:id", controllers.DeleteCategory)

    }

    books := api.Group("/books")
    books.Use(middlewares.JWTAuth())
    {
        books.GET("", controllers.GetAllBooks)
        books.POST("", controllers.CreateBook)
        books.GET("/:id", controllers.GetBookDetail)
        books.DELETE("/:id", controllers.DeleteBook)
    }
}
