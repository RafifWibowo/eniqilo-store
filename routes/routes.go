package routes

import (
	"eniqilo_store/controllers"
	"eniqilo_store/middleware"

	"github.com/gin-gonic/gin"
)

func Init() *gin.Engine {
	staffController := new(controllers.StaffController)
	customerController := new(controllers.CustomerController)

	router := gin.New()
	v1 := router.Group("/v1")
	{
		staff := v1.Group("/staff")
		{
			staff.POST("/register", staffController.Register)
			staff.POST("/login", staffController.Login)
		}
		v1.Use(middleware.AuthMiddleware)
		customer := v1.Group("/customer")
		{
			customer.POST("/register", customerController.CustomerRegister)
		}
	}

	return router
}