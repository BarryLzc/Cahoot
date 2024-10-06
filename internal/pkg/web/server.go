package web

import (
	"github.com/english-learning/internal/controller"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	router := gin.Default()

	// 注册陆游
	userRoutes := router.Group("/users")
	{
		userRoutes.GET("/", controller.GetUsers)
		userRoutes.POST("/", controller.CreateUser)
		userRoutes.GET("/:id", controller.GetUser)
		userRoutes.PUT("/:id", controller.UpdateUser)
		userRoutes.DELETE("/:id", controller.DeleteUser)
	}

	return router
}
