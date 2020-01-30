package controller

import (
	"github.com/Penetration-Platform-Go/Project-Service/model"
	"github.com/gin-gonic/gin"
)

// CreateProject Method
func CreateProject(ctx *gin.Context) {
	//username, _ := ctx.Get("username")

}

// QueryProject Method
func QueryProject(ctx *gin.Context) {
	username, _ := ctx.Get("username")
	result, err := model.QueryProject(username.(string))
	if err != nil {
		ctx.Status(400)
	} else {
		ctx.JSON(200, result)
	}
}

// UpdateProject Method
func UpdateProject(ctx *gin.Context) {
	//username, _ := ctx.Get("username")
}

// DeleteProject Method
func DeleteProject(ctx *gin.Context) {
	//username, _ := ctx.Get("username")
}
