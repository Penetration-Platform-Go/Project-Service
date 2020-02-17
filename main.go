package main

import (
	"github.com/Penetration-Platform-Go/Project-Service/middleware"
	"github.com/Penetration-Platform-Go/Project-Service/route"
	"github.com/gin-gonic/gin"
)

func main() {
	app := gin.Default()

	app.Use(middleware.Cors())
	// 动态绑定路由
	route.MainRoute(app)
	app.Run(":8002")
}
