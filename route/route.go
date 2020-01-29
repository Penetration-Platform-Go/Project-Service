package route

import "github.com/gin-gonic/gin"

// MainRoute 路由分组
func MainRoute(app *gin.Engine) {
	mainSerivice := app.Group("/api")
	mainServiceRoute(mainSerivice)
}
