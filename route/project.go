package route

import (
	"github.com/gin-gonic/gin"

	"github.com/Penetration-Platform-Go/Project-Service/middleware"

	"github.com/Penetration-Platform-Go/Project-Service/controller"
)

func projectServiceRoute(route *gin.RouterGroup) {

	route.GET("/", middleware.Auth(), controller.QueryProject)
	route.POST("/", middleware.Auth(), controller.CreateProject)
	route.PUT("/", middleware.Auth(), controller.UpdateProject)
	route.DELETE("/", middleware.Auth(), controller.DeleteProject)

}
