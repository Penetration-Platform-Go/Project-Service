package route

import "github.com/gin-gonic/gin"

// MainRoute 路由分组
func MainRoute(app *gin.Engine) {
	projectService := app.Group("/api/project")
	projectServiceRoute(projectService)
}
