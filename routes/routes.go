package routes

import (
	"eniqilo_store/controllers"

	"github.com/gin-gonic/gin"
)

func Init() *gin.Engine {
	staffController := new(controllers.StaffController)

	router := gin.New()
	v1 := router.Group("/v1")
	{
		staff := v1.Group("/staff")
		{
			staff.POST("/register", staffController.Register)
			staff.POST("/login", staffController.Login)
		}
	}

	return router
}